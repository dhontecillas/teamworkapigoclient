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

// FormsTokenApiService FormsTokenApi service
type FormsTokenApiService service

type ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest struct {
	ctx _context.Context
	ApiService *FormsTokenApiService
	formId int32
	tokenRequest *TokenRequest
}

func (r ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest) TokenRequest(tokenRequest TokenRequest) ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest {
	r.tokenRequest = &tokenRequest
	return r
}

func (r ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest) Execute() (TokenResponse, *_nethttp.Response, error) {
	return r.ApiService.PATCHProjectsApiV3FormsformIdTokenJsonExecute(r)
}

/*
 * PATCHProjectsApiV3FormsformIdTokenJson Update an existing token.
 * Update a form token
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param formId
 * @return ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest
 */
func (a *FormsTokenApiService) PATCHProjectsApiV3FormsformIdTokenJson(ctx _context.Context, formId int32) ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest {
	return ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest{
		ApiService: a,
		ctx: ctx,
		formId: formId,
	}
}

/*
 * Execute executes the request
 * @return TokenResponse
 */
func (a *FormsTokenApiService) PATCHProjectsApiV3FormsformIdTokenJsonExecute(r ApiPATCHProjectsApiV3FormsformIdTokenJsonRequest) (TokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormsTokenApiService.PATCHProjectsApiV3FormsformIdTokenJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/forms/{formId}/token.json"
	localVarPath = strings.Replace(localVarPath, "{"+"formId"+"}", _neturl.PathEscape(parameterToString(r.formId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.tokenRequest == nil {
		return localVarReturnValue, nil, reportError("tokenRequest is required and must be specified")
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
	localVarPostBody = r.tokenRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ViewErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest struct {
	ctx _context.Context
	ApiService *FormsTokenApiService
	formId int32
	tokenRequest *TokenRequest
}

func (r ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest) TokenRequest(tokenRequest TokenRequest) ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest {
	r.tokenRequest = &tokenRequest
	return r
}

func (r ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest) Execute() (TokenResponse, *_nethttp.Response, error) {
	return r.ApiService.PUTProjectsApiV3FormsformIdTokenRefreshJsonExecute(r)
}

/*
 * PUTProjectsApiV3FormsformIdTokenRefreshJson Refresh the value of a token
 * Used for invalidating current form URLs and generating
a new one.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param formId
 * @return ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest
 */
func (a *FormsTokenApiService) PUTProjectsApiV3FormsformIdTokenRefreshJson(ctx _context.Context, formId int32) ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest {
	return ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest{
		ApiService: a,
		ctx: ctx,
		formId: formId,
	}
}

/*
 * Execute executes the request
 * @return TokenResponse
 */
func (a *FormsTokenApiService) PUTProjectsApiV3FormsformIdTokenRefreshJsonExecute(r ApiPUTProjectsApiV3FormsformIdTokenRefreshJsonRequest) (TokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormsTokenApiService.PUTProjectsApiV3FormsformIdTokenRefreshJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/forms/{formId}/token/refresh.json"
	localVarPath = strings.Replace(localVarPath, "{"+"formId"+"}", _neturl.PathEscape(parameterToString(r.formId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.tokenRequest == nil {
		return localVarReturnValue, nil, reportError("tokenRequest is required and must be specified")
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
	localVarPostBody = r.tokenRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ViewErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
