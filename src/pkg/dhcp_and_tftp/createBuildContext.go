package dhcp_and_tftp

import (
	"archive/tar"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Creates buildcontext-tar-file and returns its path
//
// Deleting it after the build is up to the caller
func createBuildContext(buildPath string) string {
	// Creating tempfileName
	b := make([]rune, 16) // random string of length 16
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	var tempFileName = "docker-" + string(b) + ".buildcontext"

	// Create tar-file
	file, err := os.Create(tempFileName)
	if err != nil {
		log.Println("ERR Couldn't create buildcontext file:", err)
	}
	defer file.Close()

	// Fill tar-file with content
	tarWriter := tar.NewWriter(file)
	defer tarWriter.Close()
	filepath.Walk(buildPath, func(file string, fileInfo os.FileInfo, err error) error {
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

	// Return path to build-context-file
	return tempFileName
}
