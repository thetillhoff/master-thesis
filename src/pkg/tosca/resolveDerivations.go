package tosca

import (
	"errors"
	"log"
)

// Resolve all derivations of servicetemplate
func (serviceTemplate ServiceTemplate) ResolveDerivations() ServiceTemplate {
	var (
		err error
	)

	// Resolve NodeType derivations
	for key, nodeType := range serviceTemplate.NodeTypes {
		nodeType, err = nodeType.resolveNodeTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for nodeType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.NodeTypes[key] = nodeType
	}

	// Resolve ArtifactType derivations
	for key, artifactType := range serviceTemplate.ArtifactTypes {
		artifactType, err = artifactType.resolveArtifactTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for artifactType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.ArtifactTypes[key] = artifactType
	}

	// Resolve DataType derivations
	for key, dataType := range serviceTemplate.DataTypes {
		dataType, err = dataType.resolveDataTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for dataType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.DataTypes[key] = dataType
	}

	// Resolve CapabilityType derivations
	for key, capabilityType := range serviceTemplate.CapabilityTypes {
		capabilityType, err = capabilityType.resolveCapabilityTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for capabilityType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.CapabilityTypes[key] = capabilityType
	}

	// Resolve InterfaceType derivations
	for key, interfaceType := range serviceTemplate.InterfaceTypes {
		interfaceType, err = interfaceType.resolveInterfaceTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for interfaceType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.InterfaceTypes[key] = interfaceType
	}

	// Resolve RelationshipType derivations
	for key, relationshipType := range serviceTemplate.RelationshipTypes {
		relationshipType, err = relationshipType.resolveRelationshipTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for relationshipType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.RelationshipTypes[key] = relationshipType
	}

	// Resolve GroupType derivations
	for key, groupType := range serviceTemplate.GroupTypes {
		groupType, err = groupType.resolveGroupTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for groupType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.GroupTypes[key] = groupType
	}

	// Resolve PolicyType derivations
	for key, policyType := range serviceTemplate.PolicyTypes {
		policyType, err = policyType.resolvePolicyTypeDerivation(serviceTemplate)
		if err != nil {
			log.Fatalln("derivation for policyType '"+key+"' couldn't be resolved:", err)
		}
		serviceTemplate.PolicyTypes[key] = policyType
	}

	return serviceTemplate
}

