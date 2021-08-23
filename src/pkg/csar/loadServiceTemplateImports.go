package csar

import (
	"log"
	"path"

	"github.com/thetillhoff/eat/pkg/tosca"
)

// Takes ServiceTemplate and currentPath in CSAR
//
// Returns map of imported ServiceTemplates with namespace as key. Default namespace is "".
func loadServiceTemplateImports(serviceTemplate tosca.ServiceTemplate, currentPath string) map[string]tosca.ServiceTemplate {
	var (
		importedServiceTemplate  tosca.ServiceTemplate
		importedServiceTemplates map[string]tosca.ServiceTemplate
	)

	importedServiceTemplates = make(map[string]tosca.ServiceTemplate)

	// TODO: resolve repositories

	// For each import in provided serviceTemplate
	for _, providedImport := range serviceTemplate.Imports {
		// If import only contains url (only one implemented right now)
		if providedImport.Url != "" && providedImport.Repository == "" && providedImport.Profile == "" {
			// Prepare path of file
			// Since path.Join("a/b/c.yaml","../x/y.yaml") returns "/a/b/x/y.yaml", the former must use path.Dir first.
			pathToImport := path.Join(path.Dir(currentPath), providedImport.Url)

			// Parse imported ServiceTemplate
			importedServiceTemplate = parseServiceTemplate(pathToImport, "")

			// If namespace is unnamed
			if providedImport.Namespace == "" {
				importedServiceTemplates[""] = importedServiceTemplate
				if Debug {
					log.Println("INF Added imported ServiceTemplate at '" + pathToImport + "' into root-namespace.")
				}
				// TODO this overrides instead of merges
			} else { // namespace is named
				importedServiceTemplates[providedImport.Namespace] = importedServiceTemplate
				if Debug {
					log.Println("INF Added imported ServiceTemplate at '" + pathToImport + "' into namespace '" + providedImport.Namespace + "'.")
				}
			}

			for key, value := range loadServiceTemplateImports(importedServiceTemplate, pathToImport) {
				importedServiceTemplates[key] = value
			}
		} else {
			log.Fatalln("invalid import", providedImport)
		}

	}

	return importedServiceTemplates
}
