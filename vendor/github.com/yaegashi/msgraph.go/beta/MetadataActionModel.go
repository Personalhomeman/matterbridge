// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MetadataAction undocumented
type MetadataAction struct {
	// InformationProtectionAction is the base model of MetadataAction
	InformationProtectionAction
	// MetadataToRemove undocumented
	MetadataToRemove []string `json:"metadataToRemove,omitempty"`
	// MetadataToAdd undocumented
	MetadataToAdd []KeyValuePair `json:"metadataToAdd,omitempty"`
}