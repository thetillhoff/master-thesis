package tosca_orchestrator

import (
	"fmt"

	"github.com/thetillhoff/eat/pkg/tosca"
)

func Install(serviceTemplate tosca.ServiceTemplate) {

	var value tosca.DataType = serviceTemplate.DataTypes["testString"]
	fmt.Println(value.ToString())

}
