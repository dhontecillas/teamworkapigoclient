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
	"time"
)

// Linger please
var (
	_ _context.Context
)

// ProjectApiService ProjectApi service
type ProjectApiService service

type ApiGETProjectsApiV3ProjectsIdJsonRequest struct {
	ctx _context.Context
	ApiService *ProjectApiService
	updatedAfter *time.Time
	searchTerm *string
	reportType *string
	reportFormat *string
	projectType *string
	projectStatuses *string
	orderMode *string
	orderBy *string
	minLastActivityDate *string
	maxLastActivityDate *string
	projectHealths *int32
	pageSize *int32
	page *int32
	onlyStarredProjects *bool
	onlyProjectsWithExplicitMembership *bool
	onlyArchivedProjects *bool
	matchAllProjectTags *bool
	isReportDownload *bool
	includeProjectUserInfo *bool
	includeCustomFields *bool
	includeArchivedProjects *bool
	hideObservedProjects *bool
	projectTagIds *[]int32
	projectOwnerIds *[]int32
	projectIds *[]int32
	projectCompanyIds *[]int32
	projectCategoryIds *[]int32
	include *[]string
	fieldsUsers *[]string
	fieldsTags *[]string
	fieldsProjects *[]string
	fieldsProjectcategories *[]string
	fieldsProjectUpdates *[]string
	fieldsProjectBudgets *[]string
	fieldsPortfolioColumns *[]string
	fieldsPortfolioCards *[]string
	fieldsPortfolioBoards *[]string
	fieldsCustomfields *[]string
	fieldsCustomfieldProjects *[]string
	fieldsCompanies *[]string
}

