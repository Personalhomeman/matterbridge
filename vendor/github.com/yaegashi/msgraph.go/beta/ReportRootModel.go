// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ReportRoot The resource that represents an instance of Enrollment Failure Reports.
type ReportRoot struct {
	// Entity is the base model of ReportRoot
	Entity
	// ApplicationSignInDetailedSummary undocumented
	ApplicationSignInDetailedSummary []ApplicationSignInDetailedSummary `json:"applicationSignInDetailedSummary,omitempty"`
	// CredentialUserRegistrationDetails undocumented
	CredentialUserRegistrationDetails []CredentialUserRegistrationDetails `json:"credentialUserRegistrationDetails,omitempty"`
	// UserCredentialUsageDetails undocumented
	UserCredentialUsageDetails []UserCredentialUsageDetails `json:"userCredentialUsageDetails,omitempty"`
}
