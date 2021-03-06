// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PermissionRequestBuilder is request builder for Permission
type PermissionRequestBuilder struct{ BaseRequestBuilder }

// Request returns PermissionRequest
func (b *PermissionRequestBuilder) Request() *PermissionRequest {
	return &PermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PermissionRequest is request for Permission
type PermissionRequest struct{ BaseRequest }

// Get performs GET request for Permission
func (r *PermissionRequest) Get(ctx context.Context) (resObj *Permission, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Permission
func (r *PermissionRequest) Update(ctx context.Context, reqObj *Permission) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Permission
func (r *PermissionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
