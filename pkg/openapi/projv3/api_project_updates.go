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
	"time"
)

// Linger please
var (
	_ _context.Context
)

// ProjectUpdatesApiService ProjectUpdatesApi service
type ProjectUpdatesApiService service

type ApiGETProjectsApiV3ProjectsUpdatesJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectUpdatesApiService
	projectStatuses *string
	projectStatus *string
	orderMode *string
	orderBy *string
	createdAfter *time.Time
	projectId *int32
	projectHealths *int32
	projectHealth *int32
	pageSize *int32
	page *int32
	showDeleted *bool
	reactions *bool
	onlyStarredProjects *bool
	matchAllProjectTags *bool
	includeArchivedProjects *bool
	emoji *bool
	activeOnly *bool
	projectTagIds *[]int32
	projectOwnerIds *[]int32
	projectCompanyIds *[]int32
	projectCategoryIds *[]int32
	include *[]string
	fieldsUsers *[]string
	fieldsProjects *[]string
	fieldsProjectUpdates *[]string
}

func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectStatuses(projectStatuses string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectStatuses = &projectStatuses
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectStatus(projectStatus string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectStatus = &projectStatus
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) OrderMode(orderMode string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.orderMode = &orderMode
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) OrderBy(orderBy string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) CreatedAfter(createdAfter time.Time) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectId(projectId int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectId = &projectId
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectHealths(projectHealths int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectHealths = &projectHealths
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectHealth(projectHealth int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectHealth = &projectHealth
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) PageSize(pageSize int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) Page(page int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ShowDeleted(showDeleted bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.showDeleted = &showDeleted
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) Reactions(reactions bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.reactions = &reactions
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) OnlyStarredProjects(onlyStarredProjects bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.onlyStarredProjects = &onlyStarredProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) MatchAllProjectTags(matchAllProjectTags bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.matchAllProjectTags = &matchAllProjectTags
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) IncludeArchivedProjects(includeArchivedProjects bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.includeArchivedProjects = &includeArchivedProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) Emoji(emoji bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.emoji = &emoji
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ActiveOnly(activeOnly bool) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.activeOnly = &activeOnly
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectTagIds(projectTagIds []int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectTagIds = &projectTagIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectOwnerIds(projectOwnerIds []int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectOwnerIds = &projectOwnerIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectCompanyIds(projectCompanyIds []int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectCompanyIds = &projectCompanyIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) ProjectCategoryIds(projectCategoryIds []int32) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.projectCategoryIds = &projectCategoryIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) Include(include []string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.include = &include
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) FieldsUsers(fieldsUsers []string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.fieldsUsers = &fieldsUsers
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) FieldsProjects(fieldsProjects []string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.fieldsProjects = &fieldsProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) FieldsProjectUpdates(fieldsProjectUpdates []string) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	r.fieldsProjectUpdates = &fieldsProjectUpdates
	return r
}

func (r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) Execute() (UpdateProjectUpdatesResponse, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3ProjectsUpdatesJsonExecute(r)
}

/*
 * GETProjectsApiV3ProjectsUpdatesJson Get all project updates
 * Return the updates from all projects that the logged-in user can access.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3ProjectsUpdatesJsonRequest
 */
func (a *ProjectUpdatesApiService) GETProjectsApiV3ProjectsUpdatesJson(ctx _context.Context) ApiGETProjectsApiV3ProjectsUpdatesJsonRequest {
	return ApiGETProjectsApiV3ProjectsUpdatesJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return UpdateProjectUpdatesResponse
 */
func (a *ProjectUpdatesApiService) GETProjectsApiV3ProjectsUpdatesJsonExecute(r ApiGETProjectsApiV3ProjectsUpdatesJsonRequest) (UpdateProjectUpdatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateProjectUpdatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectUpdatesApiService.GETProjectsApiV3ProjectsUpdatesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/projects/updates.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.projectStatuses != nil {
		localVarQueryParams.Add("projectStatuses", parameterToString(*r.projectStatuses, ""))
	}
	if r.projectStatus != nil {
		localVarQueryParams.Add("projectStatus", parameterToString(*r.projectStatus, ""))
	}
	if r.orderMode != nil {
		localVarQueryParams.Add("orderMode", parameterToString(*r.orderMode, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("createdAfter", parameterToString(*r.createdAfter, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.projectHealths != nil {
		localVarQueryParams.Add("projectHealths", parameterToString(*r.projectHealths, ""))
	}
	if r.projectHealth != nil {
		localVarQueryParams.Add("projectHealth", parameterToString(*r.projectHealth, ""))
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
	if r.reactions != nil {
		localVarQueryParams.Add("reactions", parameterToString(*r.reactions, ""))
	}
	if r.onlyStarredProjects != nil {
		localVarQueryParams.Add("onlyStarredProjects", parameterToString(*r.onlyStarredProjects, ""))
	}
	if r.matchAllProjectTags != nil {
		localVarQueryParams.Add("matchAllProjectTags", parameterToString(*r.matchAllProjectTags, ""))
	}
	if r.includeArchivedProjects != nil {
		localVarQueryParams.Add("includeArchivedProjects", parameterToString(*r.includeArchivedProjects, ""))
	}
	if r.emoji != nil {
		localVarQueryParams.Add("emoji", parameterToString(*r.emoji, ""))
	}
	if r.activeOnly != nil {
		localVarQueryParams.Add("activeOnly", parameterToString(*r.activeOnly, ""))
	}
	if r.projectTagIds != nil {
		localVarQueryParams.Add("projectTagIds", parameterToString(*r.projectTagIds, "csv"))
	}
	if r.projectOwnerIds != nil {
		localVarQueryParams.Add("projectOwnerIds", parameterToString(*r.projectOwnerIds, "csv"))
	}
	if r.projectCompanyIds != nil {
		localVarQueryParams.Add("projectCompanyIds", parameterToString(*r.projectCompanyIds, "csv"))
	}
	if r.projectCategoryIds != nil {
		localVarQueryParams.Add("projectCategoryIds", parameterToString(*r.projectCategoryIds, "csv"))
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
	if r.fieldsProjectUpdates != nil {
		localVarQueryParams.Add("fields[projectUpdates]", parameterToString(*r.fieldsProjectUpdates, "csv"))
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

type ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectUpdatesApiService
	projectIds int32
	projectStatuses *string
	projectStatus *string
	orderMode *string
	orderBy *string
	createdAfter *time.Time
	projectId *int32
	projectHealths *int32
	projectHealth *int32
	pageSize *int32
	page *int32
	showDeleted *bool
	reactions *bool
	onlyStarredProjects *bool
	matchAllProjectTags *bool
	includeArchivedProjects *bool
	emoji *bool
	activeOnly *bool
	projectTagIds *[]int32
	projectOwnerIds *[]int32
	projectCompanyIds *[]int32
	projectCategoryIds *[]int32
	include *[]string
	fieldsUsers *[]string
	fieldsProjects *[]string
	fieldsProjectUpdates *[]string
}

func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectStatuses(projectStatuses string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectStatuses = &projectStatuses
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectStatus(projectStatus string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectStatus = &projectStatus
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) OrderMode(orderMode string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.orderMode = &orderMode
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) OrderBy(orderBy string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) CreatedAfter(createdAfter time.Time) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectId(projectId int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectId = &projectId
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectHealths(projectHealths int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectHealths = &projectHealths
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectHealth(projectHealth int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectHealth = &projectHealth
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) PageSize(pageSize int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) Page(page int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ShowDeleted(showDeleted bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.showDeleted = &showDeleted
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) Reactions(reactions bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.reactions = &reactions
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) OnlyStarredProjects(onlyStarredProjects bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.onlyStarredProjects = &onlyStarredProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) MatchAllProjectTags(matchAllProjectTags bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.matchAllProjectTags = &matchAllProjectTags
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) IncludeArchivedProjects(includeArchivedProjects bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.includeArchivedProjects = &includeArchivedProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) Emoji(emoji bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.emoji = &emoji
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ActiveOnly(activeOnly bool) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.activeOnly = &activeOnly
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectTagIds(projectTagIds []int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectTagIds = &projectTagIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectOwnerIds(projectOwnerIds []int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectOwnerIds = &projectOwnerIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectCompanyIds(projectCompanyIds []int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectCompanyIds = &projectCompanyIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) ProjectCategoryIds(projectCategoryIds []int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.projectCategoryIds = &projectCategoryIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) Include(include []string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.include = &include
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) FieldsUsers(fieldsUsers []string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.fieldsUsers = &fieldsUsers
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) FieldsProjects(fieldsProjects []string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.fieldsProjects = &fieldsProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) FieldsProjectUpdates(fieldsProjectUpdates []string) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	r.fieldsProjectUpdates = &fieldsProjectUpdates
	return r
}

func (r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) Execute() (UpdateProjectUpdatesResponse, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3ProjectsprojectIdsUpdatesJsonExecute(r)
}

/*
 * GETProjectsApiV3ProjectsprojectIdsUpdatesJson Get all updates from a specific project
 * Return the updates from a specific project.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectIds
 * @return ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest
 */
func (a *ProjectUpdatesApiService) GETProjectsApiV3ProjectsprojectIdsUpdatesJson(ctx _context.Context, projectIds int32) ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest {
	return ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest{
		ApiService: a,
		ctx: ctx,
		projectIds: projectIds,
	}
}

/*
 * Execute executes the request
 * @return UpdateProjectUpdatesResponse
 */
func (a *ProjectUpdatesApiService) GETProjectsApiV3ProjectsprojectIdsUpdatesJsonExecute(r ApiGETProjectsApiV3ProjectsprojectIdsUpdatesJsonRequest) (UpdateProjectUpdatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateProjectUpdatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectUpdatesApiService.GETProjectsApiV3ProjectsprojectIdsUpdatesJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/projects/{projectIds}/updates.json"
	localVarPath = strings.Replace(localVarPath, "{"+"projectIds"+"}", _neturl.PathEscape(parameterToString(r.projectIds, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.projectStatuses != nil {
		localVarQueryParams.Add("projectStatuses", parameterToString(*r.projectStatuses, ""))
	}
	if r.projectStatus != nil {
		localVarQueryParams.Add("projectStatus", parameterToString(*r.projectStatus, ""))
	}
	if r.orderMode != nil {
		localVarQueryParams.Add("orderMode", parameterToString(*r.orderMode, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("createdAfter", parameterToString(*r.createdAfter, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.projectHealths != nil {
		localVarQueryParams.Add("projectHealths", parameterToString(*r.projectHealths, ""))
	}
	if r.projectHealth != nil {
		localVarQueryParams.Add("projectHealth", parameterToString(*r.projectHealth, ""))
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
	if r.reactions != nil {
		localVarQueryParams.Add("reactions", parameterToString(*r.reactions, ""))
	}
	if r.onlyStarredProjects != nil {
		localVarQueryParams.Add("onlyStarredProjects", parameterToString(*r.onlyStarredProjects, ""))
	}
	if r.matchAllProjectTags != nil {
		localVarQueryParams.Add("matchAllProjectTags", parameterToString(*r.matchAllProjectTags, ""))
	}
	if r.includeArchivedProjects != nil {
		localVarQueryParams.Add("includeArchivedProjects", parameterToString(*r.includeArchivedProjects, ""))
	}
	if r.emoji != nil {
		localVarQueryParams.Add("emoji", parameterToString(*r.emoji, ""))
	}
	if r.activeOnly != nil {
		localVarQueryParams.Add("activeOnly", parameterToString(*r.activeOnly, ""))
	}
	if r.projectTagIds != nil {
		localVarQueryParams.Add("projectTagIds", parameterToString(*r.projectTagIds, "csv"))
	}
	if r.projectOwnerIds != nil {
		localVarQueryParams.Add("projectOwnerIds", parameterToString(*r.projectOwnerIds, "csv"))
	}
	if r.projectCompanyIds != nil {
		localVarQueryParams.Add("projectCompanyIds", parameterToString(*r.projectCompanyIds, "csv"))
	}
	if r.projectCategoryIds != nil {
		localVarQueryParams.Add("projectCategoryIds", parameterToString(*r.projectCategoryIds, "csv"))
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
	if r.fieldsProjectUpdates != nil {
		localVarQueryParams.Add("fields[projectUpdates]", parameterToString(*r.fieldsProjectUpdates, "csv"))
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
