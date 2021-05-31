/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ChartApiService ChartApi service
type ChartApiService service

type ApiGetChartRequest struct {
	ctx _context.Context
	ApiService *ChartApiService
	symbol string
	interval *Interval
	period1 *int64
	period2 *int64
	region *string
	includePrePost *bool
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
func (r ApiGetChartRequest) Region(region string) ApiGetChartRequest {
	r.region = &region
	return r
}
func (r ApiGetChartRequest) IncludePrePost(includePrePost bool) ApiGetChartRequest {
	r.includePrePost = &includePrePost
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

func (r ApiGetChartRequest) Execute() (ChartResponse, *_nethttp.Response, error) {
	return r.ApiService.GetChartExecute(r)
}

/*
 * GetChart Method for GetChart
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol
 * @return ApiGetChartRequest
 */
func (a *ChartApiService) GetChart(ctx _context.Context, symbol string) ApiGetChartRequest {
	return ApiGetChartRequest{
		ApiService: a,
		ctx: ctx,
		symbol: symbol,
	}
}

/*
 * Execute executes the request
 * @return ChartResponse
 */
func (a *ChartApiService) GetChartExecute(r ApiGetChartRequest) (ChartResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ChartResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChartApiService.GetChart")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v8/finance/chart/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", _neturl.PathEscape(parameterToString(r.symbol, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.period1 == nil {
		return localVarReturnValue, nil, reportError("period1 is required and must be specified")
	}
	if r.period2 == nil {
		return localVarReturnValue, nil, reportError("period2 is required and must be specified")
	}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.includePrePost != nil {
		localVarQueryParams.Add("includePrePost", parameterToString(*r.includePrePost, ""))
	}
	if r.lang != nil {
		localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	}
	localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	if r.useYfid != nil {
		localVarQueryParams.Add("useYfid", parameterToString(*r.useYfid, ""))
	}
	localVarQueryParams.Add("period1", parameterToString(*r.period1, ""))
	localVarQueryParams.Add("period2", parameterToString(*r.period2, ""))
	if r.corsDomain != nil {
		localVarQueryParams.Add("corsDomain", parameterToString(*r.corsDomain, ""))
	}
	if r.tsrc != nil {
		localVarQueryParams.Add(".tsrc", parameterToString(*r.tsrc, ""))
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
