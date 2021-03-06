// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// WindowsInformationProtectionDeviceRegistration Represents device registration records for Bring-Your-Own-Device(BYOD) Windows devices.
type WindowsInformationProtectionDeviceRegistration struct {
	// Entity is the base model of WindowsInformationProtectionDeviceRegistration
	Entity
	// UserID UserId associated with this device registration record.
	UserID *string `json:"userId,omitempty"`
	// DeviceRegistrationID Device identifier for this device registration record.
	DeviceRegistrationID *string `json:"deviceRegistrationId,omitempty"`
	// DeviceName Device name.
	DeviceName *string `json:"deviceName,omitempty"`
	// DeviceType Device type, for example, Windows laptop VS Windows phone.
	DeviceType *string `json:"deviceType,omitempty"`
	// DeviceMacAddress Device Mac address.
	DeviceMacAddress *string `json:"deviceMacAddress,omitempty"`
	// LastCheckInDateTime Last checkin time of the device.
	LastCheckInDateTime *time.Time `json:"lastCheckInDateTime,omitempty"`
}
