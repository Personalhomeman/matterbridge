// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OpenShift undocumented
type OpenShift struct {
	// ChangeTrackedEntity is the base model of OpenShift
	ChangeTrackedEntity
	// SharedOpenShift undocumented
	SharedOpenShift *OpenShiftItem `json:"sharedOpenShift,omitempty"`
	// DraftOpenShift undocumented
	DraftOpenShift *OpenShiftItem `json:"draftOpenShift,omitempty"`
	// SchedulingGroupID undocumented
	SchedulingGroupID *string `json:"schedulingGroupId,omitempty"`
}
