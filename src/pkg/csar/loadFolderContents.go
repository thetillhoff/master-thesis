package csar

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path"
)

// Input <fromPath> to declare root-level of directory-tree to load from
//
// Returns contents of all files in format of map[<path>]<content> where <path> begins from <fromPath>.
func loadFolderContents(fromPath string) map[string]string {
	var (
		err               error
		dirContents       []fs.FileInfo
		element           fs.FileInfo
		elementPath       string
		contents          = make(map[string]string)
		subFolderContents = make(map[string]string)
		key               string
		value             string
	)

	dirContents, err = ioutil.ReadDir(fromPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, element = range dirContents {
		// if fromPath is "." and elementPath starts with "..", path.Join results in "..."
		if fromPath == "." {
			elementPath = element.Name()
		} else {
			elementPath = path.Join(fromPath, element.Name())
		}

		if element.IsDir() {
			subFolderContents = loadFolderContents(elementPath)
			for key, value = range subFolderContents {
				contents[key] = value
			}
		} else {
			fileContent, err := ioutil.ReadFile(elementPath)
			if err != nil {
				log.Fatalln(err)
			}
			contents[elementPath] = string(fileContent)
		}
	}

	return contents
}
