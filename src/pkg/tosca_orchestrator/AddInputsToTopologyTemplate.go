package tosca_orchestrator

import (
	"log"
	"strings"

	"github.com/thetillhoff/eat/pkg/tosca"
)

// Parses inputs and puts them at correct place in topologyTemplate
func AddInputsToTopologyTemplate(topologyTemplate tosca.TopologyTemplate, inputs []string) tosca.TopologyTemplate {

	var (
		splittedInputString []string
		inputName           string
		inputValue          string
		assignableInput     tosca.ParameterDefinition
	)

	// For each provided input
	for _, inputString := range inputs {

		splittedInputString = strings.SplitN(inputString, "=", 2)
		inputName = splittedInputString[0]
		inputValue = splittedInputString[1]

		// Check if TopologyTemplate contains corresponding input
		if _, ok := topologyTemplate.Inputs[inputName]; !ok {
			log.Fatalln("ERR Unkown input with name '" + inputName + "'.")
		}

		// try to parse & write parsed value into interface
		switch topologyTemplate.Inputs[inputName].DataType {
		case "string":
			var value, err = tosca.ParseString(inputValue)
			if err != nil {
				log.Fatalln("ERR Cannot parse input '"+inputName+"' to string", err)
			}

			// Workaround for assigning values to interfaces in maps, see https://stackoverflow.com/questions/17438253/accessing-struct-fields-inside-a-map-value-without-copying
			assignableInput = topologyTemplate.Inputs[inputName]
			assignableInput.Value = value
			topologyTemplate.Inputs[inputName] = assignableInput
		// case "integer": // TODO
		// case "float": // TODO
		// TODO: add other normative types
		default:
			log.Fatalln("ERR Unkown type:", topologyTemplate.Inputs[inputName].DataType)
		}

		if debug {
			log.Println("INF Input '" + inputName + "' parsed.")
		}
	}

	return topologyTemplate
}
