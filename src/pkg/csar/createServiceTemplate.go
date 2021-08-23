package csar

func (csar CSAR) createSingleServiceTemplate() CSAR {

	// Load serviceTemplates in Imports
	csar.imports = loadServiceTemplateImports(csar.ServiceTemplate, csar.EntryDefinition)

	// Resolve derivations
	csar = csar.resolveDerivations()

	return csar
}
