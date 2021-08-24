package tosca

import (
	"log"
)

// Resolve derivation for all types in provided serviceTemplate
func (serviceTemplate ServiceTemplate) ResolveTypeDerivations() ServiceTemplate {

	// ArtifactType
	for typeName, typeDefinition := range serviceTemplate.ArtifactTypes {
		serviceTemplate.ArtifactTypes[typeName] = ResolveArtifactTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// DataType
	for typeName, typeDefinition := range serviceTemplate.DataTypes {
		serviceTemplate.DataTypes[typeName] = ResolveDataTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// CapabilityType
	for typeName, typeDefinition := range serviceTemplate.CapabilityTypes {
		serviceTemplate.CapabilityTypes[typeName] = ResolveCapabilityTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// InterfaceType
	for typeName, typeDefinition := range serviceTemplate.InterfaceTypes {
		serviceTemplate.InterfaceTypes[typeName] = ResolveInterfaceTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// RelationshipType
	for typeName, typeDefinition := range serviceTemplate.RelationshipTypes {
		serviceTemplate.RelationshipTypes[typeName] = ResolveRelationshipTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// NodeType
	for typeName, typeDefinition := range serviceTemplate.NodeTypes {
		serviceTemplate.NodeTypes[typeName] = ResolveNodeTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// GroupType
	for typeName, typeDefinition := range serviceTemplate.GroupTypes {
		serviceTemplate.GroupTypes[typeName] = ResolveGroupTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// PolicyType
	for typeName, typeDefinition := range serviceTemplate.PolicyTypes {
		serviceTemplate.PolicyTypes[typeName] = ResolvePolicyTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	return serviceTemplate
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveArtifactTypeDerivation(thisTypeName string, thisType ArtifactType, serviceTemplate ServiceTemplate) ArtifactType {
	var (
		parent  ArtifactType
		newType ArtifactType = NewArtifactType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving ArtifactType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.ArtifactTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No ArtifactType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.ArtifactTypes[type.derivedFrom]
		parent = serviceTemplate.ArtifactTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveArtifactTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// Take value from <src> (no derivation)
		newType.MimeType = thisType.MimeType

		// Take value from <src> (no derivation)
		newType.FileExt = thisType.FileExt

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveDataTypeDerivation(thisTypeName string, thisType DataType, serviceTemplate ServiceTemplate) DataType {
	var (
		parent  DataType
		newType DataType = NewDataType()

		emptySchema SchemaDefinition
	)
	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving DataType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.DataTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No DataType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.DataTypes[type.derivedFrom]
		parent = serviceTemplate.DataTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveDataTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// TODO
		// if dataType.DerivedFrom != "" { // Only if derived_from is set
		// 	switch dataType.DerivedFrom { // "standard datatypes"
		// 	case "string":
		// 		parentDataType = DataType{}
		// 	case "integer":
		// 		parentDataType = DataType{}
		// 	default:

		// First, add the parent Constraints
		newType.Constraints = append(newType.Constraints, parent.Constraints...)
		// Then, add the child Constraints
		newType.Constraints = append(newType.Constraints, thisType.Constraints...)

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
		}

		if thisType.KeySchema.Equals(emptySchema) { // If empty in <src>
			newType.KeySchema = parent.KeySchema // Use parent's value
		} else {
			newType.KeySchema = thisType.KeySchema // Use src's value
		}

		if thisType.Entryschema.Equals(emptySchema) { // If empty in <src>
			newType.Entryschema = parent.Entryschema // Use parent's value
		} else {
			newType.Entryschema = thisType.Entryschema // Use src's value
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveCapabilityTypeDerivation(thisTypeName string, thisType CapabilityType, serviceTemplate ServiceTemplate) CapabilityType {
	var (
		parent  CapabilityType
		newType CapabilityType = NewCapabilityType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving CapabilityType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.CapabilityTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No CapabilityType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.CapabilityTypes[type.derivedFrom]
		parent = serviceTemplate.CapabilityTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveCapabilityTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
		}

		// First, derive the parent Attributes
		for key, value := range parent.Attributes {
			newType.Attributes[key] = value
		}
		// Then, add/overwrite with child Attributes
		for key, value := range thisType.Attributes {
			newType.Attributes[key] = value
		}

		// if valid_source_types is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if valid_source_types is not defined in the parent type then no restrictions are applied.
		if len(parent.ValidSourceTypes) > 0 {
			// for _, sourceType := range src.ValidSourceTypes {
			// 	//TODO; add validation
			// }
		} else {
			newType.ValidSourceTypes = thisType.ValidSourceTypes
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveInterfaceTypeDerivation(thisTypeName string, thisType InterfaceType, serviceTemplate ServiceTemplate) InterfaceType {
	var (
		parent  InterfaceType
		newType InterfaceType = NewInterfaceType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving InterfaceType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.InterfaceTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No InterfaceType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.InterfaceTypes[type.derivedFrom]
		parent = serviceTemplate.InterfaceTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveInterfaceTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// First, derive the parent Inputs
		for key, value := range parent.Inputs {
			newType.Inputs[key] = value
		}
		// Then, add/overwrite with child Inputs
		for key, value := range thisType.Inputs {
			newType.Inputs[key] = value
		}

		// First, derive the parent Operations
		for key, value := range parent.Operations {
			newType.Operations[key] = value
		}
		// Then, add/overwrite with child Operations
		for key, value := range thisType.Operations {
			newType.Operations[key] = value
		}

		// First, derive the parent Notifications
		for key, value := range parent.Notifications {
			newType.Notifications[key] = value
		}
		// Then, add/overwrite with child Notifications
		for key, value := range thisType.Notifications {
			newType.Notifications[key] = value
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveRelationshipTypeDerivation(thisTypeName string, thisType RelationshipType, serviceTemplate ServiceTemplate) RelationshipType {
	var (
		parent  RelationshipType
		newType RelationshipType = NewRelationshipType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving RelationshipType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.RelationshipTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No RelationshipType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.RelationshipTypes[type.derivedFrom]
		parent = serviceTemplate.RelationshipTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveRelationshipTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
		}

		// First, derive the parent Attributes
		for key, value := range parent.Attributes {
			newType.Attributes[key] = value
		}
		// Then, add/overwrite with child Attributes
		for key, value := range thisType.Attributes {
			newType.Attributes[key] = value
		}

		// First, derive the parent Interfaces
		for key, value := range parent.Interfaces {
			newType.Interfaces[key] = value
		}
		// Then, add/overwrite with child Interfaces
		for key, value := range thisType.Interfaces {
			newType.Interfaces[key] = value
		}

		// if valid_target_types is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if valid_target_types is not defined in the parent type then no restrictions are applied.
		if len(parent.ValidTargetTypes) > 0 {
			// for _, sourceType := range src.ValidTargetTypes {
			// 	//TODO; add validation
			// }
		} else {
			newType.ValidTargetTypes = thisType.ValidTargetTypes
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveNodeTypeDerivation(thisTypeName string, thisType NodeType, serviceTemplate ServiceTemplate) NodeType {
	var (
		parent  NodeType
		newType NodeType = NewNodeType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving NodeType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.NodeTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No NodeType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.NodeTypes[type.derivedFrom]
		parent = serviceTemplate.NodeTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveNodeTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
			if debug {
				log.Println("INF parent property key:", key)
			}
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
			if debug {
				log.Println("INF child property key:", key)
			}
		}

		// First, derive the parent Attributes
		for key, value := range parent.Attributes {
			newType.Attributes[key] = value
		}
		// Then, add/overwrite with child Attributes
		for key, value := range thisType.Attributes {
			newType.Attributes[key] = value
		}

		// First, derive the parent Capabilities
		for key, value := range parent.Capabilities {
			newType.Capabilities[key] = value
		}
		// Then, add/overwrite with child Capabilities
		for key, value := range thisType.Capabilities {
			newType.Capabilities[key] = value
		}

		// First, derive the parent Requirements
		newType.Requirements = append(newType.Requirements, parent.Requirements...)
		// The, add/overwrite with child Requirements
		for _, value := range thisType.Requirements { // for each requirement in child
			var exists bool = false
			for _, existingRequirement := range newType.Requirements { // check all already copied requirements from parent node
				if existingRequirement.Equals(value) { // if one is equal, do nothing
					exists = true
				}
			}
			if !exists {
				newType.Requirements = append(newType.Requirements, value)
			}
		}

		// First, derive the parent Interfaces
		for key, value := range parent.Interfaces {
			newType.Interfaces[key] = value
		}
		// Then, add/overwrite with child Interfaces
		for key, value := range thisType.Interfaces {
			newType.Interfaces[key] = value
		}

		// First, derive the parent Artifacts
		for key, value := range parent.Artifacts {
			newType.Artifacts[key] = value
		}
		// Then, add/overwrite with child Artifacts
		for key, value := range thisType.Artifacts {
			newType.Artifacts[key] = value
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolveGroupTypeDerivation(thisTypeName string, thisType GroupType, serviceTemplate ServiceTemplate) GroupType {
	var (
		parent  GroupType
		newType GroupType = NewGroupType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving GroupType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.GroupTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No GroupType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.GroupTypes[type.derivedFrom]
		parent = serviceTemplate.GroupTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolveGroupTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
		}

		// First, derive the parent Attributes
		for key, value := range parent.Attributes {
			newType.Attributes[key] = value
		}
		// Then, add/overwrite with child Attributes
		for key, value := range thisType.Attributes {
			newType.Attributes[key] = value
		}

		// if the members keyname is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if the members keyname is not defined in the parent type then no restrictions are applied to the definition.
		if len(parent.Members) > 0 {
			// for _, sourceType := range src.ValidSourceTypes {
			// 	//TODO; add validation
			// }
		} else {
			newType.Members = thisType.Members
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func ResolvePolicyTypeDerivation(thisTypeName string, thisType PolicyType, serviceTemplate ServiceTemplate) PolicyType {
	var (
		parent  PolicyType
		newType PolicyType = NewPolicyType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !listContainsString(thisType.derivedFromAncestors, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving PolicyType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// check whether parent exists, if not: fail
		if _, ok := serviceTemplate.PolicyTypes[thisType.DerivedFrom]; !ok {
			log.Fatalln("ERR No PolicyType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
		}

		// retrieve parent type by name with serviceTemplate.PolicyTypes[type.derivedFrom]
		parent = serviceTemplate.PolicyTypes[thisType.DerivedFrom]

		// run same derivation for parent (recursion), which returns fully derived parent
		parent = ResolvePolicyTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

		// Resolve derivation with type and (now fully derived) parent

		// First, derive the parent Properties
		for key, value := range parent.Properties {
			newType.Properties[key] = value
		}
		// Then, add/overwrite with child Properties
		for key, value := range thisType.Properties {
			newType.Properties[key] = value
		}

		// if the targets keyname is defined in the parent type, each element in this list must either be in the parent type list or derived from an element in the parent type list; if the targets keyname is not defined in the parent type then no restrictions are applied to this definition.
		if len(parent.Targets) > 0 {
			// for _, sourceType := range src.ValidSourceTypes {
			// 	//TODO; add validation
			// }
		} else {
			newType.Targets = thisType.Targets
		}

		// First, derive the parent Triggers
		for key, value := range parent.Triggers {
			newType.Triggers[key] = value
		}
		// Then, add/overwrite with child Triggers
		for key, value := range thisType.Triggers {
			newType.Triggers[key] = value
		}

		// Add derivedFrom to derivedFromAncestors AND append parent.derivedFromAncestors to it.
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, parent.derivedFromAncestors...)
		newType.derivedFromAncestors = append(newType.derivedFromAncestors, thisType.DerivedFrom)
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}
