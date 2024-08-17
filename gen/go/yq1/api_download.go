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
	"strings"
	"os"
)


// DownloadAPIService DownloadAPI service
type DownloadAPIService service

type ApiDownloadRequest struct {
	ctx context.Context
	ApiService *DownloadAPIService
	symbol string
	period1 *int64
	period2 *int64
	interval *Interval
	crumb *string
	a1 *string
	events *[]string
}

// The start period for the data (Unix timestamp)
func (r ApiDownloadRequest) Period1(period1 int64) ApiDownloadRequest {
	r.period1 = &period1
	return r
}

// The end period for the data (Unix timestamp)
func (r ApiDownloadRequest) Period2(period2 int64) ApiDownloadRequest {
	r.period2 = &period2
	return r
}

// The data interval (e.g., 1d, 1wk, 1mo)
func (r ApiDownloadRequest) Interval(interval Interval) ApiDownloadRequest {
	r.interval = &interval
	return r
}

// Yahoo cookie crumb
func (r ApiDownloadRequest) Crumb(crumb string) ApiDownloadRequest {
	r.crumb = &crumb
	return r
}

// Yahoo cookie A1 for authentication
func (r ApiDownloadRequest) A1(a1 string) ApiDownloadRequest {
	r.a1 = &a1
	return r
}

// The events to include (e.g., div, split)
func (r ApiDownloadRequest) Events(events []string) ApiDownloadRequest {
	r.events = &events
	return r
}

func (r ApiDownloadRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.DownloadExecute(r)
}

/*
Download Download financial data for a specific symbol

Retrieve financial data in CSV format for the given symbol

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param symbol The stock symbol to download data for
 @return ApiDownloadRequest
*/
func (a *DownloadAPIService) Download(ctx context.Context, symbol string) ApiDownloadRequest {
	return ApiDownloadRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

// Execute executes the request
//  @return *os.File
func (a *DownloadAPIService) DownloadExecute(r ApiDownloadRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DownloadAPIService.Download")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v7/finance/download/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", url.PathEscape(parameterValueToString(r.symbol, "symbol")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.period1 == nil {
		return localVarReturnValue, nil, reportError("period1 is required and must be specified")
	}
	if r.period2 == nil {
		return localVarReturnValue, nil, reportError("period2 is required and must be specified")
	}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.crumb == nil {
		return localVarReturnValue, nil, reportError("crumb is required and must be specified")
	}
	if r.a1 == nil {
		return localVarReturnValue, nil, reportError("a1 is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "period1", r.period1, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "period2", r.period2, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
	if r.events != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "events", r.events, "pipes")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/csv"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	
	// to determine the Cookies header
	localVarHTTPCookies := []http.Cookie{}
	if *r.crumb != "" {
		localVarHTTPCookies = append(localVarHTTPCookies, http.Cookie{Name:"Crumb",Value:*r.crumb})
	}
	if *r.a1 != "" {
		localVarHTTPCookies = append(localVarHTTPCookies, http.Cookie{Name:"A1",Value:*r.a1})
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
