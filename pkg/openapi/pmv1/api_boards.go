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

// BoardsApiService BoardsApi service
type BoardsApiService service

type ApiDELETEBoardsColumnsCardsIdJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id int32
}


func (r ApiDELETEBoardsColumnsCardsIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.DELETEBoardsColumnsCardsIdJsonExecute(r)
}

/*
 * DELETEBoardsColumnsCardsIdJson Remove a Card
 * Remove a Card from a Column (that task isn't deleted from the Tasklist).

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiDELETEBoardsColumnsCardsIdJsonRequest
 */
func (a *BoardsApiService) DELETEBoardsColumnsCardsIdJson(ctx _context.Context, id int32) ApiDELETEBoardsColumnsCardsIdJsonRequest {
	return ApiDELETEBoardsColumnsCardsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *BoardsApiService) DELETEBoardsColumnsCardsIdJsonExecute(r ApiDELETEBoardsColumnsCardsIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.DELETEBoardsColumnsCardsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/boards/columns/cards/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiDELETEBoardsColumnsIdJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id int32
}


func (r ApiDELETEBoardsColumnsIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.DELETEBoardsColumnsIdJsonExecute(r)
}

/*
 * DELETEBoardsColumnsIdJson Delete a Column
 * Remove a Column from your boards view.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiDELETEBoardsColumnsIdJsonRequest
 */
func (a *BoardsApiService) DELETEBoardsColumnsIdJson(ctx _context.Context, id int32) ApiDELETEBoardsColumnsIdJsonRequest {
	return ApiDELETEBoardsColumnsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *BoardsApiService) DELETEBoardsColumnsIdJsonExecute(r ApiDELETEBoardsColumnsIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.DELETEBoardsColumnsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/boards/columns/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiGETBoardsColumnsIdCardsJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id string
	page *string
	pageSize *int32
	showDeleted *bool
	deletedAfterDate *string
	updatedAfterDate *string
	searchTerm *string
}

func (r ApiGETBoardsColumnsIdCardsJsonRequest) Page(page string) ApiGETBoardsColumnsIdCardsJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETBoardsColumnsIdCardsJsonRequest) PageSize(pageSize int32) ApiGETBoardsColumnsIdCardsJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETBoardsColumnsIdCardsJsonRequest) ShowDeleted(showDeleted bool) ApiGETBoardsColumnsIdCardsJsonRequest {
	r.showDeleted = &showDeleted
	return r
}
func (r ApiGETBoardsColumnsIdCardsJsonRequest) DeletedAfterDate(deletedAfterDate string) ApiGETBoardsColumnsIdCardsJsonRequest {
	r.deletedAfterDate = &deletedAfterDate
	return r
}
func (r ApiGETBoardsColumnsIdCardsJsonRequest) UpdatedAfterDate(updatedAfterDate string) ApiGETBoardsColumnsIdCardsJsonRequest {
	r.updatedAfterDate = &updatedAfterDate
	return r
}
func (r ApiGETBoardsColumnsIdCardsJsonRequest) SearchTerm(searchTerm string) ApiGETBoardsColumnsIdCardsJsonRequest {
	r.searchTerm = &searchTerm
	return r
}

func (r ApiGETBoardsColumnsIdCardsJsonRequest) Execute() (InlineResponse2002, *_nethttp.Response, error) {
	return r.ApiService.GETBoardsColumnsIdCardsJsonExecute(r)
}

/*
 * GETBoardsColumnsIdCardsJson List Cards in a Column
 * List the existing Cards inside a Column.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiGETBoardsColumnsIdCardsJsonRequest
 */
