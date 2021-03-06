// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IOSVpnConfigurationRequestBuilder is request builder for IOSVpnConfiguration
type IOSVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVpnConfigurationRequest
func (b *IOSVpnConfigurationRequestBuilder) Request() *IOSVpnConfigurationRequest {
	return &IOSVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSVpnConfigurationRequest is request for IOSVpnConfiguration
type IOSVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSVpnConfiguration
func (r *IOSVpnConfigurationRequest) Get(ctx context.Context) (resObj *IOSVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSVpnConfiguration
func (r *IOSVpnConfigurationRequest) Update(ctx context.Context, reqObj *IOSVpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSVpnConfiguration
func (r *IOSVpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DerivedCredentialSettings is navigation property
func (b *IOSVpnConfigurationRequestBuilder) DerivedCredentialSettings() *DeviceManagementDerivedCredentialSettingsRequestBuilder {
	bb := &DeviceManagementDerivedCredentialSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/derivedCredentialSettings"
	return bb
}

// IdentityCertificate is navigation property
func (b *IOSVpnConfigurationRequestBuilder) IdentityCertificate() *IOSCertificateProfileBaseRequestBuilder {
	bb := &IOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}
