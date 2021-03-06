// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DimensionValueRequestBuilder is request builder for DimensionValue
type DimensionValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns DimensionValueRequest
func (b *DimensionValueRequestBuilder) Request() *DimensionValueRequest {
	return &DimensionValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DimensionValueRequest is request for DimensionValue
type DimensionValueRequest struct{ BaseRequest }

// Get performs GET request for DimensionValue
func (r *DimensionValueRequest) Get(ctx context.Context) (resObj *DimensionValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DimensionValue
func (r *DimensionValueRequest) Update(ctx context.Context, reqObj *DimensionValue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DimensionValue
func (r *DimensionValueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
