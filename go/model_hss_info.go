/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// HssInfo - Information of an HSS NF Instance
type HssInfo struct {

	GroupId string `json:"groupId,omitempty"`

	ImsRanges []ImsiRange `json:"imsRanges,omitempty"`
}
