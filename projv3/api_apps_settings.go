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

// AppsSettingsApiService AppsSettingsApi service
type AppsSettingsApiService service

type ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest struct {
	ctx _context.Context
	ApiService *AppsSettingsApiService
	settingId int32
	appId int32
	settingRequest *SettingRequest
}

func (r ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest) SettingRequest(settingRequest SettingRequest) ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest {
	r.settingRequest = &settingRequest
	return r
}

func (r ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest) Execute() (SettingResponse, *_nethttp.Response, error) {
	return r.ApiService.PUTProjectsApiV3AppsappIdSettingssettingIdJsonExecute(r)
}

/*
 * PUTProjectsApiV3AppsappIdSettingssettingIdJson Update an existing setting.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param settingId
 * @param appId
 * @return ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest
 */
func (a *AppsSettingsApiService) PUTProjectsApiV3AppsappIdSettingssettingIdJson(ctx _context.Context, settingId int32, appId int32) ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest {
	return ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		settingId: settingId,
		appId: appId,
	}
}

/*
 * Execute executes the request
 * @return SettingResponse
 */
func (a *AppsSettingsApiService) PUTProjectsApiV3AppsappIdSettingssettingIdJsonExecute(r ApiPUTProjectsApiV3AppsappIdSettingssettingIdJsonRequest) (SettingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SettingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppsSettingsApiService.PUTProjectsApiV3AppsappIdSettingssettingIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/apps/{appId}/settings/{settingId}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", _neturl.PathEscape(parameterToString(r.settingId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", _neturl.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.settingRequest == nil {
		return localVarReturnValue, nil, reportError("settingRequest is required and must be specified")
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
	localVarPostBody = r.settingRequest
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