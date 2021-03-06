// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSTrustedRootCertificate OS X Trusted Root Certificate configuration profile.
type MacOSTrustedRootCertificate struct {
	// DeviceConfiguration is the base model of MacOSTrustedRootCertificate
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate.
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
}
