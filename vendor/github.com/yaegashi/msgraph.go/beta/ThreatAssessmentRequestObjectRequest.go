// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ThreatAssessmentRequestObjectRequestBuilder is request builder for ThreatAssessmentRequestObject
type ThreatAssessmentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns ThreatAssessmentRequestObjectRequest
func (b *ThreatAssessmentRequestObjectRequestBuilder) Request() *ThreatAssessmentRequestObjectRequest {
	return &ThreatAssessmentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ThreatAssessmentRequestObjectRequest is request for ThreatAssessmentRequestObject
type ThreatAssessmentRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for ThreatAssessmentRequestObject
func (r *ThreatAssessmentRequestObjectRequest) Get(ctx context.Context) (resObj *ThreatAssessmentRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ThreatAssessmentRequestObject
func (r *ThreatAssessmentRequestObjectRequest) Update(ctx context.Context, reqObj *ThreatAssessmentRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ThreatAssessmentRequestObject
func (r *ThreatAssessmentRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Results returns request builder for ThreatAssessmentResult collection
func (b *ThreatAssessmentRequestObjectRequestBuilder) Results() *ThreatAssessmentRequestObjectResultsCollectionRequestBuilder {
	bb := &ThreatAssessmentRequestObjectResultsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/results"
	return bb
}

// ThreatAssessmentRequestObjectResultsCollectionRequestBuilder is request builder for ThreatAssessmentResult collection
type ThreatAssessmentRequestObjectResultsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ThreatAssessmentResult collection
func (b *ThreatAssessmentRequestObjectResultsCollectionRequestBuilder) Request() *ThreatAssessmentRequestObjectResultsCollectionRequest {
	return &ThreatAssessmentRequestObjectResultsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ThreatAssessmentResult item
func (b *ThreatAssessmentRequestObjectResultsCollectionRequestBuilder) ID(id string) *ThreatAssessmentResultRequestBuilder {
	bb := &ThreatAssessmentResultRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ThreatAssessmentRequestObjectResultsCollectionRequest is request for ThreatAssessmentResult collection
type ThreatAssessmentRequestObjectResultsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ThreatAssessmentResult collection
func (r *ThreatAssessmentRequestObjectResultsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ThreatAssessmentResult, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ThreatAssessmentResult
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ThreatAssessmentResult
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ThreatAssessmentResult collection
func (r *ThreatAssessmentRequestObjectResultsCollectionRequest) Get(ctx context.Context) ([]ThreatAssessmentResult, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ThreatAssessmentResult collection
func (r *ThreatAssessmentRequestObjectResultsCollectionRequest) Add(ctx context.Context, reqObj *ThreatAssessmentResult) (resObj *ThreatAssessmentResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
