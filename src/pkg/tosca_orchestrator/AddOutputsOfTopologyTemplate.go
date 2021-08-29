package tosca_orchestrator

import (
	"log"

	"github.com/thetillhoff/eat/pkg/tosca"
)

func AddOutputsOfTopologyTemplate(topologyTemplate tosca.TopologyTemplate) tosca.TopologyTemplate {

	for outputName, output := range topologyTemplate.Outputs {
		if mapping, ok := output.Value.(map[string]interface{}); ok {
			if len(mapping) > 1 {
				log.Fatalln("ERR Output '" + outputName + "' contains more than one main function")
			}
			for key, value := range mapping { // Only one element, but this is the easiest & probably fastest way to get key and value
				if listing, ok := value.([]interface{}); ok {
					var typedListing = []string{}
					for _, element := range listing {
						if typedValue, ok := element.(string); ok {
							typedListing = append(typedListing, typedValue)
						}
					}
					switch key {
					case "get_property":
						output.Value = topologyTemplate.GetProperty(typedListing)
					case "get_input":
						output.Value = topologyTemplate.GetInput(typedListing)
					default:
						log.Fatalln("ERR Invalid function name '" + key + "' at output '" + outputName + "'.")
					}
				} else {
					log.Fatalln("ERR Function in output '" + outputName + "' is invalid. Parameters not a list.")
				}
			}
		} else {
			log.Fatalln("ERR Output '" + outputName + "' is invalid. Not of type map[string].")
		}
		topologyTemplate.Outputs[outputName] = output
		if debug {
			log.Println("INF Output '" + outputName + "' filled.")
		}
	}

	return topologyTemplate
}
