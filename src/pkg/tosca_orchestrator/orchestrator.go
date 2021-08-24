package tosca_orchestrator

import (
	"fmt"

	"github.com/thetillhoff/eat/pkg/tosca"
)

func Install(serviceTemplate tosca.ServiceTemplate) {
	var (
	//requirements []tosca.RequirementDefinition
	)

	fmt.Println("len(NodeTypes):", len(serviceTemplate.NodeTypes))

	// for key, value := range serviceTemplate.NodeTypes {
	// 	fmt.Println(key, value)
	// }

	serviceTemplate.NodeTypes["derivedFromImportable"].Properties["test"] = tosca.PropertyDefinition{Description: "Testproperty description"}

	fmt.Println(serviceTemplate.NodeTypes["derivedFromImportable"])

	fmt.Println("len(NodeTypes[derivedFromImportable].Properties):", len(serviceTemplate.NodeTypes["derivedFromImportable"].Properties))
	// for key, value := range serviceTemplate.NodeTypes["derivedFromImportable"].Properties {
	// 	fmt.Println(key, value)
	// }

	// requirements = serviceTemplate.TopologyTemplate.NodeTemplates[0].NodeType.Requirements

	// TODO
	// Retrieve

}
