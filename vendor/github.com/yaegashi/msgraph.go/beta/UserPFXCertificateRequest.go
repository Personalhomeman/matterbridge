// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UserPFXCertificateRequestBuilder is request builder for UserPFXCertificate
type UserPFXCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserPFXCertificateRequest
func (b *UserPFXCertificateRequestBuilder) Request() *UserPFXCertificateRequest {
	return &UserPFXCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserPFXCertificateRequest is request for UserPFXCertificate
type UserPFXCertificateRequest struct{ BaseRequest }

// Get performs GET request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Get(ctx context.Context) (resObj *UserPFXCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Update(ctx context.Context, reqObj *UserPFXCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}