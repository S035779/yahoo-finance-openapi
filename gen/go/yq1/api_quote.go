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


// QuoteAPIService QuoteAPI service
type QuoteAPIService service

type ApiGetQuoteRequest struct {
	ctx context.Context
	ApiService *QuoteAPIService
	symbols *string
	crumb *string
	a1 *string
	formatted *bool
	region *string
	lang *string
	includePrePost *bool
	fields *string
	corsDomain *string
}

func (r ApiGetQuoteRequest) Symbols(symbols string) ApiGetQuoteRequest {
	r.symbols = &symbols
	return r
}

// Yahoo cookie crumb
func (r ApiGetQuoteRequest) Crumb(crumb string) ApiGetQuoteRequest {
	r.crumb = &crumb
	return r
}

// Yahoo cookie A1 for authentication
func (r ApiGetQuoteRequest) A1(a1 string) ApiGetQuoteRequest {
	r.a1 = &a1
	return r
}

func (r ApiGetQuoteRequest) Formatted(formatted bool) ApiGetQuoteRequest {
	r.formatted = &formatted
	return r
}

func (r ApiGetQuoteRequest) Region(region string) ApiGetQuoteRequest {
	r.region = &region
	return r
}

func (r ApiGetQuoteRequest) Lang(lang string) ApiGetQuoteRequest {
	r.lang = &lang
	return r
}

func (r ApiGetQuoteRequest) IncludePrePost(includePrePost bool) ApiGetQuoteRequest {
	r.includePrePost = &includePrePost
	return r
}

func (r ApiGetQuoteRequest) Fields(fields string) ApiGetQuoteRequest {
	r.fields = &fields
	return r
}

func (r ApiGetQuoteRequest) CorsDomain(corsDomain string) ApiGetQuoteRequest {
	r.corsDomain = &corsDomain
	return r
}

func (r ApiGetQuoteRequest) Execute() (*QuoteResponse, *http.Response, error) {
	return r.ApiService.GetQuoteExecute(r)
}

/*
GetQuote Returns quotes for the specified symbols

Returns quotes for the specified symbols

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQuoteRequest
*/
func (a *QuoteAPIService) GetQuote(ctx context.Context) ApiGetQuoteRequest {
	return ApiGetQuoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuoteResponse
func (a *QuoteAPIService) GetQuoteExecute(r ApiGetQuoteRequest) (*QuoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuoteAPIService.GetQuote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v7/finance/quote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbols == nil {
		return localVarReturnValue, nil, reportError("symbols is required and must be specified")
	}
	if r.crumb == nil {
		return localVarReturnValue, nil, reportError("crumb is required and must be specified")
	}
	if r.a1 == nil {
		return localVarReturnValue, nil, reportError("a1 is required and must be specified")
	}

	if r.formatted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "formatted", r.formatted, "")
	} else {
		var defaultValue bool = false
		r.formatted = &defaultValue
	}
	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	if r.lang != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lang", r.lang, "")
	}
	if r.includePrePost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includePrePost", r.includePrePost, "")
	} else {
		var defaultValue bool = false
		r.includePrePost = &defaultValue
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.corsDomain != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "corsDomain", r.corsDomain, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "symbols", r.symbols, "")
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
	localVarHTTPCookies := make(map[string]string)
	if *r.crumb != "" {
		localVarHTTPCookies["Crumb"] = *r.crumb
	}
	if *r.a1 != "" {
		localVarHTTPCookies["A1"] = *r.a1
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
