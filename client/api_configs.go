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
)


// ConfigsAPIService ConfigsAPI service
type ConfigsAPIService service

type ApiConsoleV1ChangeValidationControllerChangeValidationRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	changeValidationDto *ChangeValidationDto
	xRespectReviewSettings *string
}

func (r ApiConsoleV1ChangeValidationControllerChangeValidationRequest) ChangeValidationDto(changeValidationDto ChangeValidationDto) ApiConsoleV1ChangeValidationControllerChangeValidationRequest {
	r.changeValidationDto = &changeValidationDto
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1ChangeValidationControllerChangeValidationRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1ChangeValidationControllerChangeValidationRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1ChangeValidationControllerChangeValidationRequest) Execute() (*ConsoleV1ChangeValidationControllerChangeValidation200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1ChangeValidationControllerChangeValidationExecute(r)
}

/*
ConsoleV1ChangeValidationControllerChangeValidation Change Validation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1ChangeValidationControllerChangeValidationRequest
*/
func (a *ConfigsAPIService) ConsoleV1ChangeValidationControllerChangeValidation(ctx context.Context) ApiConsoleV1ChangeValidationControllerChangeValidationRequest {
	return ApiConsoleV1ChangeValidationControllerChangeValidationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1ChangeValidationControllerChangeValidation200Response
func (a *ConfigsAPIService) ConsoleV1ChangeValidationControllerChangeValidationExecute(r ApiConsoleV1ChangeValidationControllerChangeValidationRequest) (*ConsoleV1ChangeValidationControllerChangeValidation200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1ChangeValidationControllerChangeValidation200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.ConsoleV1ChangeValidationControllerChangeValidation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/change_validation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.changeValidationDto == nil {
		return localVarReturnValue, nil, reportError("changeValidationDto is required and must be specified")
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
	localVarPostBody = r.changeValidationDto
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

type ApiConsoleV1ExposureCountsControllerGenExposureCountRequest struct {
	ctx context.Context
	ApiService *ConfigsAPIService
	experiments *interface{}
	gates *interface{}
	xRespectReviewSettings *string
}

func (r ApiConsoleV1ExposureCountsControllerGenExposureCountRequest) Experiments(experiments interface{}) ApiConsoleV1ExposureCountsControllerGenExposureCountRequest {
	r.experiments = &experiments
	return r
}

func (r ApiConsoleV1ExposureCountsControllerGenExposureCountRequest) Gates(gates interface{}) ApiConsoleV1ExposureCountsControllerGenExposureCountRequest {
	r.gates = &gates
	return r
}

// Optional header to respect review settings for mutation endpoints.
func (r ApiConsoleV1ExposureCountsControllerGenExposureCountRequest) XRespectReviewSettings(xRespectReviewSettings string) ApiConsoleV1ExposureCountsControllerGenExposureCountRequest {
	r.xRespectReviewSettings = &xRespectReviewSettings
	return r
}

func (r ApiConsoleV1ExposureCountsControllerGenExposureCountRequest) Execute() (*ConsoleV1ExposureCountsControllerGenExposureCount200Response, *http.Response, error) {
	return r.ApiService.ConsoleV1ExposureCountsControllerGenExposureCountExecute(r)
}

/*
ConsoleV1ExposureCountsControllerGenExposureCount Read Exposure Event Count

Get the count of exposure events recently received by Statsig.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConsoleV1ExposureCountsControllerGenExposureCountRequest
*/
func (a *ConfigsAPIService) ConsoleV1ExposureCountsControllerGenExposureCount(ctx context.Context) ApiConsoleV1ExposureCountsControllerGenExposureCountRequest {
	return ApiConsoleV1ExposureCountsControllerGenExposureCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleV1ExposureCountsControllerGenExposureCount200Response
func (a *ConfigsAPIService) ConsoleV1ExposureCountsControllerGenExposureCountExecute(r ApiConsoleV1ExposureCountsControllerGenExposureCountRequest) (*ConsoleV1ExposureCountsControllerGenExposureCount200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleV1ExposureCountsControllerGenExposureCount200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.ConsoleV1ExposureCountsControllerGenExposureCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/console/v1/exposure_count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.experiments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "experiments", r.experiments, "form", "")
	}
	if r.gates != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gates", r.gates, "form", "")
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
