/*
Yahoo Finance

Yahoo Finance API specification

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// CrumbAPIService CrumbAPI service
type CrumbAPIService service

type ApiGetCrumbRequest struct {
	ctx context.Context
	ApiService *CrumbAPIService
	a1 *string
}

// Yahoo cookie A1 for authentication
func (r ApiGetCrumbRequest) A1(a1 string) ApiGetCrumbRequest {
	r.a1 = &a1
	return r
}

func (r ApiGetCrumbRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GetCrumbExecute(r)
}

/*
GetCrumb Retrieve crumb for authentication

Retrieve crumb for Yahoo Finance API authentication

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCrumbRequest
*/
func (a *CrumbAPIService) GetCrumb(ctx context.Context) ApiGetCrumbRequest {
	return ApiGetCrumbRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *CrumbAPIService) GetCrumbExecute(r ApiGetCrumbRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CrumbAPIService.GetCrumb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/test/getcrumb"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.a1 == nil {
		return localVarReturnValue, nil, reportError("a1 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	
	// to determine the Cookies header
	localVarHTTPCookies := []string{}
	if r.a1 != nil && *r.a1 != "" {
		localVarHTTPCookies = append(localVarHTTPCookies, *r.a1)
	}

	// set Cookie header
	localVarHTTPCookie := selectHeaderCookie(localVarHTTPCookies)
	if localVarHTTPCookie != "" {
		localVarHeaderParams["Cookie"] = localVarHTTPCookie
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
