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
		mappedInputs        map[string]string = make(map[string]string)
	)

	// For each provided input
	for _, inputString := range inputs {

		// get name and value
		splittedInputString = strings.SplitN(inputString, "=", 2)
		inputName = splittedInputString[0]
		inputValue = splittedInputString[1]

		// Check if TopologyTemplate contains input with that name
		if _, ok := (topologyTemplate.Inputs)[inputName]; !ok {
			if debug {
				inputNames := []string{}
				for inputName := range topologyTemplate.Inputs {
					inputNames = append(inputNames, inputName)
				}
				log.Println("INF Valid input names are " + strings.Join(inputNames, ","))
			}
			log.Fatalln("ERR Unkown input with name '" + inputName + "'.")
		}

		// Add input to map (name and value still split)
		mappedInputs[inputName] = inputValue
	}

	for inputName, inputDefinition := range topologyTemplate.Inputs {
		input := (topologyTemplate.Inputs)[inputName]

		// Check if input was provided
		if &inputDefinition != nil {

		} else if _, ok := mappedInputs[inputName]; ok && (&inputDefinition == nil) {
			input.Value = mappedInputs[inputName]
			(topologyTemplate.Inputs)[inputName] = input
		} else if input.DefaultValue != nil {
			input.Value = input.DefaultValue
		} else {
			log.Fatalln("ERR Missing required input with name '" + inputName + "'.")
		}

		// Check if TopologyTemplate contains corresponding input
		if _, ok := (topologyTemplate.Inputs)[inputName]; !ok {
			inputNames := []string{}
			for inputName := range topologyTemplate.Inputs {
				inputNames = append(inputNames, inputName)
			}
			log.Println("INF Valid input names are " + strings.Join(inputNames, ","))
			log.Fatalln("ERR Unkown input with name '" + inputName + "'.")
		}

		// try to parse & write parsed value into interface
		switch *(topologyTemplate.Inputs)[inputName].DataType {
		case "string":
			var value, err = tosca.ParseString(inputValue)
			if err != nil {
				log.Fatalln("ERR Cannot parse input '"+inputName+"' to string", err)
			}

			// Workaround for assigning values to interfaces in maps, see https://stackoverflow.com/questions/17438253/accessing-struct-fields-inside-a-map-value-without-copying
			// assignableInput = &tosca.ParameterDefinition{}
			assignableInput = (topologyTemplate.Inputs)[inputName]
			assignableInput.Value = new(interface{})
			assignableInput.Value = value
			(topologyTemplate.Inputs)[inputName] = assignableInput
		// case "integer": // TODO
		// case "float": // TODO
		// TODO: add other normative types
		default:
			log.Fatalln("ERR Unkown type:", (topologyTemplate.Inputs)[inputName].DataType)
		}

		if debug {
			log.Println("INF Input '" + inputName + "' parsed.")
		}
	}

	return topologyTemplate
}
