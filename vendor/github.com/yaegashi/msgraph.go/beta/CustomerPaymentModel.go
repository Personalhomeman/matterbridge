// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// CustomerPayment undocumented
type CustomerPayment struct {
	// Entity is the base model of CustomerPayment
	Entity
	// JournalDisplayName undocumented
	JournalDisplayName *string `json:"journalDisplayName,omitempty"`
	// LineNumber undocumented
	LineNumber *int `json:"lineNumber,omitempty"`
	// CustomerID undocumented
	CustomerID *UUID `json:"customerId,omitempty"`
	// CustomerNumber undocumented
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// ContactID undocumented
	ContactID *string `json:"contactId,omitempty"`
	// PostingDate undocumented
	PostingDate *time.Time `json:"postingDate,omitempty"`
	// DocumentNumber undocumented
	DocumentNumber *string `json:"documentNumber,omitempty"`
	// ExternalDocumentNumber undocumented
	ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
	// Amount undocumented
	Amount *int `json:"amount,omitempty"`
	// AppliesToInvoiceID undocumented
	AppliesToInvoiceID *UUID `json:"appliesToInvoiceId,omitempty"`
	// AppliesToInvoiceNumber undocumented
	AppliesToInvoiceNumber *string `json:"appliesToInvoiceNumber,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Customer undocumented
	Customer *Customer `json:"customer,omitempty"`
}
