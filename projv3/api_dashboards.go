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
)

// Linger please
var (
	_ _context.Context
)

// DashboardsApiService DashboardsApi service
type DashboardsApiService service

type ApiGETProjectsApiV3DashboardsJsonRequest struct {
	ctx _context.Context
	ApiService *DashboardsApiService
	updatedAfter *string
	orderMode *string
	orderBy *string
	userId *int32
	projectId *int32
	pageSize *int32
	page *int32
	showDeleted *bool
	emoji *bool
	include *[]string
	fieldsUsers *[]string
	fieldsProjects *[]string
	fieldsDashboards *[]string
	fieldsDashboardSettings *[]string
	fieldsDashboardPanels *[]string
	fieldsDashboardPanelSettings *[]string
	dashboardIds *[]int32
}

func (r ApiGETProjectsApiV3DashboardsJsonRequest) UpdatedAfter(updatedAfter string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.updatedAfter = &updatedAfter
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) OrderMode(orderMode string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.orderMode = &orderMode
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) OrderBy(orderBy string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) UserId(userId int32) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.userId = &userId
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) ProjectId(projectId int32) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.projectId = &projectId
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) PageSize(pageSize int32) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) Page(page int32) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) ShowDeleted(showDeleted bool) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.showDeleted = &showDeleted
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) Emoji(emoji bool) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.emoji = &emoji
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) Include(include []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.include = &include
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) FieldsUsers(fieldsUsers []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.fieldsUsers = &fieldsUsers
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) FieldsProjects(fieldsProjects []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.fieldsProjects = &fieldsProjects
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) FieldsDashboards(fieldsDashboards []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.fieldsDashboards = &fieldsDashboards
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) FieldsDashboardSettings(fieldsDashboardSettings []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.fieldsDashboardSettings = &fieldsDashboardSettings
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) FieldsDashboardPanels(fieldsDashboardPanels []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.fieldsDashboardPanels = &fieldsDashboardPanels
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) FieldsDashboardPanelSettings(fieldsDashboardPanelSettings []string) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.fieldsDashboardPanelSettings = &fieldsDashboardPanelSettings
	return r
}
func (r ApiGETProjectsApiV3DashboardsJsonRequest) DashboardIds(dashboardIds []int32) ApiGETProjectsApiV3DashboardsJsonRequest {
	r.dashboardIds = &dashboardIds
	return r
}

func (r ApiGETProjectsApiV3DashboardsJsonRequest) Execute() (DashboardUserDashboardsResponse, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3DashboardsJsonExecute(r)
}

/*
 * GETProjectsApiV3DashboardsJson Get all dashboards
 * Return all dashboards that the logged-in user can access. Dashboards contains
panels that are used to summarize many activities from the installation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3DashboardsJsonRequest
 */
func (a *DashboardsApiService) GETProjectsApiV3DashboardsJson(ctx _context.Context) ApiGETProjectsApiV3DashboardsJsonRequest {
	return ApiGETProjectsApiV3DashboardsJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return DashboardUserDashboardsResponse
 */
func (a *DashboardsApiService) GETProjectsApiV3DashboardsJsonExecute(r ApiGETProjectsApiV3DashboardsJsonRequest) (DashboardUserDashboardsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DashboardUserDashboardsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardsApiService.GETProjectsApiV3DashboardsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/dashboards.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.updatedAfter != nil {
		localVarQueryParams.Add("updatedAfter", parameterToString(*r.updatedAfter, ""))
	}
	if r.orderMode != nil {
		localVarQueryParams.Add("orderMode", parameterToString(*r.orderMode, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.userId != nil {
		localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.showDeleted != nil {
		localVarQueryParams.Add("showDeleted", parameterToString(*r.showDeleted, ""))
	}
	if r.emoji != nil {
		localVarQueryParams.Add("emoji", parameterToString(*r.emoji, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.fieldsUsers != nil {
		localVarQueryParams.Add("fields[users]", parameterToString(*r.fieldsUsers, "csv"))
	}
	if r.fieldsProjects != nil {
		localVarQueryParams.Add("fields[projects]", parameterToString(*r.fieldsProjects, "csv"))
	}
	if r.fieldsDashboards != nil {
		localVarQueryParams.Add("fields[dashboards]", parameterToString(*r.fieldsDashboards, "csv"))
	}
	if r.fieldsDashboardSettings != nil {
		localVarQueryParams.Add("fields[dashboardSettings]", parameterToString(*r.fieldsDashboardSettings, "csv"))
	}
	if r.fieldsDashboardPanels != nil {
		localVarQueryParams.Add("fields[dashboardPanels]", parameterToString(*r.fieldsDashboardPanels, "csv"))
	}
	if r.fieldsDashboardPanelSettings != nil {
		localVarQueryParams.Add("fields[dashboardPanelSettings]", parameterToString(*r.fieldsDashboardPanelSettings, "csv"))
	}
	if r.dashboardIds != nil {
		localVarQueryParams.Add("dashboardIds", parameterToString(*r.dashboardIds, "csv"))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ViewErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
