// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceProtectionOverview undocumented
type DeviceProtectionOverview struct {
	// Object is the base model of DeviceProtectionOverview
	Object
	// TotalReportedDeviceCount Total device count.
	TotalReportedDeviceCount *int `json:"totalReportedDeviceCount,omitempty"`
	// InactiveThreatAgentDeviceCount Device with inactive threat agent count
	InactiveThreatAgentDeviceCount *int `json:"inactiveThreatAgentDeviceCount,omitempty"`
	// UnknownStateThreatAgentDeviceCount Device with threat agent state as unknown count.
	UnknownStateThreatAgentDeviceCount *int `json:"unknownStateThreatAgentDeviceCount,omitempty"`
	// PendingSignatureUpdateDeviceCount Device with old signature count.
	PendingSignatureUpdateDeviceCount *int `json:"pendingSignatureUpdateDeviceCount,omitempty"`
	// CleanDeviceCount Clean device count.
	CleanDeviceCount *int `json:"cleanDeviceCount,omitempty"`
	// PendingFullScanDeviceCount Pending full scan device count.
	PendingFullScanDeviceCount *int `json:"pendingFullScanDeviceCount,omitempty"`
	// PendingRestartDeviceCount Pending restart device count.
	PendingRestartDeviceCount *int `json:"pendingRestartDeviceCount,omitempty"`
	// PendingManualStepsDeviceCount Pending manual steps device count.
	PendingManualStepsDeviceCount *int `json:"pendingManualStepsDeviceCount,omitempty"`
	// PendingOfflineScanDeviceCount Pending offline scan device count.
	PendingOfflineScanDeviceCount *int `json:"pendingOfflineScanDeviceCount,omitempty"`
	// CriticalFailuresDeviceCount Critical failures device count.
	CriticalFailuresDeviceCount *int `json:"criticalFailuresDeviceCount,omitempty"`
}
