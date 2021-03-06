// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Item undocumented
type Item struct {
	// Entity is the base model of Item
	Entity
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// ItemCategoryID undocumented
	ItemCategoryID *UUID `json:"itemCategoryId,omitempty"`
	// ItemCategoryCode undocumented
	ItemCategoryCode *string `json:"itemCategoryCode,omitempty"`
	// Blocked undocumented
	Blocked *bool `json:"blocked,omitempty"`
	// BaseUnitOfMeasureID undocumented
	BaseUnitOfMeasureID *UUID `json:"baseUnitOfMeasureId,omitempty"`
	// Gtin undocumented
	Gtin *string `json:"gtin,omitempty"`
	// Inventory undocumented
	Inventory *int `json:"inventory,omitempty"`
	// UnitPrice undocumented
	UnitPrice *int `json:"unitPrice,omitempty"`
	// PriceIncludesTax undocumented
	PriceIncludesTax *bool `json:"priceIncludesTax,omitempty"`
	// UnitCost undocumented
	UnitCost *int `json:"unitCost,omitempty"`
	// TaxGroupID undocumented
	TaxGroupID *UUID `json:"taxGroupId,omitempty"`
	// TaxGroupCode undocumented
	TaxGroupCode *string `json:"taxGroupCode,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Picture undocumented
	Picture []Picture `json:"picture,omitempty"`
	// ItemCategory undocumented
	ItemCategory *ItemCategory `json:"itemCategory,omitempty"`
}
