// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceComplianceDeviceOverviewRequestBuilder is request builder for DeviceComplianceDeviceOverview
type DeviceComplianceDeviceOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceComplianceDeviceOverviewRequest
func (b *DeviceComplianceDeviceOverviewRequestBuilder) Request() *DeviceComplianceDeviceOverviewRequest {
	return &DeviceComplianceDeviceOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceComplianceDeviceOverviewRequest is request for DeviceComplianceDeviceOverview
type DeviceComplianceDeviceOverviewRequest struct{ BaseRequest }

// Get performs GET request for DeviceComplianceDeviceOverview
func (r *DeviceComplianceDeviceOverviewRequest) Get(ctx context.Context) (resObj *DeviceComplianceDeviceOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceComplianceDeviceOverview
func (r *DeviceComplianceDeviceOverviewRequest) Update(ctx context.Context, reqObj *DeviceComplianceDeviceOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceComplianceDeviceOverview
func (r *DeviceComplianceDeviceOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
