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

// CalendarEventsApiService CalendarEventsApi service
type CalendarEventsApiService service

type ApiGETProjectsApiV3CalendarEventsCsvRequest struct {
	ctx _context.Context
	ApiService *CalendarEventsApiService
	startDate *string
	reportFormat *string
	endDate *string
	withTasks *bool
	withMilestones *bool
	withEvents *bool
	isReportDownload *bool
	includeTags *bool
	attendingOnly *bool
	typeIDs *[]int32
	targetUserIDs *[]int32
	targetProjectIDs *[]int32
}

func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) StartDate(startDate string) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) ReportFormat(reportFormat string) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.reportFormat = &reportFormat
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) EndDate(endDate string) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) WithTasks(withTasks bool) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.withTasks = &withTasks
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) WithMilestones(withMilestones bool) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.withMilestones = &withMilestones
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) WithEvents(withEvents bool) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.withEvents = &withEvents
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) IsReportDownload(isReportDownload bool) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.isReportDownload = &isReportDownload
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) IncludeTags(includeTags bool) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.includeTags = &includeTags
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) AttendingOnly(attendingOnly bool) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.attendingOnly = &attendingOnly
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) TypeIDs(typeIDs []int32) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.typeIDs = &typeIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) TargetUserIDs(targetUserIDs []int32) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.targetUserIDs = &targetUserIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) TargetProjectIDs(targetProjectIDs []int32) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	r.targetProjectIDs = &targetProjectIDs
	return r
}

func (r ApiGETProjectsApiV3CalendarEventsCsvRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3CalendarEventsCsvExecute(r)
}

