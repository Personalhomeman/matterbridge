// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MeteredConnectionLimitType undocumented
type MeteredConnectionLimitType int

const (
	// MeteredConnectionLimitTypeVUnrestricted undocumented
	MeteredConnectionLimitTypeVUnrestricted MeteredConnectionLimitType = 0
	// MeteredConnectionLimitTypeVFixed undocumented
	MeteredConnectionLimitTypeVFixed MeteredConnectionLimitType = 1
	// MeteredConnectionLimitTypeVVariable undocumented
	MeteredConnectionLimitTypeVVariable MeteredConnectionLimitType = 2
)

// MeteredConnectionLimitTypePUnrestricted returns a pointer to MeteredConnectionLimitTypeVUnrestricted
func MeteredConnectionLimitTypePUnrestricted() *MeteredConnectionLimitType {
	v := MeteredConnectionLimitTypeVUnrestricted
	return &v
}

// MeteredConnectionLimitTypePFixed returns a pointer to MeteredConnectionLimitTypeVFixed
func MeteredConnectionLimitTypePFixed() *MeteredConnectionLimitType {
	v := MeteredConnectionLimitTypeVFixed
	return &v
}

// MeteredConnectionLimitTypePVariable returns a pointer to MeteredConnectionLimitTypeVVariable
func MeteredConnectionLimitTypePVariable() *MeteredConnectionLimitType {
	v := MeteredConnectionLimitTypeVVariable
	return &v
}
