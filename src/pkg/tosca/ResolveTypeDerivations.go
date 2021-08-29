package tosca

import (
	"log"
)

// Resolve derivation for all types in provided serviceTemplate
func ResolveTypeDerivations(serviceTemplate ServiceTemplate) ServiceTemplate {

	// ArtifactType
	for typeName, typeDefinition := range serviceTemplate.ArtifactTypes {
		serviceTemplate.ArtifactTypes[typeName] = resolveArtifactTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// DataType
	for typeName, typeDefinition := range serviceTemplate.DataTypes {
		serviceTemplate.DataTypes[typeName] = resolveDataTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// CapabilityType
	for typeName, typeDefinition := range serviceTemplate.CapabilityTypes {
		serviceTemplate.CapabilityTypes[typeName] = resolveCapabilityTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// InterfaceType
	for typeName, typeDefinition := range serviceTemplate.InterfaceTypes {
		serviceTemplate.InterfaceTypes[typeName] = resolveInterfaceTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// RelationshipType
	for typeName, typeDefinition := range serviceTemplate.RelationshipTypes {
		serviceTemplate.RelationshipTypes[typeName] = resolveRelationshipTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// NodeType
	for typeName, typeDefinition := range serviceTemplate.NodeTypes {
		serviceTemplate.NodeTypes[typeName] = resolveNodeTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// GroupType
	for typeName, typeDefinition := range serviceTemplate.GroupTypes {
		serviceTemplate.GroupTypes[typeName] = resolveGroupTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	// PolicyType
	for typeName, typeDefinition := range serviceTemplate.PolicyTypes {
		serviceTemplate.PolicyTypes[typeName] = resolvePolicyTypeDerivation(typeName, typeDefinition, serviceTemplate)
	}

	return serviceTemplate
}

// Resolve derivation of this type only - but completely (recursive)
func resolveArtifactTypeDerivation(thisTypeName string, thisType ArtifactType, serviceTemplate ServiceTemplate) ArtifactType {
	var (
		parent  ArtifactType
		newType ArtifactType = NewArtifactType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolveArtifactTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func resolveDataTypeDerivation(thisTypeName string, thisType DataType, serviceTemplate ServiceTemplate) DataType {
	var (
		parent  DataType
		newType DataType = NewDataType()

		emptySchema SchemaDefinition

		normativeTypes []string = []string{"string", "integer", "float", "boolean", "byte", "frequency", "time", "timestamp", "size", "range", "map", "list"}
	)

	// If derivation is necessary AND not already done AND not derived from a normativeType
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
		if debug {
			log.Println("INF Deriving DataType '" + thisTypeName + "' from parent '" + thisType.DerivedFrom + "'.")
		}

		// If thisType is derived from normativeType, don't derive further
		if listContainsString(normativeTypes, thisType.DerivedFrom) {
			return thisType
		} else {
			// check whether parent exists, if not: fail
			if _, ok := serviceTemplate.DataTypes[thisType.DerivedFrom]; !ok {
				log.Fatalln("ERR No DataType '" + thisType.DerivedFrom + "' in ServiceTemplate (parent of '" + thisTypeName + "').")
			}

			// retrieve parent type by name with serviceTemplate.DataTypes[type.derivedFrom]
			parent = serviceTemplate.DataTypes[thisType.DerivedFrom]

			// run same derivation for parent (recursion), which returns fully derived parent
			parent = resolveDataTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)
		}

		// Resolve derivation with type and (now fully derived) parent

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
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func resolveCapabilityTypeDerivation(thisTypeName string, thisType CapabilityType, serviceTemplate ServiceTemplate) CapabilityType {
	var (
		parent  CapabilityType
		newType CapabilityType = NewCapabilityType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolveCapabilityTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
func resolveInterfaceTypeDerivation(thisTypeName string, thisType InterfaceType, serviceTemplate ServiceTemplate) InterfaceType {
	var (
		parent  InterfaceType
		newType InterfaceType = NewInterfaceType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolveInterfaceTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func resolveRelationshipTypeDerivation(thisTypeName string, thisType RelationshipType, serviceTemplate ServiceTemplate) RelationshipType {
	var (
		parent  RelationshipType
		newType RelationshipType = NewRelationshipType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolveRelationshipTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func resolveNodeTypeDerivation(thisTypeName string, thisType NodeType, serviceTemplate ServiceTemplate) NodeType {
	var (
		parent  NodeType
		newType NodeType = NewNodeType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolveNodeTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
func resolveGroupTypeDerivation(thisTypeName string, thisType GroupType, serviceTemplate ServiceTemplate) GroupType {
	var (
		parent  GroupType
		newType GroupType = NewGroupType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolveGroupTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}

// Resolve derivation of this type only - but completely (recursive)
func resolvePolicyTypeDerivation(thisTypeName string, thisType PolicyType, serviceTemplate ServiceTemplate) PolicyType {
	var (
		parent  PolicyType
		newType PolicyType = NewPolicyType()
	)

	// If derivation is necessary AND not already done
	if thisType.DerivedFrom != "" && !IsDerivedFrom(thisType.AbstractType, thisType.DerivedFrom) {
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
		parent = resolvePolicyTypeDerivation(thisType.DerivedFrom, parent, serviceTemplate)

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
	} else if thisType.DerivedFrom == "" { // If derivation is NOT necessary
		newType = thisType
	}

	// return fully derived type (== derivedFromAncestors is filled with all necessary Ancestors AND properties etc contain all derived values)
	return newType
}