/*
 * GETProjectsApiV3CalendarEventsCsv Generate agenda report in CSV format
 * Generates an agenda report in CSV format containing all the events for the
provided filters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3CalendarEventsCsvRequest
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsCsv(ctx _context.Context) ApiGETProjectsApiV3CalendarEventsCsvRequest {
	return ApiGETProjectsApiV3CalendarEventsCsvRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsCsvExecute(r ApiGETProjectsApiV3CalendarEventsCsvRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsApiService.GETProjectsApiV3CalendarEventsCsv")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/calendar/events.csv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.reportFormat != nil {
		localVarQueryParams.Add("reportFormat", parameterToString(*r.reportFormat, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.withTasks != nil {
		localVarQueryParams.Add("withTasks", parameterToString(*r.withTasks, ""))
	}
	if r.withMilestones != nil {
		localVarQueryParams.Add("withMilestones", parameterToString(*r.withMilestones, ""))
	}
	if r.withEvents != nil {
		localVarQueryParams.Add("withEvents", parameterToString(*r.withEvents, ""))
	}
	if r.isReportDownload != nil {
		localVarQueryParams.Add("isReportDownload", parameterToString(*r.isReportDownload, ""))
	}
	if r.includeTags != nil {
		localVarQueryParams.Add("includeTags", parameterToString(*r.includeTags, ""))
	}
	if r.attendingOnly != nil {
		localVarQueryParams.Add("attendingOnly", parameterToString(*r.attendingOnly, ""))
	}
	if r.typeIDs != nil {
		localVarQueryParams.Add("typeIDs", parameterToString(*r.typeIDs, "csv"))
	}
	if r.targetUserIDs != nil {
		localVarQueryParams.Add("targetUserIDs", parameterToString(*r.targetUserIDs, "csv"))
	}
	if r.targetProjectIDs != nil {
		localVarQueryParams.Add("targetProjectIDs", parameterToString(*r.targetProjectIDs, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGETProjectsApiV3CalendarEventsHtmlRequest struct {
	ctx _context.Context
	ApiService *CalendarEventsApiService
	startDate *string
	reportFormat *string
	endDate *string
	withTasks *bool
	withMilestones *bool
	withEvents *bool
	isReportDownload *bool
	includeTags *bool
	attendingOnly *bool
	typeIDs *[]int32
	targetUserIDs *[]int32
	targetProjectIDs *[]int32
}

func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) StartDate(startDate string) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) ReportFormat(reportFormat string) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.reportFormat = &reportFormat
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) EndDate(endDate string) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) WithTasks(withTasks bool) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.withTasks = &withTasks
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) WithMilestones(withMilestones bool) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.withMilestones = &withMilestones
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) WithEvents(withEvents bool) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.withEvents = &withEvents
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) IsReportDownload(isReportDownload bool) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.isReportDownload = &isReportDownload
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) IncludeTags(includeTags bool) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.includeTags = &includeTags
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) AttendingOnly(attendingOnly bool) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.attendingOnly = &attendingOnly
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) TypeIDs(typeIDs []int32) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.typeIDs = &typeIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) TargetUserIDs(targetUserIDs []int32) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.targetUserIDs = &targetUserIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) TargetProjectIDs(targetProjectIDs []int32) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	r.targetProjectIDs = &targetProjectIDs
	return r
}

func (r ApiGETProjectsApiV3CalendarEventsHtmlRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3CalendarEventsHtmlExecute(r)
}

/*
 * GETProjectsApiV3CalendarEventsHtml Generate agenda report in HTML format
 * Generates an agenda report in HTML format containing all the events for the
provided filters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3CalendarEventsHtmlRequest
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsHtml(ctx _context.Context) ApiGETProjectsApiV3CalendarEventsHtmlRequest {
	return ApiGETProjectsApiV3CalendarEventsHtmlRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsHtmlExecute(r ApiGETProjectsApiV3CalendarEventsHtmlRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsApiService.GETProjectsApiV3CalendarEventsHtml")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/calendar/events.html"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.reportFormat != nil {
		localVarQueryParams.Add("reportFormat", parameterToString(*r.reportFormat, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.withTasks != nil {
		localVarQueryParams.Add("withTasks", parameterToString(*r.withTasks, ""))
	}
	if r.withMilestones != nil {
		localVarQueryParams.Add("withMilestones", parameterToString(*r.withMilestones, ""))
	}
	if r.withEvents != nil {
		localVarQueryParams.Add("withEvents", parameterToString(*r.withEvents, ""))
	}
	if r.isReportDownload != nil {
		localVarQueryParams.Add("isReportDownload", parameterToString(*r.isReportDownload, ""))
	}
	if r.includeTags != nil {
		localVarQueryParams.Add("includeTags", parameterToString(*r.includeTags, ""))
	}
	if r.attendingOnly != nil {
		localVarQueryParams.Add("attendingOnly", parameterToString(*r.attendingOnly, ""))
	}
	if r.typeIDs != nil {
		localVarQueryParams.Add("typeIDs", parameterToString(*r.typeIDs, "csv"))
	}
	if r.targetUserIDs != nil {
		localVarQueryParams.Add("targetUserIDs", parameterToString(*r.targetUserIDs, "csv"))
	}
	if r.targetProjectIDs != nil {
		localVarQueryParams.Add("targetProjectIDs", parameterToString(*r.targetProjectIDs, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGETProjectsApiV3CalendarEventsJsonRequest struct {
	ctx _context.Context
	ApiService *CalendarEventsApiService
	startDate *string
	projectStatuses *string
	orderMode *string
	orderBy *string
	endDate *string
	projectHealths *int32
	pageSize *int32
	page *int32
	onlyStarredProjects *bool
	matchAllProjectTags *bool
	projectTagIds *[]int32
	projectOwnerIds *[]int32
	projectIds *[]int32
	projectCompanyIds *[]int32
	projectCategoryIds *[]int32
	fieldsCalendarEvents *[]string
}

func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) StartDate(startDate string) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectStatuses(projectStatuses string) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectStatuses = &projectStatuses
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) OrderMode(orderMode string) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.orderMode = &orderMode
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) OrderBy(orderBy string) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) EndDate(endDate string) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectHealths(projectHealths int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectHealths = &projectHealths
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) PageSize(pageSize int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) Page(page int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.page = &page
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) OnlyStarredProjects(onlyStarredProjects bool) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.onlyStarredProjects = &onlyStarredProjects
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) MatchAllProjectTags(matchAllProjectTags bool) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.matchAllProjectTags = &matchAllProjectTags
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectTagIds(projectTagIds []int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectTagIds = &projectTagIds
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectOwnerIds(projectOwnerIds []int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectOwnerIds = &projectOwnerIds
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectIds(projectIds []int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectIds = &projectIds
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectCompanyIds(projectCompanyIds []int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectCompanyIds = &projectCompanyIds
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) ProjectCategoryIds(projectCategoryIds []int32) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.projectCategoryIds = &projectCategoryIds
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) FieldsCalendarEvents(fieldsCalendarEvents []string) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	r.fieldsCalendarEvents = &fieldsCalendarEvents
	return r
}

func (r ApiGETProjectsApiV3CalendarEventsJsonRequest) Execute() (CalendarEventsResponse, *_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3CalendarEventsJsonExecute(r)
}

/*
 * GETProjectsApiV3CalendarEventsJson Retrieve the calendar events
 * This call will return all calendar events that the current user can see in
the provided date range.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3CalendarEventsJsonRequest
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsJson(ctx _context.Context) ApiGETProjectsApiV3CalendarEventsJsonRequest {
	return ApiGETProjectsApiV3CalendarEventsJsonRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return CalendarEventsResponse
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsJsonExecute(r ApiGETProjectsApiV3CalendarEventsJsonRequest) (CalendarEventsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CalendarEventsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsApiService.GETProjectsApiV3CalendarEventsJson")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/calendar/events.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
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
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
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
	if r.matchAllProjectTags != nil {
		localVarQueryParams.Add("matchAllProjectTags", parameterToString(*r.matchAllProjectTags, ""))
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
	if r.fieldsCalendarEvents != nil {
		localVarQueryParams.Add("fields[calendarEvents]", parameterToString(*r.fieldsCalendarEvents, "csv"))
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

type ApiGETProjectsApiV3CalendarEventsPdfRequest struct {
	ctx _context.Context
	ApiService *CalendarEventsApiService
	startDate *string
	reportFormat *string
	endDate *string
	withTasks *bool
	withMilestones *bool
	withEvents *bool
	isReportDownload *bool
	includeTags *bool
	attendingOnly *bool
	typeIDs *[]int32
	targetUserIDs *[]int32
	targetProjectIDs *[]int32
}

func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) StartDate(startDate string) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) ReportFormat(reportFormat string) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.reportFormat = &reportFormat
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) EndDate(endDate string) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) WithTasks(withTasks bool) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.withTasks = &withTasks
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) WithMilestones(withMilestones bool) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.withMilestones = &withMilestones
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) WithEvents(withEvents bool) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.withEvents = &withEvents
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) IsReportDownload(isReportDownload bool) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.isReportDownload = &isReportDownload
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) IncludeTags(includeTags bool) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.includeTags = &includeTags
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) AttendingOnly(attendingOnly bool) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.attendingOnly = &attendingOnly
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) TypeIDs(typeIDs []int32) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.typeIDs = &typeIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) TargetUserIDs(targetUserIDs []int32) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.targetUserIDs = &targetUserIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) TargetProjectIDs(targetProjectIDs []int32) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	r.targetProjectIDs = &targetProjectIDs
	return r
}

func (r ApiGETProjectsApiV3CalendarEventsPdfRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3CalendarEventsPdfExecute(r)
}

/*
 * GETProjectsApiV3CalendarEventsPdf Generate agenda report in PDF format
 * Generates an agenda report in PDF format containing all the events for the
provided filters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3CalendarEventsPdfRequest
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsPdf(ctx _context.Context) ApiGETProjectsApiV3CalendarEventsPdfRequest {
	return ApiGETProjectsApiV3CalendarEventsPdfRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsPdfExecute(r ApiGETProjectsApiV3CalendarEventsPdfRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsApiService.GETProjectsApiV3CalendarEventsPdf")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/calendar/events.pdf"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.reportFormat != nil {
		localVarQueryParams.Add("reportFormat", parameterToString(*r.reportFormat, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.withTasks != nil {
		localVarQueryParams.Add("withTasks", parameterToString(*r.withTasks, ""))
	}
	if r.withMilestones != nil {
		localVarQueryParams.Add("withMilestones", parameterToString(*r.withMilestones, ""))
	}
	if r.withEvents != nil {
		localVarQueryParams.Add("withEvents", parameterToString(*r.withEvents, ""))
	}
	if r.isReportDownload != nil {
		localVarQueryParams.Add("isReportDownload", parameterToString(*r.isReportDownload, ""))
	}
	if r.includeTags != nil {
		localVarQueryParams.Add("includeTags", parameterToString(*r.includeTags, ""))
	}
	if r.attendingOnly != nil {
		localVarQueryParams.Add("attendingOnly", parameterToString(*r.attendingOnly, ""))
	}
	if r.typeIDs != nil {
		localVarQueryParams.Add("typeIDs", parameterToString(*r.typeIDs, "csv"))
	}
	if r.targetUserIDs != nil {
		localVarQueryParams.Add("targetUserIDs", parameterToString(*r.targetUserIDs, "csv"))
	}
	if r.targetProjectIDs != nil {
		localVarQueryParams.Add("targetProjectIDs", parameterToString(*r.targetProjectIDs, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/pdf"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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

type ApiGETProjectsApiV3CalendarEventsXlsxRequest struct {
	ctx _context.Context
	ApiService *CalendarEventsApiService
	startDate *string
	reportFormat *string
	endDate *string
	withTasks *bool
	withMilestones *bool
	withEvents *bool
	isReportDownload *bool
	includeTags *bool
	attendingOnly *bool
	typeIDs *[]int32
	targetUserIDs *[]int32
	targetProjectIDs *[]int32
}

func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) StartDate(startDate string) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.startDate = &startDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) ReportFormat(reportFormat string) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.reportFormat = &reportFormat
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) EndDate(endDate string) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.endDate = &endDate
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) WithTasks(withTasks bool) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.withTasks = &withTasks
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) WithMilestones(withMilestones bool) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.withMilestones = &withMilestones
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) WithEvents(withEvents bool) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.withEvents = &withEvents
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) IsReportDownload(isReportDownload bool) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.isReportDownload = &isReportDownload
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) IncludeTags(includeTags bool) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.includeTags = &includeTags
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) AttendingOnly(attendingOnly bool) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.attendingOnly = &attendingOnly
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) TypeIDs(typeIDs []int32) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.typeIDs = &typeIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) TargetUserIDs(targetUserIDs []int32) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.targetUserIDs = &targetUserIDs
	return r
}
func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) TargetProjectIDs(targetProjectIDs []int32) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	r.targetProjectIDs = &targetProjectIDs
	return r
}

func (r ApiGETProjectsApiV3CalendarEventsXlsxRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GETProjectsApiV3CalendarEventsXlsxExecute(r)
}

/*
 * GETProjectsApiV3CalendarEventsXlsx Generate agenda report in XLSX format
 * Generates an agenda report in XLSX format containing all the events for the
provided filters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGETProjectsApiV3CalendarEventsXlsxRequest
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsXlsx(ctx _context.Context) ApiGETProjectsApiV3CalendarEventsXlsxRequest {
	return ApiGETProjectsApiV3CalendarEventsXlsxRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CalendarEventsApiService) GETProjectsApiV3CalendarEventsXlsxExecute(r ApiGETProjectsApiV3CalendarEventsXlsxRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarEventsApiService.GETProjectsApiV3CalendarEventsXlsx")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/projects/api/v3/calendar/events.xlsx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.reportFormat != nil {
		localVarQueryParams.Add("reportFormat", parameterToString(*r.reportFormat, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.withTasks != nil {
		localVarQueryParams.Add("withTasks", parameterToString(*r.withTasks, ""))
	}
	if r.withMilestones != nil {
		localVarQueryParams.Add("withMilestones", parameterToString(*r.withMilestones, ""))
	}
	if r.withEvents != nil {
		localVarQueryParams.Add("withEvents", parameterToString(*r.withEvents, ""))
	}
	if r.isReportDownload != nil {
		localVarQueryParams.Add("isReportDownload", parameterToString(*r.isReportDownload, ""))
	}
	if r.includeTags != nil {
		localVarQueryParams.Add("includeTags", parameterToString(*r.includeTags, ""))
	}
	if r.attendingOnly != nil {
		localVarQueryParams.Add("attendingOnly", parameterToString(*r.attendingOnly, ""))
	}
	if r.typeIDs != nil {
		localVarQueryParams.Add("typeIDs", parameterToString(*r.typeIDs, "csv"))
	}
	if r.targetUserIDs != nil {
		localVarQueryParams.Add("targetUserIDs", parameterToString(*r.targetUserIDs, "csv"))
	}
	if r.targetProjectIDs != nil {
		localVarQueryParams.Add("targetProjectIDs", parameterToString(*r.targetProjectIDs, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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