func (a *BoardsApiService) GETBoardsColumnsIdCardsJson(ctx _context.Context, id string) ApiGETBoardsColumnsIdCardsJsonRequest {
	return ApiGETBoardsColumnsIdCardsJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2002
 */
func (a *BoardsApiService) GETBoardsColumnsIdCardsJsonExecute(r ApiGETBoardsColumnsIdCardsJsonRequest) (InlineResponse2002, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2002
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.GETBoardsColumnsIdCardsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/boards/columns/{id}/cards.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.showDeleted != nil {
		localVarQueryParams.Add("showDeleted", parameterToString(*r.showDeleted, ""))
	}
	if r.deletedAfterDate != nil {
		localVarQueryParams.Add("deletedAfterDate", parameterToString(*r.deletedAfterDate, ""))
	}
	if r.updatedAfterDate != nil {
		localVarQueryParams.Add("updatedAfterDate", parameterToString(*r.updatedAfterDate, ""))
	}
	if r.searchTerm != nil {
		localVarQueryParams.Add("searchTerm", parameterToString(*r.searchTerm, ""))
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

type ApiGETProjectsIdBoardsColumnsJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id int32
	page *int32
	pageSize *int32
	showDeleted *bool
	deletedAfterDate *string
	updatedAfterDate *string
}

func (r ApiGETProjectsIdBoardsColumnsJsonRequest) Page(page int32) ApiGETProjectsIdBoardsColumnsJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsIdBoardsColumnsJsonRequest) PageSize(pageSize int32) ApiGETProjectsIdBoardsColumnsJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsIdBoardsColumnsJsonRequest) ShowDeleted(showDeleted bool) ApiGETProjectsIdBoardsColumnsJsonRequest {
	r.showDeleted = &showDeleted
	return r
}
func (r ApiGETProjectsIdBoardsColumnsJsonRequest) DeletedAfterDate(deletedAfterDate string) ApiGETProjectsIdBoardsColumnsJsonRequest {
	r.deletedAfterDate = &deletedAfterDate
	return r
}
func (r ApiGETProjectsIdBoardsColumnsJsonRequest) UpdatedAfterDate(updatedAfterDate string) ApiGETProjectsIdBoardsColumnsJsonRequest {
	r.updatedAfterDate = &updatedAfterDate
	return r
}

func (r ApiGETProjectsIdBoardsColumnsJsonRequest) Execute() (InlineResponse20061, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsIdBoardsColumnsJsonExecute(r)
}

/*
 * GETProjectsIdBoardsColumnsJson List Columns
 * Allows you to list all existing Columns in your Boards view.

Example:
https://yoursite.teamwork.com/projects/{id}/boards/columns.json?showDeleted=true&updatedAfterDate=20171119120101

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiGETProjectsIdBoardsColumnsJsonRequest
 */
