// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ObjectFlowTypes undocumented
type ObjectFlowTypes int

const (
	// ObjectFlowTypesVNone undocumented
	ObjectFlowTypesVNone ObjectFlowTypes = 0
	// ObjectFlowTypesVAdd undocumented
	ObjectFlowTypesVAdd ObjectFlowTypes = 1
	// ObjectFlowTypesVUpdate undocumented
	ObjectFlowTypesVUpdate ObjectFlowTypes = 2
	// ObjectFlowTypesVDelete undocumented
	ObjectFlowTypesVDelete ObjectFlowTypes = 4
)

// ObjectFlowTypesPNone returns a pointer to ObjectFlowTypesVNone
func ObjectFlowTypesPNone() *ObjectFlowTypes {
	v := ObjectFlowTypesVNone
	return &v
}

// ObjectFlowTypesPAdd returns a pointer to ObjectFlowTypesVAdd
func ObjectFlowTypesPAdd() *ObjectFlowTypes {
	v := ObjectFlowTypesVAdd
	return &v
}

// ObjectFlowTypesPUpdate returns a pointer to ObjectFlowTypesVUpdate
func ObjectFlowTypesPUpdate() *ObjectFlowTypes {
	v := ObjectFlowTypesVUpdate
	return &v
}

// ObjectFlowTypesPDelete returns a pointer to ObjectFlowTypesVDelete
func ObjectFlowTypesPDelete() *ObjectFlowTypes {
	v := ObjectFlowTypesVDelete
	return &v
}
