// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsKioskMultipleApps undocumented
type WindowsKioskMultipleApps struct {
	// WindowsKioskAppConfiguration is the base model of WindowsKioskMultipleApps
	WindowsKioskAppConfiguration
	// Apps These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can contain a maximum of 128 elements.
	Apps []WindowsKioskAppBase `json:"apps,omitempty"`
	// ShowTaskBar This setting allows the admin to specify whether the Task Bar is shown or not.
	ShowTaskBar *bool `json:"showTaskBar,omitempty"`
	// AllowAccessToDownloadsFolder This setting allows access to Downloads folder in file explorer.
	AllowAccessToDownloadsFolder *bool `json:"allowAccessToDownloadsFolder,omitempty"`
	// DisallowDesktopApps This setting indicates that desktop apps are allowed. Default to true.
	DisallowDesktopApps *bool `json:"disallowDesktopApps,omitempty"`
	// StartMenuLayoutXML Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
	StartMenuLayoutXML *Binary `json:"startMenuLayoutXml,omitempty"`
}
