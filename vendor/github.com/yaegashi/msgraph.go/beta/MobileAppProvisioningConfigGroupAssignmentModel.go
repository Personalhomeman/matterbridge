// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppProvisioningConfigGroupAssignment Contains the properties used to assign an App provisioning configuration to a group.
type MobileAppProvisioningConfigGroupAssignment struct {
	// Entity is the base model of MobileAppProvisioningConfigGroupAssignment
	Entity
	// TargetGroupID The ID of the AAD group in which the app provisioning configuration is being targeted.
	TargetGroupID *string `json:"targetGroupId,omitempty"`
}
