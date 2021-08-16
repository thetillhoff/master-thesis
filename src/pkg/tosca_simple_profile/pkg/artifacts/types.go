package artifacts

// This is the default (root) TOSCA Artifact Type definition that all other TOSCA base Artifact Types derive from.
type Root struct {
}

// This artifact type is used when an artifact definition needs to have its associated file simply treated as a file and no special handling/handlers are invoked (i.e., it is not treated as either an implementation or deployment artifact type).
type File Root

// This artifact type represents the parent type for all deployment artifacts in TOSCA. This class of artifacts typically represents a binary packaging of an application or service that is used to install/create or deploy it as part of a node’s lifecycle.
type Deployment Root

// This artifact type represents a parent type for any “image” which is an opaque packaging of a TOSCA Node’s deployment (whether real or virtual) whose contents are typically already installed and pre-configured (i.e., “stateful”) and prepared to be run on a known target container.
type DeploymentImage Deployment

// This artifact represents the parent type for all Virtual Machine (VM) image and container formatted deployment artifacts. These images contain a stateful capture of a machine (e.g., server) including operating system and installed software along with any configurations and can be run on another machine using a hypervisor which virtualizes typical server (i.e., hardware) resources.
type DeploymentImageVM DeploymentImage

// This artifact type represents the parent type for all implementation artifacts in TOSCA. These artifacts are used to implement operations of TOSCA interfaces either directly (e.g., scripts) or indirectly (e.g., config. files).
type Implementation Root

// This artifact type represents a Bash script type that contains Bash commands that can be executed on the Unix Bash shell.
type ImplementationBash Implementation

// This artifact type represents a Python file that contains Python language constructs that can be executed within a Python interpreter.
type ImplementationPython Implementation

// This artifact type represents the parent type for all template type s in TOSCA. This class of artifacts typically represent template files that are dependent artifacts for implementation of an interface or deployment of a node.
//
// Like the case of other dependent artifacts, the TOSCA orchestrator copies the dependent artifacts to the same location as the primary artifact for its access during execution. However, the template artifact processor need not be deployed in the environment where the primary artifact executes.  At run time, the Orchestrator can invoke the artifact processor (i.e. template engine) to fill in run time values and provide the “filled template” to the primary artifact processor for further processing.
//
// This reduces the requirements on the primary artifact target environment and the processing time of template artifacts.
type Template Root
