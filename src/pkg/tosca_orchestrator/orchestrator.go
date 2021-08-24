package tosca_orchestrator

import (
	"fmt"

	"github.com/thetillhoff/eat/pkg/tosca"
)

func Install(serviceTemplate tosca.ServiceTemplate) {
	var (
	//requirements []tosca.RequirementDefinition
	)

	fmt.Println(serviceTemplate.NodeTypes["derivedFromImportable"].ToString())

	// for key, value := range serviceTemplate.NodeTypes["derivedFromImportable"].Properties {
	// 	fmt.Println(key, value)
	// }

	// requirements = serviceTemplate.TopologyTemplate.NodeTemplates[0].NodeType.Requirements

	// TODO
	// Retrieve

}
