// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OutlookTaskGroup undocumented
type OutlookTaskGroup struct {
	// Entity is the base model of OutlookTaskGroup
	Entity
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// IsDefaultGroup undocumented
	IsDefaultGroup *bool `json:"isDefaultGroup,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// GroupKey undocumented
	GroupKey *UUID `json:"groupKey,omitempty"`
	// TaskFolders undocumented
	TaskFolders []OutlookTaskFolder `json:"taskFolders,omitempty"`
}
