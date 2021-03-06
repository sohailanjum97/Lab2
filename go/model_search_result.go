/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SearchResult - Contains the list of NF Profiles returned in a Discovery response
type SearchResult struct {

	ValidityPeriod int32 `json:"validityPeriod,omitempty"`

	NfInstances []NfProfile `json:"nfInstances"`

	SearchId string `json:"searchId,omitempty"`

	NumNfInstComplete int32 `json:"numNfInstComplete,omitempty"`

	PreferredSearch PreferredSearch `json:"preferredSearch,omitempty"`

	NrfSupportedFeatures string `json:"nrfSupportedFeatures,omitempty"`
}
