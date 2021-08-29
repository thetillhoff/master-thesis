package tosca

import "log"

//
// Intrinsic functions
// The functions are supported within the TOSCA template for manipulation of template data.
//

// The concat function is used to concatenate two or more string values within a TOSCA service template.
// concat: [<strings_or_expressions>*]

// The join function is used to join an array of strings into a single string with optional delimiter.
// join: [<strings_or_expressions>* [<delimiter>]]

// The token function is used within a TOSCA service template on a string to parse out (tokenize) substrings separated by one or more token characters within a larger string.
// token: [<string_with_tokens>, <string_with_chars>, <substring_index>]

//
// Property functions
// The get_input function is used within a service template to obtain template input parameter values. The get_property function is used to get property values from property definitions declared in the same service template (e.g. node or relationship templates).
// Note that the get_input and get_property functions may only retrieve the static values of parameter or property definitions of a TOSCA application as defined in the TOSCA Service Template.  The get_attribute function should be used to retrieve values for attribute definitions (or property definitions reflected as attribute definitions) from the runtime instance model of the TOSCA application (as realized by the TOSCA orchestrator).
//

// The get_input function is used to retrieve the values of parameters declared within the inputs section of a TOSCA Service Template.
//
// get_input: <input_parameter_name>
// get_input: [ <input_parameter_name>, <nested_input_parameter_name_or_index_1>, ..., <nested_input_parameter_name_or_index_n> ]
//
// <input_parameter_name>: The name of the parameter as defined in the inputs section of the service template.
// <nested_input_paratmer_name_or_index_*>: Some TOSCA input parameters are complex (i.e., composed as nested structures).  These parameters are used to dereference into the names of these nested structures when needed. Some parameters represent list types. In these cases, an index may be provided to reference a specific entry in the list (as identified by the previous parameter) to return.
func (topologyTemplate TopologyTemplate) GetInput(path []string) interface{} {
	var (
		value interface{}
	)

	// Check minimal length of path
	if len(path) < 1 {
		log.Fatalln("ERR get_input requires at least one element (input name).")
	} else if len(path) > 1 {
		log.Fatalln("ERR get_input with multiple parameters is not implemented (yet).")
	}

	if input, ok := topologyTemplate.Inputs[path[0]]; ok {
		value = input.Value
	} else {
		log.Fatalln("ERR No input with name '" + path[0] + "'.")
	}

	return value
}

// The get_property function is used to retrieve property values between modelable entities (e.g. NodeTemplate or RelationshipTemplate) defined in the same service template.
//
// get_property: [ <modelable_entity_name>, <optional_req_or_cap_name>, <property_name>, <nested_property_name_or_index_1>, ..., <nested_property_name_or_index_n> ]
//
// <modelable entity name> | SELF | SOURCE | TARGET | HOST: The mandatory name of a modelable entity (e.g., Node Template or Relationship Template name) as declared in the service template that contains the property definition the function will return the value from. See section B.1 for valid keywords.
// <optional_req_or_cap_name>: The optional name of the requirement or capability name within the modelable entity (i.e., the <modelable_entity_name> which contains the property definition the function will return the value from. Note: If the property definition is located in the modelable entity directly, then this parameter MAY be omitted.
// <property_name>: The name of the property definition the function will return the value from.
// <nested_property_name_or_index_*>: Some TOSCA properties are complex (i.e., composed as nested structures). These parameters are used to dereference into the names of these nested structures when needed. Some properties represent list types. In these cases, an index may be provided to reference a specific entry in the list (as identified by the previous parameter) to return.
func (topologyTemplate TopologyTemplate) GetProperty(path []string) interface{} {
	var (
		value interface{}
	)

	// Check minimal length of path
	if len(path) < 2 {
		log.Fatalln("ERR get_property requires at least elements (entity name, property name).")
	}

	if nodeTemplate, ok := topologyTemplate.NodeTemplates[path[0]]; ok {
		if nodeProperty, ok := nodeTemplate.Properties[path[1]]; ok {
			return nodeProperty
		} else {
			log.Println("WRN NodeTemplate '" + path[0] + "' found, but property '" + path[1] + "' doesn't exist.")
		}
	}
	if relationshipTemplate, ok := topologyTemplate.RelationshipTemplates[path[0]]; ok {
		if relationshipProperty, ok := relationshipTemplate.Properties[path[1]]; ok {
			return relationshipProperty
		} else {
			log.Println("WRN RelationshipTemplate '" + path[0] + "' found, but property '" + path[1] + "' doesn't exist.")
		}
	} else {
		log.Fatalln("ERR No property with that path:", path)
	}

	return value
}

