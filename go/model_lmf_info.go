/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LmfInfo struct {

	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`

	LmfId string `json:"lmfId,omitempty"`

	ServingAccessTypes []AccessType `json:"servingAccessTypes,omitempty"`

	ServingAnNodeTypes []AnNodeType `json:"servingAnNodeTypes,omitempty"`

	ServingRatTypes []RatType `json:"servingRatTypes,omitempty"`
}
