// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnServer undocumented
type VpnServer struct {
	// Object is the base model of VpnServer
	Object
	// Description Description.
	Description *string `json:"description,omitempty"`
	// Address Address (IP address, FQDN or URL)
	Address *string `json:"address,omitempty"`
	// IsDefaultServer Default server.
	IsDefaultServer *bool `json:"isDefaultServer,omitempty"`
}
