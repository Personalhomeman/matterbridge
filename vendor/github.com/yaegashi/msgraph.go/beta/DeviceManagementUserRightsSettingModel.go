// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementUserRightsSetting undocumented
type DeviceManagementUserRightsSetting struct {
	// Object is the base model of DeviceManagementUserRightsSetting
	Object
	// State Representing the current state of this user rights setting
	State *StateManagementSetting `json:"state,omitempty"`
	// LocalUsersOrGroups Representing a collection of local users or groups which will be set on device if the state of this setting is Allowed. This collection can contain a maximum of 500 elements.
	LocalUsersOrGroups []DeviceManagementUserRightsLocalUserOrGroup `json:"localUsersOrGroups,omitempty"`
}
