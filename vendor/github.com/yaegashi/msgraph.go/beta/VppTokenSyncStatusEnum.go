// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VppTokenSyncStatus undocumented
type VppTokenSyncStatus int

const (
	// VppTokenSyncStatusVNone undocumented
	VppTokenSyncStatusVNone VppTokenSyncStatus = 0
	// VppTokenSyncStatusVInProgress undocumented
	VppTokenSyncStatusVInProgress VppTokenSyncStatus = 1
	// VppTokenSyncStatusVCompleted undocumented
	VppTokenSyncStatusVCompleted VppTokenSyncStatus = 2
	// VppTokenSyncStatusVFailed undocumented
	VppTokenSyncStatusVFailed VppTokenSyncStatus = 3
)

// VppTokenSyncStatusPNone returns a pointer to VppTokenSyncStatusVNone
func VppTokenSyncStatusPNone() *VppTokenSyncStatus {
	v := VppTokenSyncStatusVNone
	return &v
}

// VppTokenSyncStatusPInProgress returns a pointer to VppTokenSyncStatusVInProgress
func VppTokenSyncStatusPInProgress() *VppTokenSyncStatus {
	v := VppTokenSyncStatusVInProgress
	return &v
}

// VppTokenSyncStatusPCompleted returns a pointer to VppTokenSyncStatusVCompleted
func VppTokenSyncStatusPCompleted() *VppTokenSyncStatus {
	v := VppTokenSyncStatusVCompleted
	return &v
}

// VppTokenSyncStatusPFailed returns a pointer to VppTokenSyncStatusVFailed
func VppTokenSyncStatusPFailed() *VppTokenSyncStatus {
	v := VppTokenSyncStatusVFailed
	return &v
}
