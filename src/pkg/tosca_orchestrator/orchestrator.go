package tosca_orchestrator

import (
	"fmt"

	"github.com/thetillhoff/eat/pkg/tosca"
)

func Install(serviceTemplate tosca.ServiceTemplate, inputs []string) {

	// TODO: Add derivation from NodeType to NodeTemplate, so all properties etc are available.

	serviceTemplate.TopologyTemplate = AddInputsToTopologyTemplate(serviceTemplate.TopologyTemplate, inputs)

	// TODO: Add installation here

	serviceTemplate.TopologyTemplate = AddOutputsOfTopologyTemplate(serviceTemplate.TopologyTemplate)

	fmt.Println(serviceTemplate.ToString())

}
