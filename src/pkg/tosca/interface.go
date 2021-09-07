package tosca

import (
	"bytes"
	"log"

	"gopkg.in/yaml.v3"
)

// interfaceTypes MUST NOT include any implementations for defined operations or notifications.
type InterfaceType struct {
	AbstractType `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The optional map of input parameter definitions available to all operations defined for this interface.
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// The optional map of operations defined for this interface.
	Operations map[string]OperationDefinition `yaml:"operations,omitempty" json:"operations,omitempty"`

	// The optional map of notifications defined for this interface.
	Notifications map[string]NotificationDefinition `yaml:"notifications,omitempty" json:"notifications,omitempty"`
}

func NewInterfaceType() InterfaceType {
	return InterfaceType{
		Inputs:        map[string]ParameterDefinition{},
		Operations:    map[string]OperationDefinition{},
		Notifications: map[string]NotificationDefinition{},
	}
}

func (interfaceType InterfaceType) ToString() string {
	var (
		err         error
		buffer      bytes.Buffer
		yamlEncoder *yaml.Encoder
	)

	yamlEncoder = yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2) // Default is 4 spaces
	err = yamlEncoder.Encode(&interfaceType)
	if err != nil {
		log.Fatalln(err)
	}
	defer yamlEncoder.Close()

	return buffer.String()
}

// An Interface definition defines an interface (containing operations and notifications definitions) that can be associated with (i.e. defined within) a Node or Relationship Type definition (including Interface definitions in Requirements definitions). An Interface definition may be refined in subsequent Node or Relationship Type derivations.
type InterfaceDefinition struct {

	// The mandatory name of the Interface Type this interface definition is based upon.
	InterfaceType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The optional description for this interface definition.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// The optional map of input parameter refinements and new input parameter definitions available to all operations defined for this interface (the input parameters to be refined have been defined in the Interface Type definition).
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// The optional map of operations refinements for this interface. The referred operations must have been defined in the Interface Type definition.
	Operations map[string]OperationDefinition `yaml:"operations,omitempty" json:"operations,omitempty"`

	// The optional map of notifications refinements for this interface. The referred operations must have been defined in the Interface Type definition.
	Notifications map[string]NotificationDefinition `yaml:"notifications,omitempty" json:"notifications,omitempty"`
}

type InterfaceAssignment struct {
	// The optional map of input parameter assignments. Template authors MAY provide parameter assignments for interface inputs that are not defined in their corresponding Interface Type.
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// The optional map of operations assignments specified for this interface.
	Operations map[string]OperationAssignment `yaml:"operations,omitempty" json:"operations,omitempty"`

	// The optional map of notifications assignments specified for this interface.
	Notifications map[string]NotificationAssignment `yaml:"notifications,omitempty" json:"notifications,omitempty"`
}

func NewInterfaceAssignment() InterfaceAssignment {
	return InterfaceAssignment{
		Inputs:        make(map[string]ParameterDefinition),
		Operations:    make(map[string]OperationAssignment),
		Notifications: make(map[string]NotificationAssignment),
	}
}

// func InitInterfaceAssignment(interfaceDefinition InterfaceDefinition, interfaceAssignment InterfaceAssignment) (InterfaceAssignment, error) {
// 	var (
// 		err        error
// 		assignment InterfaceAssignment = NewInterfaceAssignment()
// 	)

// 	for parameterName, parameterContent := range interfaceDefinition.Inputs {
// 		assignment.Inputs[parameterName] = parameterContent
// 	}

// 	for operationName, operationContent := range interfaceDefinition.Operations {
// 		assignment.Operations[operationName], err = InitOperationAssignment(operationContent, interfaceAssignment.Operations[operationName])
// 		if err != nil {
// 			return assignment, err
// 		}
// 	}

// 	for notificationName, notificationContent := range interfaceDefinition.Notifications {
// 		assignment.Notifications[notificationName], err = InitNotificationAssignment(notificationContent, interfaceAssignment.Notifications[notificationName])
// 		if err != nil {
// 			return assignment, err
// 		}
// 	}

// 	return assignment, nil
// }
