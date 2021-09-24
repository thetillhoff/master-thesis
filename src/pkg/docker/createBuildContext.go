package docker

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Input folder-path, return file-path
//
// Deleting it after the build is up to the caller
func createBuildContext(dirPath string) string {

	// Creating tempfile
	tempFile, err := ioutil.TempFile("", "docker-*.buildcontext")
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()
	// defer os.Remove(file.Name()) -> up to the caller

	// Fill tar-file with content
	tarWriter := tar.NewWriter(tempFile)
	defer tarWriter.Close()
	filepath.Walk(dirPath, func(file string, fileInfo os.FileInfo, err error) error {
		// return on any error
		if err != nil {
			return err
		}

		// return on non-regular files (links)
		if !fileInfo.Mode().IsRegular() {
			return nil
		}

		// create a new dir/file header
		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}
		// write the header
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(tarWriter, f); err != nil {
			return err
		}

		// manually close here after each file operation; defering would cause each file close
		// to wait until all operations have completed.
		f.Close()

		return nil
	})

	if debug {
		log.Println("INF Created buildContext at " + tempFile.Name())
	}

	// Return path to build-context-file
	return tempFile.Name()
}