func (a *BoardsApiService) GETProjectsIdBoardsColumnsJson(ctx _context.Context, id int32) ApiGETProjectsIdBoardsColumnsJsonRequest {
	return ApiGETProjectsIdBoardsColumnsJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse20061
 */
func (a *BoardsApiService) GETProjectsIdBoardsColumnsJsonExecute(r ApiGETProjectsIdBoardsColumnsJsonRequest) (InlineResponse20061, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20061
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.GETProjectsIdBoardsColumnsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{id}/boards/columns.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.showDeleted != nil {
		localVarQueryParams.Add("showDeleted", parameterToString(*r.showDeleted, ""))
	}
	if r.deletedAfterDate != nil {
		localVarQueryParams.Add("deletedAfterDate", parameterToString(*r.deletedAfterDate, ""))
	}
	if r.updatedAfterDate != nil {
		localVarQueryParams.Add("updatedAfterDate", parameterToString(*r.updatedAfterDate, ""))
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

type ApiPOSTBoardsColumnsIdCardsJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id string
	body *InlineObject3
}

func (r ApiPOSTBoardsColumnsIdCardsJsonRequest) Body(body InlineObject3) ApiPOSTBoardsColumnsIdCardsJsonRequest {
	r.body = &body
	return r
}

func (r ApiPOSTBoardsColumnsIdCardsJsonRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.POSTBoardsColumnsIdCardsJsonExecute(r)
}

/*
 * POSTBoardsColumnsIdCardsJson Add a Task from the Backlog to a Column
 * Add a task from your list of Tasks in to a Column.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPOSTBoardsColumnsIdCardsJsonRequest
 */
func (a *BoardsApiService) POSTBoardsColumnsIdCardsJson(ctx _context.Context, id string) ApiPOSTBoardsColumnsIdCardsJsonRequest {
	return ApiPOSTBoardsColumnsIdCardsJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *BoardsApiService) POSTBoardsColumnsIdCardsJsonExecute(r ApiPOSTBoardsColumnsIdCardsJsonRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.POSTBoardsColumnsIdCardsJson")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/boards/columns/{id}/cards.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPOSTProjectsIdBoardsColumnsJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id int32
	body *InlineObject56
}

func (r ApiPOSTProjectsIdBoardsColumnsJsonRequest) Body(body InlineObject56) ApiPOSTProjectsIdBoardsColumnsJsonRequest {
	r.body = &body
	return r
}

func (r ApiPOSTProjectsIdBoardsColumnsJsonRequest) Execute() (InlineResponse201, *_nethttp.Response, error) {
	return r.ApiService.POSTProjectsIdBoardsColumnsJsonExecute(r)
}

/*
 * POSTProjectsIdBoardsColumnsJson Create a new Column
 * Allows you to create a new Column in your Boards view.

<h4>Please note: The following colours can be used when creating Columns:</h4>
#27AE60, #99DF72, #1ABC9C, #6866D0, #8E44AD, #0AD2F5, #3498DB, #3D82DE, #C0392B, #E74C3C, #A94136, #660A00, #F39C12, #F1C40F, #34495E, #7F8C8D, #D35400, #B49255, #D870AD, #BDC3C7, #9B59B6

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPOSTProjectsIdBoardsColumnsJsonRequest
 */
func (a *BoardsApiService) POSTProjectsIdBoardsColumnsJson(ctx _context.Context, id int32) ApiPOSTProjectsIdBoardsColumnsJsonRequest {
	return ApiPOSTProjectsIdBoardsColumnsJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse201
 */
func (a *BoardsApiService) POSTProjectsIdBoardsColumnsJsonExecute(r ApiPOSTProjectsIdBoardsColumnsJsonRequest) (InlineResponse201, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse201
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.POSTProjectsIdBoardsColumnsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/{id}/boards/columns.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiPUTBoardsColumnsCardsIdJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id int32
	body *InlineObject1
}

func (r ApiPUTBoardsColumnsCardsIdJsonRequest) Body(body InlineObject1) ApiPUTBoardsColumnsCardsIdJsonRequest {
	r.body = &body
	return r
}

func (r ApiPUTBoardsColumnsCardsIdJsonRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.PUTBoardsColumnsCardsIdJsonExecute(r)
}

/*
 * PUTBoardsColumnsCardsIdJson Edit a Card
 * Edit the details of a Card.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPUTBoardsColumnsCardsIdJsonRequest
 */
func (a *BoardsApiService) PUTBoardsColumnsCardsIdJson(ctx _context.Context, id int32) ApiPUTBoardsColumnsCardsIdJsonRequest {
	return ApiPUTBoardsColumnsCardsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *BoardsApiService) PUTBoardsColumnsCardsIdJsonExecute(r ApiPUTBoardsColumnsCardsIdJsonRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.PUTBoardsColumnsCardsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/boards/columns/cards/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiPUTBoardsColumnsCardsIdMoveJsonRequest struct {
	ctx _context.Context
	ApiService *BoardsApiService
	id int32
	body *InlineObject2
}

func (r ApiPUTBoardsColumnsCardsIdMoveJsonRequest) Body(body InlineObject2) ApiPUTBoardsColumnsCardsIdMoveJsonRequest {
	r.body = &body
	return r
}

func (r ApiPUTBoardsColumnsCardsIdMoveJsonRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.PUTBoardsColumnsCardsIdMoveJsonExecute(r)
}

/*
 * PUTBoardsColumnsCardsIdMoveJson Move a Card
 * Add a task from your list of Tasks in to a Column.

---
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiPUTBoardsColumnsCardsIdMoveJsonRequest
 */
func (a *BoardsApiService) PUTBoardsColumnsCardsIdMoveJson(ctx _context.Context, id int32) ApiPUTBoardsColumnsCardsIdMoveJsonRequest {
	return ApiPUTBoardsColumnsCardsIdMoveJsonRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *BoardsApiService) PUTBoardsColumnsCardsIdMoveJsonExecute(r ApiPUTBoardsColumnsCardsIdMoveJsonRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardsApiService.PUTBoardsColumnsCardsIdMoveJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/boards/columns/cards/{id}/move.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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
