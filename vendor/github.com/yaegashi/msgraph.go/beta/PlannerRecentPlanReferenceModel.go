// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PlannerRecentPlanReference undocumented
type PlannerRecentPlanReference struct {
	// Object is the base model of PlannerRecentPlanReference
	Object
	// LastAccessedDateTime undocumented
	LastAccessedDateTime *time.Time `json:"lastAccessedDateTime,omitempty"`
	// PlanTitle undocumented
	PlanTitle *string `json:"planTitle,omitempty"`
}