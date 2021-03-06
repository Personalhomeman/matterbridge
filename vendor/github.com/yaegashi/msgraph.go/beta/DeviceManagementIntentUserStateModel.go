// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceManagementIntentUserState Entity that represents user state for an intent
type DeviceManagementIntentUserState struct {
	// Entity is the base model of DeviceManagementIntentUserState
	Entity
	// UserPrincipalName The user principal name that is being reported on a device
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserName The user name that is being reported on a device
	UserName *string `json:"userName,omitempty"`
	// DeviceCount Count of Devices that belongs to a user for an intent
	DeviceCount *int `json:"deviceCount,omitempty"`
	// LastReportedDateTime Last modified date time of an intent report
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
	// State User state for an intent
	State *ComplianceStatus `json:"state,omitempty"`
}
