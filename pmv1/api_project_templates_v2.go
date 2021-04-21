/*
 * Teamwork Projects API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmv1

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

// ProjectTemplatesV2ApiService ProjectTemplatesV2Api service
type ProjectTemplatesV2ApiService service

type ApiDELETEProjectsTemplateTemplateIDJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectTemplatesV2ApiService
	templateID string
}


func (r ApiDELETEProjectsTemplateTemplateIDJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.DELETEProjectsTemplateTemplateIDJsonExecute(r)
}

/*
 * DELETEProjectsTemplateTemplateIDJson Delete a Project Template
 * **NOTE: This is a V2 endpoint. These endpoints can be used but are subject to change.**

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param templateID
 * @return ApiDELETEProjectsTemplateTemplateIDJsonRequest
 */
func (a *ProjectTemplatesV2ApiService) DELETEProjectsTemplateTemplateIDJson(ctx _context.Context, templateID string) ApiDELETEProjectsTemplateTemplateIDJsonRequest {
	return ApiDELETEProjectsTemplateTemplateIDJsonRequest{
		ApiService: a,
		ctx: ctx,
		templateID: templateID,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *ProjectTemplatesV2ApiService) DELETEProjectsTemplateTemplateIDJsonExecute(r ApiDELETEProjectsTemplateTemplateIDJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectTemplatesV2ApiService.DELETEProjectsTemplateTemplateIDJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/template/{TemplateID}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"TemplateID"+"}", _neturl.PathEscape(parameterToString(r.templateID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGETProjectsApiV2ProjectsTemplateIdJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectTemplatesV2ApiService
	templateId string
	getpermissions bool
}


func (r ApiGETProjectsApiV2ProjectsTemplateIdJsonRequest) Execute() (InlineResponse20054, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV2ProjectsTemplateIdJsonExecute(r)
}

/*
 * GETProjectsApiV2ProjectsTemplateIdJson Get a Project Template
 * **NOTE: This is a V2 endpoint. These endpoints can be used but are subject to change.**

Returns a specific project template by passing in the ID. 

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param templateId
 * @param getpermissions If this a required value in order to return the template. Values = true/false
 * @return ApiGETProjectsApiV2ProjectsTemplateIdJsonRequest
 */
func (a *ProjectTemplatesV2ApiService) GETProjectsApiV2ProjectsTemplateIdJson(ctx _context.Context, templateId string, getpermissions bool) ApiGETProjectsApiV2ProjectsTemplateIdJsonRequest {
	return ApiGETProjectsApiV2ProjectsTemplateIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
		getpermissions: getpermissions,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20054
 */
func (a *ProjectTemplatesV2ApiService) GETProjectsApiV2ProjectsTemplateIdJsonExecute(r ApiGETProjectsApiV2ProjectsTemplateIdJsonRequest) (InlineResponse20054, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20054
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectTemplatesV2ApiService.GETProjectsApiV2ProjectsTemplateIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v2/projects/{templateId}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", _neturl.PathEscape(parameterToString(r.templateId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"getpermissions"+"}", _neturl.PathEscape(parameterToString(r.getpermissions, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGETProjectsApiV2ProjectsTemplatesJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectTemplatesV2ApiService
	searchTerm *string
	orderBy *string
	projectOwnerIds *string
	projectTagIds *string
}

func (r ApiGETProjectsApiV2ProjectsTemplatesJsonRequest) SearchTerm(searchTerm string) ApiGETProjectsApiV2ProjectsTemplatesJsonRequest {
	r.searchTerm = &searchTerm
	return r
}
func (r ApiGETProjectsApiV2ProjectsTemplatesJsonRequest) OrderBy(orderBy string) ApiGETProjectsApiV2ProjectsTemplatesJsonRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGETProjectsApiV2ProjectsTemplatesJsonRequest) ProjectOwnerIds(projectOwnerIds string) ApiGETProjectsApiV2ProjectsTemplatesJsonRequest {
	r.projectOwnerIds = &projectOwnerIds
	return r
}
func (r ApiGETProjectsApiV2ProjectsTemplatesJsonRequest) ProjectTagIds(projectTagIds string) ApiGETProjectsApiV2ProjectsTemplatesJsonRequest {
	r.projectTagIds = &projectTagIds
	return r
}

func (r ApiGETProjectsApiV2ProjectsTemplatesJsonRequest) Execute() (InlineResponse20054, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV2ProjectsTemplatesJsonExecute(r)
}

/*
 * GETProjectsApiV2ProjectsTemplatesJson Get Project Templates
 * **NOTE: This is a V2 endpoint. These endpoints can be used but are subject to change.**

Returns all project templates which were set up on the account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV2ProjectsTemplatesJsonRequest
 */
func (a *ProjectTemplatesV2ApiService) GETProjectsApiV2ProjectsTemplatesJson(ctx _context.Context) ApiGETProjectsApiV2ProjectsTemplatesJsonRequest {
	return ApiGETProjectsApiV2ProjectsTemplatesJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20054
 */
func (a *ProjectTemplatesV2ApiService) GETProjectsApiV2ProjectsTemplatesJsonExecute(r ApiGETProjectsApiV2ProjectsTemplatesJsonRequest) (InlineResponse20054, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20054
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectTemplatesV2ApiService.GETProjectsApiV2ProjectsTemplatesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v2/projects/templates.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.searchTerm != nil {
		localVarQueryParams.Add("searchTerm", parameterToString(*r.searchTerm, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.projectOwnerIds != nil {
		localVarQueryParams.Add("projectOwnerIds", parameterToString(*r.projectOwnerIds, ""))
	}
	if r.projectTagIds != nil {
		localVarQueryParams.Add("projectTagIds", parameterToString(*r.projectTagIds, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiPOSTProjectsTemplateJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectTemplatesV2ApiService
	body *InlineObject51
}

func (r ApiPOSTProjectsTemplateJsonRequest) Body(body InlineObject51) ApiPOSTProjectsTemplateJsonRequest {
	r.body = &body
	return r
}

func (r ApiPOSTProjectsTemplateJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.POSTProjectsTemplateJsonExecute(r)
}

/*
 * POSTProjectsTemplateJson Create a Project Template
 * **NOTE: This is a V2 endpoint. These endpoints can be used but are subject to change.**

Creating a project template by passing in a json payload 'project' and the required parameters. 

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPOSTProjectsTemplateJsonRequest
 */
func (a *ProjectTemplatesV2ApiService) POSTProjectsTemplateJson(ctx _context.Context) ApiPOSTProjectsTemplateJsonRequest {
	return ApiPOSTProjectsTemplateJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *ProjectTemplatesV2ApiService) POSTProjectsTemplateJsonExecute(r ApiPOSTProjectsTemplateJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectTemplatesV2ApiService.POSTProjectsTemplateJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/template.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.body
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

type ApiPUTProjectsTemplateTemplateIDJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectTemplatesV2ApiService
	templateID string
	body *InlineObject52
}

func (r ApiPUTProjectsTemplateTemplateIDJsonRequest) Body(body InlineObject52) ApiPUTProjectsTemplateTemplateIDJsonRequest {
	r.body = &body
	return r
}

func (r ApiPUTProjectsTemplateTemplateIDJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.PUTProjectsTemplateTemplateIDJsonExecute(r)
}

/*
 * PUTProjectsTemplateTemplateIDJson Update a specific Project Template
 * **NOTE: This is a V2 endpoint. These endpoints can be used but are subject to change.**

Update a specific template using the ID in the URL.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param templateID
 * @return ApiPUTProjectsTemplateTemplateIDJsonRequest
 */
func (a *ProjectTemplatesV2ApiService) PUTProjectsTemplateTemplateIDJson(ctx _context.Context, templateID string) ApiPUTProjectsTemplateTemplateIDJsonRequest {
	return ApiPUTProjectsTemplateTemplateIDJsonRequest{
		ApiService: a,
		ctx: ctx,
		templateID: templateID,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *ProjectTemplatesV2ApiService) PUTProjectsTemplateTemplateIDJsonExecute(r ApiPUTProjectsTemplateTemplateIDJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectTemplatesV2ApiService.PUTProjectsTemplateTemplateIDJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/template/{TemplateID}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"TemplateID"+"}", _neturl.PathEscape(parameterToString(r.templateID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = r.body
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
