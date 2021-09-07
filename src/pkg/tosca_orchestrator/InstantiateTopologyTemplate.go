package tosca_orchestrator

import "github.com/thetillhoff/eat/pkg/tosca"

func InstantiateTopologyTemplate(serviceTemplate tosca.ServiceTemplate) tosca.TopologyTemplate {
	// TODO
	return *serviceTemplate.TopologyTemplate
}
