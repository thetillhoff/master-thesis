package tosca

import (
	"bytes"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func PrintServiceTemplate(serviceTemplate ServiceTemplate) {
	var (
		err                 error
		serviceTemplateYaml string
		buffer              bytes.Buffer
		yamlEncoder         *yaml.Encoder
	)

	yamlEncoder = yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2) // Default is 4 spaces
	err = yamlEncoder.Encode(&serviceTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	defer yamlEncoder.Close()

	serviceTemplateYaml = buffer.String()

	fmt.Println(serviceTemplateYaml)
}
