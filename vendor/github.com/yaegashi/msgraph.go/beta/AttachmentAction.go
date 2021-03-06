// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AttachmentCollectionCreateUploadSessionRequestParameter undocumented
type AttachmentCollectionCreateUploadSessionRequestParameter struct {
	// AttachmentItem undocumented
	AttachmentItem *AttachmentItem `json:"AttachmentItem,omitempty"`
}

//
type AttachmentCollectionCreateUploadSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateUploadSession action undocumented
func (b *EventAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateUploadSession action undocumented
func (b *MessageAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateUploadSession action undocumented
func (b *OutlookTaskAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateUploadSession action undocumented
func (b *PostAttachmentsCollectionRequestBuilder) CreateUploadSession(reqObj *AttachmentCollectionCreateUploadSessionRequestParameter) *AttachmentCollectionCreateUploadSessionRequestBuilder {
	bb := &AttachmentCollectionCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AttachmentCollectionCreateUploadSessionRequest struct{ BaseRequest }

//
func (b *AttachmentCollectionCreateUploadSessionRequestBuilder) Request() *AttachmentCollectionCreateUploadSessionRequest {
	return &AttachmentCollectionCreateUploadSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AttachmentCollectionCreateUploadSessionRequest) Post(ctx context.Context) (resObj *UploadSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
