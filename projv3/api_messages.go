/*
 * Teamwork.com Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package projv3

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// MessagesApiService MessagesApi service
type MessagesApiService service

type ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest struct {
	ctx _context.Context
	ApiService *MessagesApiService
	messageId int32
	messageRequest *MessageRequest
}

func (r ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest) MessageRequest(messageRequest MessageRequest) ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest {
	r.messageRequest = &messageRequest
	return r
}

func (r ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PATCHProjectsApiV3MessagesmessageIdJsonExecute(r)
}

/*
 * PATCHProjectsApiV3MessagesmessageIdJson Edit a message.
 * Update specific fields on a message. If the property `read` is sent with the
value `true` all posts from this message will be marked as read.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param messageId
 * @return ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest
 */
func (a *MessagesApiService) PATCHProjectsApiV3MessagesmessageIdJson(ctx _context.Context, messageId int32) ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest {
	return ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		messageId: messageId,
	}
}

/*
 * Execute executes the request
 */
func (a *MessagesApiService) PATCHProjectsApiV3MessagesmessageIdJsonExecute(r ApiPATCHProjectsApiV3MessagesmessageIdJsonRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesApiService.PATCHProjectsApiV3MessagesmessageIdJson")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/messages/{messageId}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", _neturl.PathEscape(parameterToString(r.messageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.messageRequest == nil {
		return nil, reportError("messageRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.messageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ViewErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}