// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationSynchronizationCustomization undocumented
type EducationSynchronizationCustomization struct {
	// Object is the base model of EducationSynchronizationCustomization
	Object
	// OptionalPropertiesToSync undocumented
	OptionalPropertiesToSync []string `json:"optionalPropertiesToSync,omitempty"`
	// SynchronizationStartDate undocumented
	SynchronizationStartDate *time.Time `json:"synchronizationStartDate,omitempty"`
	// IsSyncDeferred undocumented
	IsSyncDeferred *bool `json:"isSyncDeferred,omitempty"`
	// AllowDisplayNameUpdate undocumented
	AllowDisplayNameUpdate *bool `json:"allowDisplayNameUpdate,omitempty"`
}
