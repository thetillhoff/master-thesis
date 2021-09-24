package docker

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/docker/docker/api/types"
)

func BuildImage(contextPath string, imageName string, noCache bool) {
	var (
		resp    types.ImageBuildResponse
		scanner *bufio.Scanner
	)

	log.Println("INF Building container image '" + imageName + "'.")

	buildContextFilePath := createBuildContext(contextPath)
	buildContextFile, err := os.Open(buildContextFilePath)
	if err != nil {
		log.Fatalln("ERR Can't open buildcontext at '"+buildContextFilePath+"':", err)
	}

	// Build container image
	resp, err = cli.ImageBuild(
		ctx,
		buildContextFile,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{imageName},
			NoCache:    noCache,
			// Remove: true,
			// BuildArgs: make(map[string]*string),
		})
	if err != nil {
		log.Fatalln("ERR Can't create container image:", err)
	}
	buildContextFile.Close()
	os.Remove(buildContextFilePath)
	defer resp.Body.Close()

	if debug {
		// Output is json stream; Print it in readable way
		scanner = bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			data := map[string]interface{}{}
			_ = json.Unmarshal([]byte(line), &data) // Don't care about errors during unmarshal (trusting docker)
			if stream, ok := data["stream"]; ok {
				if parsedLine, ok := stream.(string); ok {
					log.Print(parsedLine) // No Println here, since the unmarshalled json already contains newlines
				}
			}
		}
		if err = scanner.Err(); err != nil {
			log.Fatalln("ERR There was an error while creating the container image:", err)
		}
	}
}
