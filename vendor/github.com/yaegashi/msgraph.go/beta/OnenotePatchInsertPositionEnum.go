// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnenotePatchInsertPosition undocumented
type OnenotePatchInsertPosition int

const (
	// OnenotePatchInsertPositionVAfter undocumented
	OnenotePatchInsertPositionVAfter OnenotePatchInsertPosition = 0
	// OnenotePatchInsertPositionVBefore undocumented
	OnenotePatchInsertPositionVBefore OnenotePatchInsertPosition = 1
)

// OnenotePatchInsertPositionPAfter returns a pointer to OnenotePatchInsertPositionVAfter
func OnenotePatchInsertPositionPAfter() *OnenotePatchInsertPosition {
	v := OnenotePatchInsertPositionVAfter
	return &v
}

// OnenotePatchInsertPositionPBefore returns a pointer to OnenotePatchInsertPositionVBefore
func OnenotePatchInsertPositionPBefore() *OnenotePatchInsertPosition {
	v := OnenotePatchInsertPositionVBefore
	return &v
}
