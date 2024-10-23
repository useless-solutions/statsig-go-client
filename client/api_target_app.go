/*
Console API

       The \"Console API\" is the CRUD API for performing the actions offered on console.statsig.com without needing to go through the web UI.       If you have any feature requests, drop on in to our [slack channel](https://www.statsig.com/slack) and let us know.       <br /><br />       <b>Authorization</b>       <br />       All requests must include the **STATSIG-API-KEY** field in the header. The value should be a **Console API Key** which can be created in the Project Settings on [console.statsig.com/api_keys](https://console.statsig.com/api_keys)       <br /><br />       <b>Rate Limiting</b>       <br />       Requests to the Console API are limited to <code>~ 100reqs / 10secs and ~ 900reqs / 15 mins</code>.       <br /><br />       <b>Keyboard Search</b>       <br />       Use <code>Ctrl/Cmd + K</code> to search for specific endpoints.       

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TargetAppAPIService TargetAppAPI service
type TargetAppAPIService service

type ApiConsoleV1TargetAppControllerBulkUpdateRequest struct {
	ctx context.Context
	ApiService *TargetAppAPIService
	bulkAssignConfigTargetAppDto *BulkAssignConfigTargetAppDto
	xRespectReviewSettings *string
}

func (r ApiConsoleV1TargetAppControllerBulkUpdateRequest) BulkAssignConfigTargetAppDto(bulkAssignConfigTargetAppDto BulkAssignConfigTargetAppDto) ApiConsoleV1TargetAppControllerBulkUpdateRequest {
	r.bulkAssignConfigTargetAppDto = &bulkAssignConfigTargetAppDto
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1TargetAppControllerBulkUpdateRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1TargetAppControllerBulkUpdateRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1TargetAppControllerBulkUpdateRequest) Execute() (*ConsoleV1TargetAppControllerGenUpdate200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1TargetAppControllerBulkUpdateExecute(r)
}

/*
ConsoleV1TargetAppControllerBulkUpdate Bulk Assign Target Apps

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1TargetAppControllerBulkUpdateRequest
*/
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerBulkUpdate(ctx context.Context) ApiConsoleV1TargetAppControllerBulkUpdateRequest {
	return ApiConsoleV1TargetAppControllerBulkUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1TargetAppControllerGenUpdate200Response
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerBulkUpdateExecute(r ApiConsoleV1TargetAppControllerBulkUpdateRequest) (*ConsoleV1TargetAppControllerGenUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1TargetAppControllerGenUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetAppAPIService.ConsoleV1TargetAppControllerBulkUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/target_app"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bulkAssignConfigTargetAppDto == nil {
		return localVarReturnValue, nil, reportError("bulkAssignConfigTargetAppDto is required and must be specified")
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
	if r.xRespectReviewSettings != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-respect-review-settings", r.xRespectReviewSettings, "simple", "")
	}
	// body params
	localVarPostBody = r.bulkAssignConfigTargetAppDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["STATSIG-API-KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["STATSIG-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConsoleV1DynamicConfigControllerGenCreate401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConsoleV1DynamicConfigControllerGenList404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsoleV1TargetAppControllerGenCreateRequest struct {
	ctx context.Context
	ApiService *TargetAppAPIService
	targetAppCreateDto *TargetAppCreateDto
	xRespectReviewSettings *string
}

func (r ApiConsoleV1TargetAppControllerGenCreateRequest) TargetAppCreateDto(targetAppCreateDto TargetAppCreateDto) ApiConsoleV1TargetAppControllerGenCreateRequest {
	r.targetAppCreateDto = &targetAppCreateDto
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1TargetAppControllerGenCreateRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1TargetAppControllerGenCreateRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1TargetAppControllerGenCreateRequest) Execute() (*ConsoleV1TargetAppControllerGenRead200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1TargetAppControllerGenCreateExecute(r)
}

/*
ConsoleV1TargetAppControllerGenCreate Create Target App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1TargetAppControllerGenCreateRequest
*/
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenCreate(ctx context.Context) ApiConsoleV1TargetAppControllerGenCreateRequest {
	return ApiConsoleV1TargetAppControllerGenCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1TargetAppControllerGenRead200Response
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenCreateExecute(r ApiConsoleV1TargetAppControllerGenCreateRequest) (*ConsoleV1TargetAppControllerGenRead200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1TargetAppControllerGenRead200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetAppAPIService.ConsoleV1TargetAppControllerGenCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/target_app"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targetAppCreateDto == nil {
		return localVarReturnValue, nil, reportError("targetAppCreateDto is required and must be specified")
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
	if r.xRespectReviewSettings != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-respect-review-settings", r.xRespectReviewSettings, "simple", "")
	}
	// body params
	localVarPostBody = r.targetAppCreateDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["STATSIG-API-KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["STATSIG-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsoleV1TargetAppControllerGenDeleteRequest struct {
	ctx context.Context
	ApiService *TargetAppAPIService
	id string
	xRespectReviewSettings *string
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1TargetAppControllerGenDeleteRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1TargetAppControllerGenDeleteRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1TargetAppControllerGenDeleteRequest) Execute() (*ConsoleV1TargetAppControllerGenDelete200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1TargetAppControllerGenDeleteExecute(r)
}

/*
ConsoleV1TargetAppControllerGenDelete Delete Target App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of target app
 @return ApiConsoleV1TargetAppControllerGenDeleteRequest
*/
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenDelete(ctx context.Context, id string) ApiConsoleV1TargetAppControllerGenDeleteRequest {
	return ApiConsoleV1TargetAppControllerGenDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ConsoleV1TargetAppControllerGenDelete200Response
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenDeleteExecute(r ApiConsoleV1TargetAppControllerGenDeleteRequest) (*ConsoleV1TargetAppControllerGenDelete200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1TargetAppControllerGenDelete200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetAppAPIService.ConsoleV1TargetAppControllerGenDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/target_app/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.xRespectReviewSettings != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-respect-review-settings", r.xRespectReviewSettings, "simple", "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["STATSIG-API-KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["STATSIG-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsoleV1TargetAppControllerGenListRequest struct {
	ctx context.Context
	ApiService *TargetAppAPIService
	limit *int32
	page *int32
	xRespectReviewSettings *string
}

// Results per page
func (r ApiConsoleV1TargetAppControllerGenListRequest) Limit(limit int32) ApiConsoleV1TargetAppControllerGenListRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiConsoleV1TargetAppControllerGenListRequest) Page(page int32) ApiConsoleV1TargetAppControllerGenListRequest {
	r.page = &page
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1TargetAppControllerGenListRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1TargetAppControllerGenListRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1TargetAppControllerGenListRequest) Execute() (*ConsoleV1TargetAppControllerGenList200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1TargetAppControllerGenListExecute(r)
}

/*
ConsoleV1TargetAppControllerGenList List Target Apps

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1TargetAppControllerGenListRequest
*/
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenList(ctx context.Context) ApiConsoleV1TargetAppControllerGenListRequest {
	return ApiConsoleV1TargetAppControllerGenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1TargetAppControllerGenList200Response
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenListExecute(r ApiConsoleV1TargetAppControllerGenListRequest) (*ConsoleV1TargetAppControllerGenList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1TargetAppControllerGenList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetAppAPIService.ConsoleV1TargetAppControllerGenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/target_app"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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
	if r.xRespectReviewSettings != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-respect-review-settings", r.xRespectReviewSettings, "simple", "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["STATSIG-API-KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["STATSIG-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConsoleV1DynamicConfigControllerGenCreate401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsoleV1TargetAppControllerGenReadRequest struct {
	ctx context.Context
	ApiService *TargetAppAPIService
	id string
	xRespectReviewSettings *string
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1TargetAppControllerGenReadRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1TargetAppControllerGenReadRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1TargetAppControllerGenReadRequest) Execute() (*ConsoleV1TargetAppControllerGenRead200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1TargetAppControllerGenReadExecute(r)
}

/*
ConsoleV1TargetAppControllerGenRead Read Target App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id
 @return ApiConsoleV1TargetAppControllerGenReadRequest
*/
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenRead(ctx context.Context, id string) ApiConsoleV1TargetAppControllerGenReadRequest {
	return ApiConsoleV1TargetAppControllerGenReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ConsoleV1TargetAppControllerGenRead200Response
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenReadExecute(r ApiConsoleV1TargetAppControllerGenReadRequest) (*ConsoleV1TargetAppControllerGenRead200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1TargetAppControllerGenRead200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetAppAPIService.ConsoleV1TargetAppControllerGenRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/target_app/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.xRespectReviewSettings != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-respect-review-settings", r.xRespectReviewSettings, "simple", "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["STATSIG-API-KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["STATSIG-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConsoleV1DynamicConfigControllerGenCreate401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsoleV1TargetAppControllerGenUpdateRequest struct {
	ctx context.Context
	ApiService *TargetAppAPIService
	id string
	updateTargetAppDto *UpdateTargetAppDto
	xRespectReviewSettings *string
}

func (r ApiConsoleV1TargetAppControllerGenUpdateRequest) UpdateTargetAppDto(updateTargetAppDto UpdateTargetAppDto) ApiConsoleV1TargetAppControllerGenUpdateRequest {
	r.updateTargetAppDto = &updateTargetAppDto
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1TargetAppControllerGenUpdateRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1TargetAppControllerGenUpdateRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1TargetAppControllerGenUpdateRequest) Execute() (*ConsoleV1TargetAppControllerGenUpdate200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1TargetAppControllerGenUpdateExecute(r)
}

/*
ConsoleV1TargetAppControllerGenUpdate Update Target App

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id
 @return ApiConsoleV1TargetAppControllerGenUpdateRequest
*/
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenUpdate(ctx context.Context, id string) ApiConsoleV1TargetAppControllerGenUpdateRequest {
	return ApiConsoleV1TargetAppControllerGenUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ConsoleV1TargetAppControllerGenUpdate200Response
func (a *TargetAppAPIService) ConsoleV1TargetAppControllerGenUpdateExecute(r ApiConsoleV1TargetAppControllerGenUpdateRequest) (*ConsoleV1TargetAppControllerGenUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1TargetAppControllerGenUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetAppAPIService.ConsoleV1TargetAppControllerGenUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/target_app/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateTargetAppDto == nil {
		return localVarReturnValue, nil, reportError("updateTargetAppDto is required and must be specified")
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
	if r.xRespectReviewSettings != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-respect-review-settings", r.xRespectReviewSettings, "simple", "")
	}
	// body params
	localVarPostBody = r.updateTargetAppDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["STATSIG-API-KEY"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["STATSIG-API-KEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConsoleV1DynamicConfigControllerGenCreate401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConsoleV1DynamicConfigControllerGenList404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
