package csar

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"path"
)

// Input <fromPath> to point to CSAR zip-file to load
//
// Returns contents of all files in format of map[<path>]<content> where <path> begins from <fromPath>.
func loadZipContents(filePath string) map[string]string {

	var (
		zipFile  *zip.ReadCloser
		file     io.ReadCloser
		err      error
		element  *zip.File
		contents = make(map[string]string)
	)

	// Open zipfile
	zipFile, err = zip.OpenReader(filePath)
	if err != nil {
		log.Fatalln("ERR Unable to open CSAR file.", err)
	}
	defer zipFile.Close()

	// For each element in zipfile
	for _, element = range zipFile.File {

		// Only load yaml-files
		if path.Ext(element.Name) != ".yaml" && path.Ext(element.Name) != ".yml" {
			continue
		}

		// Open file within zipfile
		file, err = element.Open()
		if err != nil {
			log.Fatalln("ERR Unable to access file within CSAR archive.", err)
		}
		defer file.Close()

		// Retrieve contents of file within zipfile
		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln("ERR Unable to read contents from file within CSAR archive.", err)
		}
		contents[element.Name] = string(fileContent)
	}

	return contents
}