func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) UpdatedAfter(updatedAfter time.Time) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.updatedAfter = &updatedAfter
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) SearchTerm(searchTerm string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.searchTerm = &searchTerm
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ReportType(reportType string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.reportType = &reportType
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ReportFormat(reportFormat string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.reportFormat = &reportFormat
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectType(projectType string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectType = &projectType
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectStatuses(projectStatuses string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectStatuses = &projectStatuses
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) OrderMode(orderMode string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.orderMode = &orderMode
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) OrderBy(orderBy string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) MinLastActivityDate(minLastActivityDate string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.minLastActivityDate = &minLastActivityDate
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) MaxLastActivityDate(maxLastActivityDate string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.maxLastActivityDate = &maxLastActivityDate
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectHealths(projectHealths int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectHealths = &projectHealths
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) PageSize(pageSize int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) Page(page int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) OnlyStarredProjects(onlyStarredProjects bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.onlyStarredProjects = &onlyStarredProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) OnlyProjectsWithExplicitMembership(onlyProjectsWithExplicitMembership bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.onlyProjectsWithExplicitMembership = &onlyProjectsWithExplicitMembership
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) OnlyArchivedProjects(onlyArchivedProjects bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.onlyArchivedProjects = &onlyArchivedProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) MatchAllProjectTags(matchAllProjectTags bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.matchAllProjectTags = &matchAllProjectTags
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) IsReportDownload(isReportDownload bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.isReportDownload = &isReportDownload
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) IncludeProjectUserInfo(includeProjectUserInfo bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.includeProjectUserInfo = &includeProjectUserInfo
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) IncludeCustomFields(includeCustomFields bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.includeCustomFields = &includeCustomFields
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) IncludeArchivedProjects(includeArchivedProjects bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.includeArchivedProjects = &includeArchivedProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) HideObservedProjects(hideObservedProjects bool) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.hideObservedProjects = &hideObservedProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectTagIds(projectTagIds []int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectTagIds = &projectTagIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectOwnerIds(projectOwnerIds []int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectOwnerIds = &projectOwnerIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectIds(projectIds []int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectIds = &projectIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectCompanyIds(projectCompanyIds []int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectCompanyIds = &projectCompanyIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) ProjectCategoryIds(projectCategoryIds []int32) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.projectCategoryIds = &projectCategoryIds
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) Include(include []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.include = &include
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsUsers(fieldsUsers []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsUsers = &fieldsUsers
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsTags(fieldsTags []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsTags = &fieldsTags
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsProjects(fieldsProjects []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsProjects = &fieldsProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsProjectcategories(fieldsProjectcategories []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsProjectcategories = &fieldsProjectcategories
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsProjectUpdates(fieldsProjectUpdates []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsProjectUpdates = &fieldsProjectUpdates
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsProjectBudgets(fieldsProjectBudgets []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsProjectBudgets = &fieldsProjectBudgets
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsPortfolioColumns(fieldsPortfolioColumns []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsPortfolioColumns = &fieldsPortfolioColumns
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsPortfolioCards(fieldsPortfolioCards []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsPortfolioCards = &fieldsPortfolioCards
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsPortfolioBoards(fieldsPortfolioBoards []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsPortfolioBoards = &fieldsPortfolioBoards
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsCustomfields(fieldsCustomfields []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsCustomfields = &fieldsCustomfields
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsCustomfieldProjects(fieldsCustomfieldProjects []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsCustomfieldProjects = &fieldsCustomfieldProjects
	return r
}
func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) FieldsCompanies(fieldsCompanies []string) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	r.fieldsCompanies = &fieldsCompanies
	return r
}

func (r ApiGETProjectsApiV3ProjectsIdJsonRequest) Execute() (ProjectSingleResponse, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3ProjectsIdJsonExecute(r)
}

/*
 * GETProjectsApiV3ProjectsIdJson Returns a project
 * On this endpoint you can filter by project custom fields. The syntax for the
query parameter is the following:

    projectCustomField[id][op]=value

Where:
  - [id] is the custom field ID
  - [op] is the operator to apply when filtering, different operators are
    allowed according to the custom field type
  - [value] is the value to apply when filtering

For example, if I want to filter a dropdown custom field with ID 10 to only
return entries that have the value "Option1" we would do the following:

    projectCustomField[10][eq]=Option1

The allowed operators are:
  - like
  - not-like
  - eq
  - not
  - lt
  - gt
  - any
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3ProjectsIdJsonRequest
 */
func (a *ProjectApiService) GETProjectsApiV3ProjectsIdJson(ctx _context.Context) ApiGETProjectsApiV3ProjectsIdJsonRequest {
	return ApiGETProjectsApiV3ProjectsIdJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ProjectSingleResponse
 */
func (a *ProjectApiService) GETProjectsApiV3ProjectsIdJsonExecute(r ApiGETProjectsApiV3ProjectsIdJsonRequest) (ProjectSingleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProjectSingleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectApiService.GETProjectsApiV3ProjectsIdJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/projects/:id.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.updatedAfter != nil {
		localVarQueryParams.Add("updatedAfter", parameterToString(*r.updatedAfter, ""))
	}
	if r.searchTerm != nil {
		localVarQueryParams.Add("searchTerm", parameterToString(*r.searchTerm, ""))
	}
	if r.reportType != nil {
		localVarQueryParams.Add("reportType", parameterToString(*r.reportType, ""))
	}
	if r.reportFormat != nil {
		localVarQueryParams.Add("reportFormat", parameterToString(*r.reportFormat, ""))
	}
	if r.projectType != nil {
		localVarQueryParams.Add("projectType", parameterToString(*r.projectType, ""))
	}
	if r.projectStatuses != nil {
		localVarQueryParams.Add("projectStatuses", parameterToString(*r.projectStatuses, ""))
	}
	if r.orderMode != nil {
		localVarQueryParams.Add("orderMode", parameterToString(*r.orderMode, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.minLastActivityDate != nil {
		localVarQueryParams.Add("minLastActivityDate", parameterToString(*r.minLastActivityDate, ""))
	}
	if r.maxLastActivityDate != nil {
		localVarQueryParams.Add("maxLastActivityDate", parameterToString(*r.maxLastActivityDate, ""))
	}
	if r.projectHealths != nil {
		localVarQueryParams.Add("projectHealths", parameterToString(*r.projectHealths, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.onlyStarredProjects != nil {
		localVarQueryParams.Add("onlyStarredProjects", parameterToString(*r.onlyStarredProjects, ""))
	}
	if r.onlyProjectsWithExplicitMembership != nil {
		localVarQueryParams.Add("onlyProjectsWithExplicitMembership", parameterToString(*r.onlyProjectsWithExplicitMembership, ""))
	}
	if r.onlyArchivedProjects != nil {
		localVarQueryParams.Add("onlyArchivedProjects", parameterToString(*r.onlyArchivedProjects, ""))
	}
	if r.matchAllProjectTags != nil {
		localVarQueryParams.Add("matchAllProjectTags", parameterToString(*r.matchAllProjectTags, ""))
	}
	if r.isReportDownload != nil {
		localVarQueryParams.Add("isReportDownload", parameterToString(*r.isReportDownload, ""))
	}
	if r.includeProjectUserInfo != nil {
		localVarQueryParams.Add("includeProjectUserInfo", parameterToString(*r.includeProjectUserInfo, ""))
	}
	if r.includeCustomFields != nil {
		localVarQueryParams.Add("includeCustomFields", parameterToString(*r.includeCustomFields, ""))
	}
	if r.includeArchivedProjects != nil {
		localVarQueryParams.Add("includeArchivedProjects", parameterToString(*r.includeArchivedProjects, ""))
	}
	if r.hideObservedProjects != nil {
		localVarQueryParams.Add("hideObservedProjects", parameterToString(*r.hideObservedProjects, ""))
	}
	if r.projectTagIds != nil {
		localVarQueryParams.Add("projectTagIds", parameterToString(*r.projectTagIds, "csv"))
	}
	if r.projectOwnerIds != nil {
		localVarQueryParams.Add("projectOwnerIds", parameterToString(*r.projectOwnerIds, "csv"))
	}
	if r.projectIds != nil {
		localVarQueryParams.Add("projectIds", parameterToString(*r.projectIds, "csv"))
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
	if r.fieldsTags != nil {
		localVarQueryParams.Add("fields[tags]", parameterToString(*r.fieldsTags, "csv"))
	}
	if r.fieldsProjects != nil {
		localVarQueryParams.Add("fields[projects]", parameterToString(*r.fieldsProjects, "csv"))
	}
	if r.fieldsProjectcategories != nil {
		localVarQueryParams.Add("fields[projectcategories]", parameterToString(*r.fieldsProjectcategories, "csv"))
	}
	if r.fieldsProjectUpdates != nil {
		localVarQueryParams.Add("fields[projectUpdates]", parameterToString(*r.fieldsProjectUpdates, "csv"))
	}
	if r.fieldsProjectBudgets != nil {
		localVarQueryParams.Add("fields[projectBudgets]", parameterToString(*r.fieldsProjectBudgets, "csv"))
	}
	if r.fieldsPortfolioColumns != nil {
		localVarQueryParams.Add("fields[portfolioColumns]", parameterToString(*r.fieldsPortfolioColumns, "csv"))
	}
	if r.fieldsPortfolioCards != nil {
		localVarQueryParams.Add("fields[portfolioCards]", parameterToString(*r.fieldsPortfolioCards, "csv"))
	}
	if r.fieldsPortfolioBoards != nil {
		localVarQueryParams.Add("fields[portfolioBoards]", parameterToString(*r.fieldsPortfolioBoards, "csv"))
	}
	if r.fieldsCustomfields != nil {
		localVarQueryParams.Add("fields[customfields]", parameterToString(*r.fieldsCustomfields, "csv"))
	}
	if r.fieldsCustomfieldProjects != nil {
		localVarQueryParams.Add("fields[customfieldProjects]", parameterToString(*r.fieldsCustomfieldProjects, "csv"))
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
