package groups

// This is the default (root) TOSCA Group Type definition that all other TOSCA base Group Types derive from.
//
// Note:
//
// - Group operations are not necessarily tied directly to member nodes that are part of a group.
//
// - Future versions of this specification will create sub types of the tosca.groups.Root type that will describe how Group Type operations are to be orchestrated.
type Root struct {
}
