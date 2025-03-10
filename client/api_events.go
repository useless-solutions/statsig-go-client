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


// EventsAPIService EventsAPI service
type EventsAPIService service

type ApiConsoleV1EventsControllerGenListRequest struct {
	ctx context.Context
	ApiService *EventsAPIService
	limit *int32
	page *int32
	xRespectReviewSettings *string
}

// Results per page
func (r ApiConsoleV1EventsControllerGenListRequest) Limit(limit int32) ApiConsoleV1EventsControllerGenListRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiConsoleV1EventsControllerGenListRequest) Page(page int32) ApiConsoleV1EventsControllerGenListRequest {
	r.page = &page
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1EventsControllerGenListRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1EventsControllerGenListRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1EventsControllerGenListRequest) Execute() (*ConsoleV1EventsControllerGenList200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1EventsControllerGenListExecute(r)
}

/*
ConsoleV1EventsControllerGenList List Events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1EventsControllerGenListRequest
*/
func (a *EventsAPIService) ConsoleV1EventsControllerGenList(ctx context.Context) ApiConsoleV1EventsControllerGenListRequest {
	return ApiConsoleV1EventsControllerGenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1EventsControllerGenList200Response
func (a *EventsAPIService) ConsoleV1EventsControllerGenListExecute(r ApiConsoleV1EventsControllerGenListRequest) (*ConsoleV1EventsControllerGenList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1EventsControllerGenList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.ConsoleV1EventsControllerGenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/events"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ConsoleV1DynamicConfigControllerGenCreate400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiConsoleV1EventsControllerGenListMetricsByEventRequest struct {
	ctx context.Context
	ApiService *EventsAPIService
	eventName string
	limit *int32
	page *int32
	xRespectReviewSettings *string
}

// Results per page
func (r ApiConsoleV1EventsControllerGenListMetricsByEventRequest) Limit(limit int32) ApiConsoleV1EventsControllerGenListMetricsByEventRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiConsoleV1EventsControllerGenListMetricsByEventRequest) Page(page int32) ApiConsoleV1EventsControllerGenListMetricsByEventRequest {
	r.page = &page
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1EventsControllerGenListMetricsByEventRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1EventsControllerGenListMetricsByEventRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1EventsControllerGenListMetricsByEventRequest) Execute() (*ConsoleV1EventsControllerGenListMetricsByEvent200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1EventsControllerGenListMetricsByEventExecute(r)
}

/*
ConsoleV1EventsControllerGenListMetricsByEvent Get metrics using event name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventName
 @return ApiConsoleV1EventsControllerGenListMetricsByEventRequest
*/
func (a *EventsAPIService) ConsoleV1EventsControllerGenListMetricsByEvent(ctx context.Context, eventName string) ApiConsoleV1EventsControllerGenListMetricsByEventRequest {
	return ApiConsoleV1EventsControllerGenListMetricsByEventRequest{
		ApiService: a,
		ctx: ctx,
		eventName: eventName,
	}
}

// Execute executes the request
//  @return ConsoleV1EventsControllerGenListMetricsByEvent200Response
func (a *EventsAPIService) ConsoleV1EventsControllerGenListMetricsByEventExecute(r ApiConsoleV1EventsControllerGenListMetricsByEventRequest) (*ConsoleV1EventsControllerGenListMetricsByEvent200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1EventsControllerGenListMetricsByEvent200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.ConsoleV1EventsControllerGenListMetricsByEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/events/{eventName}/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"eventName"+"}", url.PathEscape(parameterValueToString(r.eventName, "eventName")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ConsoleV1DynamicConfigControllerGenCreate400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiConsoleV1EventsControllerGenListSpecificEventRequest struct {
	ctx context.Context
	ApiService *EventsAPIService
	eventName string
	limit *int32
	page *int32
	xRespectReviewSettings *string
}

// Results per page
func (r ApiConsoleV1EventsControllerGenListSpecificEventRequest) Limit(limit int32) ApiConsoleV1EventsControllerGenListSpecificEventRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiConsoleV1EventsControllerGenListSpecificEventRequest) Page(page int32) ApiConsoleV1EventsControllerGenListSpecificEventRequest {
	r.page = &page
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1EventsControllerGenListSpecificEventRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1EventsControllerGenListSpecificEventRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1EventsControllerGenListSpecificEventRequest) Execute() (*ConsoleV1EventsControllerGenListSpecificEvent200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1EventsControllerGenListSpecificEventExecute(r)
}

/*
ConsoleV1EventsControllerGenListSpecificEvent Get specific events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventName
 @return ApiConsoleV1EventsControllerGenListSpecificEventRequest
*/
func (a *EventsAPIService) ConsoleV1EventsControllerGenListSpecificEvent(ctx context.Context, eventName string) ApiConsoleV1EventsControllerGenListSpecificEventRequest {
	return ApiConsoleV1EventsControllerGenListSpecificEventRequest{
		ApiService: a,
		ctx: ctx,
		eventName: eventName,
	}
}

// Execute executes the request
//  @return ConsoleV1EventsControllerGenListSpecificEvent200Response
func (a *EventsAPIService) ConsoleV1EventsControllerGenListSpecificEventExecute(r ApiConsoleV1EventsControllerGenListSpecificEventRequest) (*ConsoleV1EventsControllerGenListSpecificEvent200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1EventsControllerGenListSpecificEvent200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.ConsoleV1EventsControllerGenListSpecificEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/events/{eventName}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventName"+"}", url.PathEscape(parameterValueToString(r.eventName, "eventName")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ConsoleV1DynamicConfigControllerGenCreate400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
