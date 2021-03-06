// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DirectoryRoleRequestBuilder is request builder for DirectoryRole
type DirectoryRoleRequestBuilder struct{ BaseRequestBuilder }

// Request returns DirectoryRoleRequest
func (b *DirectoryRoleRequestBuilder) Request() *DirectoryRoleRequest {
	return &DirectoryRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DirectoryRoleRequest is request for DirectoryRole
type DirectoryRoleRequest struct{ BaseRequest }

// Get performs GET request for DirectoryRole
func (r *DirectoryRoleRequest) Get(ctx context.Context) (resObj *DirectoryRole, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DirectoryRole
func (r *DirectoryRoleRequest) Update(ctx context.Context, reqObj *DirectoryRole) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DirectoryRole
func (r *DirectoryRoleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Members returns request builder for DirectoryObject collection
func (b *DirectoryRoleRequestBuilder) Members() *DirectoryRoleMembersCollectionRequestBuilder {
	bb := &DirectoryRoleMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// DirectoryRoleMembersCollectionRequestBuilder is request builder for DirectoryObject collection
type DirectoryRoleMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DirectoryRoleMembersCollectionRequestBuilder) Request() *DirectoryRoleMembersCollectionRequest {
	return &DirectoryRoleMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DirectoryRoleMembersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryRoleMembersCollectionRequest is request for DirectoryObject collection
type DirectoryRoleMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ScopedMembers returns request builder for ScopedRoleMembership collection
func (b *DirectoryRoleRequestBuilder) ScopedMembers() *DirectoryRoleScopedMembersCollectionRequestBuilder {
	bb := &DirectoryRoleScopedMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scopedMembers"
	return bb
}

// DirectoryRoleScopedMembersCollectionRequestBuilder is request builder for ScopedRoleMembership collection
type DirectoryRoleScopedMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ScopedRoleMembership collection
func (b *DirectoryRoleScopedMembersCollectionRequestBuilder) Request() *DirectoryRoleScopedMembersCollectionRequest {
	return &DirectoryRoleScopedMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ScopedRoleMembership item
func (b *DirectoryRoleScopedMembersCollectionRequestBuilder) ID(id string) *ScopedRoleMembershipRequestBuilder {
	bb := &ScopedRoleMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryRoleScopedMembersCollectionRequest is request for ScopedRoleMembership collection
type DirectoryRoleScopedMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ScopedRoleMembership, error) {
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
	var values []ScopedRoleMembership
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
			value  []ScopedRoleMembership
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

// Get performs GET request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Get(ctx context.Context) ([]ScopedRoleMembership, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Add(ctx context.Context, reqObj *ScopedRoleMembership) (resObj *ScopedRoleMembership, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