// Resolve derivation of this type only - but completely (recursive)
func (artifactType ArtifactType) resolveArtifactTypeDerivation(serviceTemplate ServiceTemplate) (ArtifactType, error) {
	var (
		newArtifactType    ArtifactType
		parentArtifactType ArtifactType
		err                error
	)

	newArtifactType = NewArtifactType()

	if artifactType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.ArtifactTypes[artifactType.DerivedFrom]; ok {
			parentArtifactType = value
			if parentArtifactType.DerivedFrom != "" {
				parentArtifactType, err = parentArtifactType.resolveArtifactTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return artifactType, errors.New("derivation of parent artifactType '" + artifactType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return artifactType, errors.New("no ArtifactType with name '" + artifactType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentArtifactType = NewArtifactType()
	}

	// Take value from <src> (no derivation)
	newArtifactType.MimeType = artifactType.MimeType

	// Take value from <src> (no derivation)
	newArtifactType.FileExt = artifactType.FileExt

	// First, derive the parent Properties
	for key, value := range parentArtifactType.Properties {
		newArtifactType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range artifactType.Properties {
		newArtifactType.Properties[key] = value
	}

	return newArtifactType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (dataType DataType) resolveDataTypeDerivation(serviceTemplate ServiceTemplate) (DataType, error) {
	var (
		newDataType    DataType
		parentDataType DataType
		err            error
		emptySchema    SchemaDefinition
	)

	newDataType = NewDataType()

	if dataType.DerivedFrom != "" { // Only if derived_from is set
		switch dataType.DerivedFrom { // "standard datatypes"
		case "string":
			parentDataType = DataType{}
		case "integer":
			parentDataType = DataType{}
		default:
			if value, ok := serviceTemplate.DataTypes[dataType.DerivedFrom]; ok {
				parentDataType = value
				if parentDataType.DerivedFrom != "" {
					parentDataType, err = parentDataType.resolveDataTypeDerivation(serviceTemplate) // recursion
					if err != nil {
						return dataType, errors.New("derivation of parent dataType '" + dataType.DerivedFrom + "' couldn't be resolved")
					}
				}
			} else {
				return dataType, errors.New("no DataType with name '" + dataType.DerivedFrom + "' exists in serviceTemplate")
			}
		}
	} else {
		parentDataType = NewDataType()
	}

	// First, add the parent Constraints
	newDataType.Constraints = append(newDataType.Constraints, parentDataType.Constraints...)
	// Then, add the child Constraints
	newDataType.Constraints = append(newDataType.Constraints, dataType.Constraints...)

	// First, derive the parent Properties
	for key, value := range parentDataType.Properties {
		newDataType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range dataType.Properties {
		newDataType.Properties[key] = value
	}

	if dataType.KeySchema.Equals(emptySchema) { // If empty in <src>
		newDataType.KeySchema = parentDataType.KeySchema // Use parent's value
	} else {
		newDataType.KeySchema = dataType.KeySchema // Use src's value
	}

	if dataType.Entryschema.Equals(emptySchema) { // If empty in <src>
		newDataType.Entryschema = parentDataType.Entryschema // Use parent's value
	} else {
		newDataType.Entryschema = dataType.Entryschema // Use src's value
	}

	return newDataType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (capabilityType CapabilityType) resolveCapabilityTypeDerivation(serviceTemplate ServiceTemplate) (CapabilityType, error) {
	var (
		newCapabilityType    CapabilityType
		parentCapabilityType CapabilityType
		err                  error
	)

	newCapabilityType = NewCapabilityType()

	if capabilityType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.CapabilityTypes[capabilityType.DerivedFrom]; ok {
			parentCapabilityType = value
			if parentCapabilityType.DerivedFrom != "" {
				parentCapabilityType, err = parentCapabilityType.resolveCapabilityTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return capabilityType, errors.New("derivation of parent capabilityType '" + capabilityType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return capabilityType, errors.New("no CapabilityType with name '" + capabilityType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentCapabilityType = NewCapabilityType()
	}

	// First, derive the parent Properties
	for key, value := range parentCapabilityType.Properties {
		newCapabilityType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range capabilityType.Properties {
		newCapabilityType.Properties[key] = value
	}

	// First, derive the parent Attributes
	for key, value := range parentCapabilityType.Attributes {
		newCapabilityType.Attributes[key] = value
	}
	// Then, add/overwrite with child Attributes
	for key, value := range capabilityType.Attributes {
		newCapabilityType.Attributes[key] = value
	}

	// if valid_source_types is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if valid_source_types is not defined in the parent type then no restrictions are applied.
	if len(parentCapabilityType.ValidSourceTypes) > 0 {
		// for _, sourceType := range src.ValidSourceTypes {
		// 	//TODO; add validation
		// }
	} else {
		newCapabilityType.ValidSourceTypes = capabilityType.ValidSourceTypes
	}

	return newCapabilityType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (interfaceType InterfaceType) resolveInterfaceTypeDerivation(serviceTemplate ServiceTemplate) (InterfaceType, error) {
	var (
		newInterfaceType    InterfaceType
		parentInterfaceType InterfaceType
		err                 error
	)

	newInterfaceType = NewInterfaceType()

	if interfaceType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.InterfaceTypes[interfaceType.DerivedFrom]; ok {
			parentInterfaceType = value
			if parentInterfaceType.DerivedFrom != "" {
				parentInterfaceType, err = parentInterfaceType.resolveInterfaceTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return interfaceType, errors.New("derivation of parent interfaceType '" + interfaceType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return interfaceType, errors.New("no InterfaceType with name '" + interfaceType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentInterfaceType = NewInterfaceType()
	}

	// First, derive the parent Inputs
	for key, value := range parentInterfaceType.Inputs {
		newInterfaceType.Inputs[key] = value
	}
	// Then, add/overwrite with child Inputs
	for key, value := range interfaceType.Inputs {
		newInterfaceType.Inputs[key] = value
	}

	// First, derive the parent Operations
	for key, value := range parentInterfaceType.Operations {
		newInterfaceType.Operations[key] = value
	}
	// Then, add/overwrite with child Operations
	for key, value := range interfaceType.Operations {
		newInterfaceType.Operations[key] = value
	}

	// First, derive the parent Notifications
	for key, value := range parentInterfaceType.Notifications {
		newInterfaceType.Notifications[key] = value
	}
	// Then, add/overwrite with child Notifications
	for key, value := range interfaceType.Notifications {
		newInterfaceType.Notifications[key] = value
	}

	return newInterfaceType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (relationshipType RelationshipType) resolveRelationshipTypeDerivation(serviceTemplate ServiceTemplate) (RelationshipType, error) {
	var (
		newRelationshipType    RelationshipType
		parentRelationshipType RelationshipType
		err                    error
	)

	newRelationshipType = NewRelationshipType()

	if relationshipType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.RelationshipTypes[relationshipType.DerivedFrom]; ok {
			parentRelationshipType = value
			if parentRelationshipType.DerivedFrom != "" {
				parentRelationshipType, err = parentRelationshipType.resolveRelationshipTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return relationshipType, errors.New("derivation of parent relationshipType '" + relationshipType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return relationshipType, errors.New("no RelationshipType with name '" + relationshipType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentRelationshipType = NewRelationshipType()
	}

	// First, derive the parent Properties
	for key, value := range parentRelationshipType.Properties {
		newRelationshipType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range relationshipType.Properties {
		newRelationshipType.Properties[key] = value
	}

	// First, derive the parent Attributes
	for key, value := range parentRelationshipType.Attributes {
		newRelationshipType.Attributes[key] = value
	}
	// Then, add/overwrite with child Attributes
	for key, value := range relationshipType.Attributes {
		newRelationshipType.Attributes[key] = value
	}

	// First, derive the parent Interfaces
	for key, value := range parentRelationshipType.Interfaces {
		newRelationshipType.Interfaces[key] = value
	}
	// Then, add/overwrite with child Interfaces
	for key, value := range relationshipType.Interfaces {
		newRelationshipType.Interfaces[key] = value
	}

	// if valid_target_types is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if valid_target_types is not defined in the parent type then no restrictions are applied.
	if len(parentRelationshipType.ValidTargetTypes) > 0 {
		// for _, sourceType := range src.ValidTargetTypes {
		// 	//TODO; add validation
		// }
	} else {
		newRelationshipType.ValidTargetTypes = relationshipType.ValidTargetTypes
	}

	return newRelationshipType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (nodeType NodeType) resolveNodeTypeDerivation(serviceTemplate ServiceTemplate) (NodeType, error) {
	var (
		newNodeType    NodeType
		parentNodeType NodeType
		err            error
	)

	newNodeType = NewNodeType()

	if nodeType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.NodeTypes[nodeType.DerivedFrom]; ok {
			parentNodeType = value
			if parentNodeType.DerivedFrom != "" {
				parentNodeType, err = parentNodeType.resolveNodeTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return nodeType, errors.New("derivation of parent nodeType '" + nodeType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return nodeType, errors.New("no NodeType with name '" + nodeType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentNodeType = NewNodeType()
	}

	// First, derive the parent Properties
	for key, value := range parentNodeType.Properties {
		newNodeType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range nodeType.Properties {
		newNodeType.Properties[key] = value
	}

	// First, derive the parent Attributes
	for key, value := range parentNodeType.Attributes {
		newNodeType.Attributes[key] = value
	}
	// Then, add/overwrite with child Attributes
	for key, value := range nodeType.Attributes {
		newNodeType.Attributes[key] = value
	}

	// First, derive the parent Capabilities
	for key, value := range parentNodeType.Capabilities {
		newNodeType.Capabilities[key] = value
	}
	// Then, add/overwrite with child Capabilities
	for key, value := range nodeType.Capabilities {
		newNodeType.Capabilities[key] = value
	}

	// First, derive the parent Requirements
	newNodeType.Requirements = append(newNodeType.Requirements, parentNodeType.Requirements...)
	// The, add/overwrite with child Requirements
	for _, value := range nodeType.Requirements { // for each requirement in child
		var exists bool = false
		for _, existingRequirement := range newNodeType.Requirements { // check all already copied requirements from parent node
			if existingRequirement.Equals(value) { // if one is equal, do nothing
				exists = true
			}
		}
		if !exists {
			newNodeType.Requirements = append(newNodeType.Requirements, value)
		}
	}

	// First, derive the parent Interfaces
	for key, value := range parentNodeType.Interfaces {
		newNodeType.Interfaces[key] = value
	}
	// Then, add/overwrite with child Interfaces
	for key, value := range nodeType.Interfaces {
		newNodeType.Interfaces[key] = value
	}

	// First, derive the parent Artifacts
	for key, value := range parentNodeType.Artifacts {
		newNodeType.Artifacts[key] = value
	}
	// Then, add/overwrite with child Artifacts
	for key, value := range nodeType.Artifacts {
		newNodeType.Artifacts[key] = value
	}

	return newNodeType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (groupType GroupType) resolveGroupTypeDerivation(serviceTemplate ServiceTemplate) (GroupType, error) {
	var (
		newGroupType    GroupType
		parentGroupType GroupType
		err             error
	)

	newGroupType = NewGroupType()

	if groupType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.GroupTypes[groupType.DerivedFrom]; ok {
			parentGroupType = value
			if parentGroupType.DerivedFrom != "" {
				parentGroupType, err = parentGroupType.resolveGroupTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return groupType, errors.New("derivation of parent groupType '" + groupType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return groupType, errors.New("no GroupType with name '" + groupType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentGroupType = NewGroupType()
	}

	// First, derive the parent Properties
	for key, value := range parentGroupType.Properties {
		newGroupType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range groupType.Properties {
		newGroupType.Properties[key] = value
	}

	// First, derive the parent Attributes
	for key, value := range parentGroupType.Attributes {
		newGroupType.Attributes[key] = value
	}
	// Then, add/overwrite with child Attributes
	for key, value := range groupType.Attributes {
		newGroupType.Attributes[key] = value
	}

	// if the members keyname is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if the members keyname is not defined in the parent type then no restrictions are applied to the definition.
	if len(parentGroupType.Members) > 0 {
		// for _, sourceType := range src.ValidSourceTypes {
		// 	//TODO; add validation
		// }
	} else {
		newGroupType.Members = groupType.Members
	}

	return newGroupType, nil
}

// Resolve derivation of this type only - but completely (recursive)
func (policyType PolicyType) resolvePolicyTypeDerivation(serviceTemplate ServiceTemplate) (PolicyType, error) {
	var (
		newPolicyType    PolicyType
		parentPolicyType PolicyType
		err              error
	)

	newPolicyType = NewPolicyType()

	if policyType.DerivedFrom != "" { // Only if derived_from is set
		if value, ok := serviceTemplate.PolicyTypes[policyType.DerivedFrom]; ok {
			parentPolicyType = value
			if parentPolicyType.DerivedFrom != "" {
				parentPolicyType, err = parentPolicyType.resolvePolicyTypeDerivation(serviceTemplate) // recursion
				if err != nil {
					return policyType, errors.New("derivation of parent policyType '" + policyType.DerivedFrom + "' couldn't be resolved")
				}
			}
		} else {
			return policyType, errors.New("no PolicyType with name '" + policyType.DerivedFrom + "' exists in serviceTemplate")
		}
	} else {
		parentPolicyType = NewPolicyType()
	}

	// First, derive the parent Properties
	for key, value := range parentPolicyType.Properties {
		newPolicyType.Properties[key] = value
	}
	// Then, add/overwrite with child Properties
	for key, value := range policyType.Properties {
		newPolicyType.Properties[key] = value
	}

	// if the targets keyname is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if the targets keyname is not defined in the parent type then no restrictions are applied to this definition.
	if len(parentPolicyType.Targets) > 0 {
		// for _, sourceType := range src.ValidSourceTypes {
		// 	//TODO; add validation
		// }
	} else {
		newPolicyType.Targets = policyType.Targets
	}

	// First, derive the parent Triggers
	for key, value := range parentPolicyType.Triggers {
		newPolicyType.Triggers[key] = value
	}
	// Then, add/overwrite with child Triggers
	for key, value := range policyType.Triggers {
		newPolicyType.Triggers[key] = value
	}

	return newPolicyType, nil
}

// ? TopologyTemplate TopologyTemplateDefinition

// func (serviceTemplate ServiceTemplate) isDerivedFrom(child AbstractType, ancestorName string) bool {
// 	if child.DerivedFrom == ancestorName {
// 		return true
// 	} else {
// 		return isDerivedFrom()
// 	}
// 	return false
// }
