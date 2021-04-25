/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"GetCardanoTip",
			strings.ToUpper("Get"),
			"/novellia/cardano/tip",
			c.GetCardanoTip,
		},
		{
			"GetOrders",
			strings.ToUpper("Get"),
			"/novellia/orders",
			c.GetOrders,
		},
		{
			"GetProducts",
			strings.ToUpper("Get"),
			"/novellia/products",
			c.GetProducts,
		},
		{
			"GetStatus",
			strings.ToUpper("Get"),
			"/novellia/status",
			c.GetStatus,
		},
		{
			"GetWallet",
			strings.ToUpper("Get"),
			"/novellia/wallet/{wallet_address}",
			c.GetWallet,
		},
		{
			"GetWorkflowMinterNvla",
			strings.ToUpper("Get"),
			"/novellia/workflow/minter/nvla",
			c.GetWorkflowMinterNvla,
		},
		{
			"PostCardanoTransaction",
			strings.ToUpper("Post"),
			"/novellia/cardano/transaction",
			c.PostCardanoTransaction,
		},
		{
			"PostOrders",
			strings.ToUpper("Post"),
			"/novellia/orders",
			c.PostOrders,
		},
		{
			"PostWorkflowMinterNvla",
			strings.ToUpper("Post"),
			"/novellia/workflow/minter/nvla",
			c.PostWorkflowMinterNvla,
		},
	}
}

// GetCardanoTip - Your GET endpoint
func (c *DefaultApiController) GetCardanoTip(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.GetCardanoTip(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetOrders - Your GET endpoint
func (c *DefaultApiController) GetOrders(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	orderId := query.Get("order_id")
	result, err := c.service.GetOrders(r.Context(), orderId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetProducts - Your GET endpoint
func (c *DefaultApiController) GetProducts(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	marketId := query.Get("market_id")
	organizationId := query.Get("organization_id")
	productId := query.Get("product_id")
	result, err := c.service.GetProducts(r.Context(), marketId, organizationId, productId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetStatus - Your GET endpoint
func (c *DefaultApiController) GetStatus(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.GetStatus(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetWallet - Your GET endpoint
func (c *DefaultApiController) GetWallet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	walletAddress := params["wallet_address"]
	result, err := c.service.GetWallet(r.Context(), walletAddress)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetWorkflowMinterNvla - 
func (c *DefaultApiController) GetWorkflowMinterNvla(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.GetWorkflowMinterNvla(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// PostCardanoTransaction - 
func (c *DefaultApiController) PostCardanoTransaction(w http.ResponseWriter, r *http.Request) { 
	cardanoTransaction := &CardanoTransaction{}
	if err := json.NewDecoder(r.Body).Decode(&cardanoTransaction); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.PostCardanoTransaction(r.Context(), *cardanoTransaction)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// PostOrders - 
func (c *DefaultApiController) PostOrders(w http.ResponseWriter, r *http.Request) { 
	orderCreated := &OrderCreated{}
	if err := json.NewDecoder(r.Body).Decode(&orderCreated); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.PostOrders(r.Context(), *orderCreated)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// PostWorkflowMinterNvla - 
func (c *DefaultApiController) PostWorkflowMinterNvla(w http.ResponseWriter, r *http.Request) { 
	minterInfo := &MinterInfo{}
	if err := json.NewDecoder(r.Body).Decode(&minterInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.PostWorkflowMinterNvla(r.Context(), *minterInfo)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
