/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.6.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request, 
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	GetCardanoTip(http.ResponseWriter, *http.Request)
	GetOrders(http.ResponseWriter, *http.Request)
	GetProducts(http.ResponseWriter, *http.Request)
	GetStatus(http.ResponseWriter, *http.Request)
	GetWallet(http.ResponseWriter, *http.Request)
	GetWorkflowMinterNvla(http.ResponseWriter, *http.Request)
	PostCardanoTransaction(http.ResponseWriter, *http.Request)
	PostOrders(http.ResponseWriter, *http.Request)
	PostWorkflowMinterNvla(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	GetCardanoTip(context.Context) (ImplResponse, error)
	GetOrders(context.Context, string) (ImplResponse, error)
	GetProducts(context.Context, string, string, string) (ImplResponse, error)
	GetStatus(context.Context) (ImplResponse, error)
	GetWallet(context.Context, string) (ImplResponse, error)
	GetWorkflowMinterNvla(context.Context) (ImplResponse, error)
	PostCardanoTransaction(context.Context, CardanoTransaction) (ImplResponse, error)
	PostOrders(context.Context, Order) (ImplResponse, error)
	PostWorkflowMinterNvla(context.Context, MinterInfo) (ImplResponse, error)
}
