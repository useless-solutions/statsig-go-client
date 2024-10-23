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
	"reflect"
)


// AlertsAPIService AlertsAPI service
type AlertsAPIService service

type ApiConsoleV1AlertsControllerGenListRequest struct {
	ctx context.Context
	ApiService *AlertsAPIService
	creatorName *string
	creatorID *string
	tags *[]string
	limit *int32
	page *int32
	xRespectReviewSettings *string
}

// Name of the creator.
func (r ApiConsoleV1AlertsControllerGenListRequest) CreatorName(creatorName string) ApiConsoleV1AlertsControllerGenListRequest {
	r.creatorName = &creatorName
	return r
}

// ID of the user who created the entity.
func (r ApiConsoleV1AlertsControllerGenListRequest) CreatorID(creatorID string) ApiConsoleV1AlertsControllerGenListRequest {
	r.creatorID = &creatorID
	return r
}

// Filter by tags
func (r ApiConsoleV1AlertsControllerGenListRequest) Tags(tags []string) ApiConsoleV1AlertsControllerGenListRequest {
	r.tags = &tags
	return r
}

// Results per page
func (r ApiConsoleV1AlertsControllerGenListRequest) Limit(limit int32) ApiConsoleV1AlertsControllerGenListRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiConsoleV1AlertsControllerGenListRequest) Page(page int32) ApiConsoleV1AlertsControllerGenListRequest {
	r.page = &page
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1AlertsControllerGenListRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1AlertsControllerGenListRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1AlertsControllerGenListRequest) Execute() (*ConsoleV1AlertsControllerGenList200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1AlertsControllerGenListExecute(r)
}

/*
ConsoleV1AlertsControllerGenList List Alerts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1AlertsControllerGenListRequest
*/
func (a *AlertsAPIService) ConsoleV1AlertsControllerGenList(ctx context.Context) ApiConsoleV1AlertsControllerGenListRequest {
	return ApiConsoleV1AlertsControllerGenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1AlertsControllerGenList200Response
func (a *AlertsAPIService) ConsoleV1AlertsControllerGenListExecute(r ApiConsoleV1AlertsControllerGenListRequest) (*ConsoleV1AlertsControllerGenList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1AlertsControllerGenList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlertsAPIService.ConsoleV1AlertsControllerGenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/alerts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.creatorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creatorName", r.creatorName, "form", "")
	}
	if r.creatorID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creatorID", r.creatorID, "form", "")
	}
	if r.tags != nil {
		t := *r.tags
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "tags", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "tags", t, "form", "multi")
		}
	}
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ConsoleV1AlertsControllerGenList403Response
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