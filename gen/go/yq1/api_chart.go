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
)


// ChartAPIService ChartAPI service
type ChartAPIService service

type ApiGetChartRequest struct {
	ctx context.Context
	ApiService *ChartAPIService
	symbol string
	interval *Interval
	period1 *int64
	period2 *int64
	crumb *string
	a1 *string
	region *string
	includePrePost *bool
	events *[]string
	lang *string
	useYfid *bool
	corsDomain *string
	tsrc *string
}

func (r ApiGetChartRequest) Interval(interval Interval) ApiGetChartRequest {
	r.interval = &interval
	return r
}

func (r ApiGetChartRequest) Period1(period1 int64) ApiGetChartRequest {
	r.period1 = &period1
	return r
}

func (r ApiGetChartRequest) Period2(period2 int64) ApiGetChartRequest {
	r.period2 = &period2
	return r
}

// Yahoo cookie crumb
func (r ApiGetChartRequest) Crumb(crumb string) ApiGetChartRequest {
	r.crumb = &crumb
	return r
}

// Yahoo cookie A1 for authentication
func (r ApiGetChartRequest) A1(a1 string) ApiGetChartRequest {
	r.a1 = &a1
	return r
}

func (r ApiGetChartRequest) Region(region string) ApiGetChartRequest {
	r.region = &region
	return r
}

func (r ApiGetChartRequest) IncludePrePost(includePrePost bool) ApiGetChartRequest {
	r.includePrePost = &includePrePost
	return r
}

func (r ApiGetChartRequest) Events(events []string) ApiGetChartRequest {
	r.events = &events
	return r
}

func (r ApiGetChartRequest) Lang(lang string) ApiGetChartRequest {
	r.lang = &lang
	return r
}

func (r ApiGetChartRequest) UseYfid(useYfid bool) ApiGetChartRequest {
	r.useYfid = &useYfid
	return r
}

func (r ApiGetChartRequest) CorsDomain(corsDomain string) ApiGetChartRequest {
	r.corsDomain = &corsDomain
	return r
}

func (r ApiGetChartRequest) Tsrc(tsrc string) ApiGetChartRequest {
	r.tsrc = &tsrc
	return r
}

func (r ApiGetChartRequest) Execute() (*ChartResponse, *http.Response, error) {
	return r.ApiService.GetChartExecute(r)
}

/*
GetChart Method for GetChart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param symbol
 @return ApiGetChartRequest
*/
func (a *ChartAPIService) GetChart(ctx context.Context, symbol string) ApiGetChartRequest {
	return ApiGetChartRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

// Execute executes the request
//  @return ChartResponse
func (a *ChartAPIService) GetChartExecute(r ApiGetChartRequest) (*ChartResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChartResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartAPIService.GetChart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v8/finance/chart/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", url.PathEscape(parameterValueToString(r.symbol, "symbol")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.period1 == nil {
		return localVarReturnValue, nil, reportError("period1 is required and must be specified")
	}
	if r.period2 == nil {
		return localVarReturnValue, nil, reportError("period2 is required and must be specified")
	}
	if r.crumb == nil {
		return localVarReturnValue, nil, reportError("crumb is required and must be specified")
	}
	if r.a1 == nil {
		return localVarReturnValue, nil, reportError("a1 is required and must be specified")
	}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	if r.includePrePost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includePrePost", r.includePrePost, "")
	} else {
		var defaultValue bool = false
		r.includePrePost = &defaultValue
	}
	if r.events != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "events", r.events, "pipes")
	}
	if r.lang != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lang", r.lang, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
	if r.useYfid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "useYfid", r.useYfid, "")
	} else {
		var defaultValue bool = true
		r.useYfid = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "period1", r.period1, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "period2", r.period2, "")
	if r.corsDomain != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "corsDomain", r.corsDomain, "")
	}
	if r.tsrc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, ".tsrc", r.tsrc, "")
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
	
	// to determine the Cookies header
	localVarHTTPCookies := []string{}
	if r.a1 != nil && *r.a1 != "" {
		localVarHTTPCookies = append(localVarHTTPCookies, *r.a1)
	} else {
		localVarHTTPCookies = append(localVarHTTPCookies, "")
	}

	if r.crumb != nil && *r.crumb != "" {
		localVarHTTPCookies = append(localVarHTTPCookies, *r.crumb)
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
