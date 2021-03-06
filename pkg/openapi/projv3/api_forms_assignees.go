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

// FormsAssigneesApiService FormsAssigneesApi service
type FormsAssigneesApiService service

type ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest struct {
	ctx _context.Context
	ApiService *FormsAssigneesApiService
	formId int32
	orderMode *string
	userId *int32
	pageSize *int32
	page *int32
	include *[]string
	fieldsUsers *[]string
	fieldsTeams *[]string
	fieldsFormAssignees *[]string
	fieldsCompanies *[]string
}

func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) OrderMode(orderMode string) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.orderMode = &orderMode
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) UserId(userId int32) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.userId = &userId
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) PageSize(pageSize int32) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) Page(page int32) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) Include(include []string) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.include = &include
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) FieldsUsers(fieldsUsers []string) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.fieldsUsers = &fieldsUsers
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) FieldsTeams(fieldsTeams []string) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.fieldsTeams = &fieldsTeams
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) FieldsFormAssignees(fieldsFormAssignees []string) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.fieldsFormAssignees = &fieldsFormAssignees
	return r
}
func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) FieldsCompanies(fieldsCompanies []string) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	r.fieldsCompanies = &fieldsCompanies
	return r
}

func (r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) Execute() (AssigneeFormAssigneesResponse, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3FormsformIdAssigneesJsonExecute(r)
}

/*
 * GETProjectsApiV3FormsformIdAssigneesJson Get all assignees for a given form.
 * Get all assignees for a given form.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param formId
 * @return ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest
 */
func (a *FormsAssigneesApiService) GETProjectsApiV3FormsformIdAssigneesJson(ctx _context.Context, formId int32) ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest {
	return ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest{
		ApiService: a,
		ctx: ctx,
		formId: formId,
	}
}

/*
 * Execute executes the request
 * @return AssigneeFormAssigneesResponse
 */
func (a *FormsAssigneesApiService) GETProjectsApiV3FormsformIdAssigneesJsonExecute(r ApiGETProjectsApiV3FormsformIdAssigneesJsonRequest) (AssigneeFormAssigneesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AssigneeFormAssigneesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormsAssigneesApiService.GETProjectsApiV3FormsformIdAssigneesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/forms/{formId}/assignees.json"
	localVarPath = strings.Replace(localVarPath, "{"+"formId"+"}", _neturl.PathEscape(parameterToString(r.formId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.orderMode != nil {
		localVarQueryParams.Add("orderMode", parameterToString(*r.orderMode, ""))
	}
	if r.userId != nil {
		localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.fieldsUsers != nil {
		localVarQueryParams.Add("fields[users]", parameterToString(*r.fieldsUsers, "csv"))
	}
	if r.fieldsTeams != nil {
		localVarQueryParams.Add("fields[teams]", parameterToString(*r.fieldsTeams, "csv"))
	}
	if r.fieldsFormAssignees != nil {
		localVarQueryParams.Add("fields[formAssignees]", parameterToString(*r.fieldsFormAssignees, "csv"))
	}
	if r.fieldsCompanies != nil {
		localVarQueryParams.Add("fields[companies]", parameterToString(*r.fieldsCompanies, "csv"))
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

type ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest struct {
	ctx _context.Context
	ApiService *FormsAssigneesApiService
	formId int32
	assigneeRequest *AssigneeRequest
}

func (r ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest) AssigneeRequest(assigneeRequest AssigneeRequest) ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest {
	r.assigneeRequest = &assigneeRequest
	return r
}

func (r ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest) Execute() (AssigneeResponse, *_nethttp.Response, error) {
	return r.ApiService.PUTProjectsApiV3FormformIdAssigneesJsonExecute(r)
}

/*
 * PUTProjectsApiV3FormformIdAssigneesJson Update the existing assignees.
 * Replace the current assignees. An empty list deletes.

If some of the assignees are found to error, the correct
ones will be created, and the incorrect assignees will be
reported back in the response, under `.errors.[]`

If every assignee is found to not exist, the response will be
a 404.

If every assignee is found to be invalid, or a mix of invalid
and doesn't exist, the response will be a 400.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param formId
 * @return ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest
 */
func (a *FormsAssigneesApiService) PUTProjectsApiV3FormformIdAssigneesJson(ctx _context.Context, formId int32) ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest {
	return ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest{
		ApiService: a,
		ctx: ctx,
		formId: formId,
	}
}

/*
 * Execute executes the request
 * @return AssigneeResponse
 */
func (a *FormsAssigneesApiService) PUTProjectsApiV3FormformIdAssigneesJsonExecute(r ApiPUTProjectsApiV3FormformIdAssigneesJsonRequest) (AssigneeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AssigneeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormsAssigneesApiService.PUTProjectsApiV3FormformIdAssigneesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/form/{formId}/assignees.json"
	localVarPath = strings.Replace(localVarPath, "{"+"formId"+"}", _neturl.PathEscape(parameterToString(r.formId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.assigneeRequest == nil {
		return localVarReturnValue, nil, reportError("assigneeRequest is required and must be specified")
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
	localVarPostBody = r.assigneeRequest
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
