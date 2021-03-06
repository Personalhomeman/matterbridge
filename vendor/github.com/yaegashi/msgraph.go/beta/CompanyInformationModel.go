// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// CompanyInformation undocumented
type CompanyInformation struct {
	// Entity is the base model of CompanyInformation
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Address undocumented
	Address *PostalAddressType `json:"address,omitempty"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// FaxNumber undocumented
	FaxNumber *string `json:"faxNumber,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// Website undocumented
	Website *string `json:"website,omitempty"`
	// TaxRegistrationNumber undocumented
	TaxRegistrationNumber *string `json:"taxRegistrationNumber,omitempty"`
	// CurrencyCode undocumented
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// CurrentFiscalYearStartDate undocumented
	CurrentFiscalYearStartDate *time.Time `json:"currentFiscalYearStartDate,omitempty"`
	// Industry undocumented
	Industry *string `json:"industry,omitempty"`
	// Picture undocumented
	Picture *Stream `json:"picture,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
