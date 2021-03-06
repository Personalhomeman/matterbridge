// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MobileAppCategoryRequestBuilder is request builder for MobileAppCategory
type MobileAppCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppCategoryRequest
func (b *MobileAppCategoryRequestBuilder) Request() *MobileAppCategoryRequest {
	return &MobileAppCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppCategoryRequest is request for MobileAppCategory
type MobileAppCategoryRequest struct{ BaseRequest }

// Get performs GET request for MobileAppCategory
func (r *MobileAppCategoryRequest) Get(ctx context.Context) (resObj *MobileAppCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppCategory
func (r *MobileAppCategoryRequest) Update(ctx context.Context, reqObj *MobileAppCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppCategory
func (r *MobileAppCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
