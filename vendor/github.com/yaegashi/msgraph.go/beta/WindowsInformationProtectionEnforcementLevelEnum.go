// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionEnforcementLevel undocumented
type WindowsInformationProtectionEnforcementLevel int

const (
	// WindowsInformationProtectionEnforcementLevelVNoProtection undocumented
	WindowsInformationProtectionEnforcementLevelVNoProtection WindowsInformationProtectionEnforcementLevel = 0
	// WindowsInformationProtectionEnforcementLevelVEncryptAndAuditOnly undocumented
	WindowsInformationProtectionEnforcementLevelVEncryptAndAuditOnly WindowsInformationProtectionEnforcementLevel = 1
	// WindowsInformationProtectionEnforcementLevelVEncryptAuditAndPrompt undocumented
	WindowsInformationProtectionEnforcementLevelVEncryptAuditAndPrompt WindowsInformationProtectionEnforcementLevel = 2
	// WindowsInformationProtectionEnforcementLevelVEncryptAuditAndBlock undocumented
	WindowsInformationProtectionEnforcementLevelVEncryptAuditAndBlock WindowsInformationProtectionEnforcementLevel = 3
)

// WindowsInformationProtectionEnforcementLevelPNoProtection returns a pointer to WindowsInformationProtectionEnforcementLevelVNoProtection
func WindowsInformationProtectionEnforcementLevelPNoProtection() *WindowsInformationProtectionEnforcementLevel {
	v := WindowsInformationProtectionEnforcementLevelVNoProtection
	return &v
}

// WindowsInformationProtectionEnforcementLevelPEncryptAndAuditOnly returns a pointer to WindowsInformationProtectionEnforcementLevelVEncryptAndAuditOnly
func WindowsInformationProtectionEnforcementLevelPEncryptAndAuditOnly() *WindowsInformationProtectionEnforcementLevel {
	v := WindowsInformationProtectionEnforcementLevelVEncryptAndAuditOnly
	return &v
}

// WindowsInformationProtectionEnforcementLevelPEncryptAuditAndPrompt returns a pointer to WindowsInformationProtectionEnforcementLevelVEncryptAuditAndPrompt
func WindowsInformationProtectionEnforcementLevelPEncryptAuditAndPrompt() *WindowsInformationProtectionEnforcementLevel {
	v := WindowsInformationProtectionEnforcementLevelVEncryptAuditAndPrompt
	return &v
}

// WindowsInformationProtectionEnforcementLevelPEncryptAuditAndBlock returns a pointer to WindowsInformationProtectionEnforcementLevelVEncryptAuditAndBlock
func WindowsInformationProtectionEnforcementLevelPEncryptAuditAndBlock() *WindowsInformationProtectionEnforcementLevel {
	v := WindowsInformationProtectionEnforcementLevelVEncryptAuditAndBlock
	return &v
}
