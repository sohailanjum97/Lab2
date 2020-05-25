/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	//"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A CompleteStoredSearchDocumentApiController binds http requests to an api service and writes the service results to the http response
type CompleteStoredSearchDocumentApiController struct {
	service CompleteStoredSearchDocumentApiServicer
}

// NewCompleteStoredSearchDocumentApiController creates a default api controller
func NewCompleteStoredSearchDocumentApiController(s CompleteStoredSearchDocumentApiServicer) Router {
	return &CompleteStoredSearchDocumentApiController{service: s}
}

// Routes returns all of the api route for the CompleteStoredSearchDocumentApiController
func (c *CompleteStoredSearchDocumentApiController) Routes() Routes {
	return Routes{
		{
			"RetrieveCompleteSearch",
			strings.ToUpper("Get"),
			"/nnrf-disc/v1/searches/{searchId}/complete",
			c.RetrieveCompleteSearch,
		},
	}
}

// RetrieveCompleteSearch -
func (c *CompleteStoredSearchDocumentApiController) RetrieveCompleteSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchId := params["searchId"]
	acceptEncoding := r.Header.Get("acceptEncoding")
	result, err := c.service.RetrieveCompleteSearch(searchId, acceptEncoding)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
