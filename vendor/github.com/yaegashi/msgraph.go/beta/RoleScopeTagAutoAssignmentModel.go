// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RoleScopeTagAutoAssignment Contains the properties for auto-assigning a Role Scope Tag to a group to be applied to Devices.
type RoleScopeTagAutoAssignment struct {
	// Entity is the base model of RoleScopeTagAutoAssignment
	Entity
	// Target The auto-assignment target for the specific Role Scope Tag.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