//
// Attribute functions
// These functions (attribute functions) are used within an instance model to obtain attribute values from instances of nodes and relationships that have been created from an application model described in a service template.  The instances of nodes or relationships can be referenced by their name as assigned in the service template or relative to the context where they are being invoked.
//

// The get_attribute function is used to retrieve the values of attributes declared by the referenced node or relationship template name.
//
// get_attribute: [ <modelable_entity_name>, <optional_req_or_cap_name>, <attribute_name>, <nested_attribute_name_or_index_1>, ..., <nested_attribute_name_or_index_n> ]
//
// <modelable entity name> | SELF | SOURCE | TARGET | HOST: The mandatory name of a modelable entity (e.g., Node Template or Relationship Template name) as declared in the service template that contains the attribute definition the function will return the value from.  See section B.1 for valid keywords.
// <optional_req_or_cap_name>: The optional name of the requirement or capability name within the modelable entity (i.e., the <modelable_entity_name> which contains the attribute definition the function will return the value from. Note:  If the attribute definition is located in the modelable entity directly, then this parameter MAY be omitted.
//
// <attribute_name>: The name of the attribute definition the function will return the value from.
// <nested_attribute_name_or_index_*>: Some TOSCA attributes are complex (i.e., composed as nested structures). These parameters are used to dereference into the names of these nested structures when needed. Some attributes represent list types. In these cases, an index may be provided to reference a specific entry in the list (as identified by the previous parameter) to return.

//
// Operation functions
// These functions are used within an instance model to obtain values from interface operations. These can be used in order to set an attribute of a node instance at runtime or to pass values from one operation to another.
//

// The get_operation_output function is used to retrieve the values of variables exposed / exported from an interface operation.
// Note: If operation failed, then ignore its outputs. Orchestrators should allow orchestrators to continue running when possible past deployment in the lifecycle. For example, if an update fails, the application should be allowed to continue running and some other method will be used to alert administrators of the failure.
//
// get_operation_output: <modelable_entity_name>, <interface_name>, <operation_name>, <output_variable_name>
//
// <modelable entity name> | SELF | SOURCE | TARGET: The mandatory name of a modelable entity (e.g., Node Template or Relationship Template name) as declared in the service template that implements the interface and operation.
// <interface_name>: The mandatory name of the interface which defines the operation.
// <operation_name>: The mandatory name of the operation whose value we would like to retrieve.
// <output_variable_name>: The mandatory name of the variable that is exposed / exported by the operation.

//
// Navigation functions
//

// The get_nodes_of_type function can be used to retrieve a list of all known instances of nodes of the declared Node Type.
//
// get_nodes_of_type: <node_type_name>
//
// <node_type_name>: The mandatory name of a Node Type that a TOSCA orchestrator will use to search a running application instance in order to return all unique, node instances of that type.
//
// Returns
// TARGETS: The list of node instances from the current application instance that match the node_type_name supplied as an input parameter of this function.

//
// Artifact functions
// The get_artifact function is used to retrieve artifact location between modelable entities defined in the same service template.
//
// get_artifact: [ <modelable_entity_name>, <artifact_name>, <location>, <remove> ]
//
// <modelable entity name> | SELF | SOURCE | TARGET | HOST: The mandatory name of a modelable entity (e.g., Node Template or Relationship Template name) as declared in the service template that contains the property definition the function will return the value from. See section B.1 for valid keywords.
// <artifact_name>: The name of the artifact definition the function will return the value from.
// <location> | LOCAL_FILE: Location value must be either a valid path e.g. ‘/etc/var/my_file’ or ‘LOCAL_FILE’. If the value is LOCAL_FILE the orchestrator is responsible for providing a path as the result of the get_artifact call where the artifact file can be accessed. The orchestrator will also remove the artifact from this location at the end of the operation. If the location is a path specified by the user the orchestrator is responsible to copy the artifact to the specified location. The orchestrator will return the path as the value of the get_artifact function and leave the file here after the execution of the operation.
// remove: Boolean flag to override the orchestrator default behavior so it will remove or not the artifact at the end of the operation execution. If not specified the removal will depends of the location e.g. removes it in case of ‘LOCAL_FILE’ and keeps it in case of a path. If true the artifact will be removed by the orchestrator at the end of the operation execution, if false it will not be removed.
