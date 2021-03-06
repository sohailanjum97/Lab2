/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// N2InterfaceAmfInfo - AMF N2 interface information
type N2InterfaceAmfInfo struct {

	Ipv4EndpointAddress []string `json:"ipv4EndpointAddress,omitempty"`

	Ipv6EndpointAddress []Ipv6Addr `json:"ipv6EndpointAddress,omitempty"`

	AmfName string `json:"amfName,omitempty"`
}
