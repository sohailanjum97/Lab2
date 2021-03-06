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
	"time"
)

// NfService - Information of a given NF Service Instance; it is part of the NFProfile of an NF Instance discovered by the NRF
type NfService struct {

	ServiceInstanceId string `json:"serviceInstanceId"`

	ServiceName ServiceName `json:"serviceName"`

	Versions []NfServiceVersion `json:"versions"`

	Scheme UriScheme `json:"scheme"`

	NfServiceStatus NfServiceStatus `json:"nfServiceStatus"`

	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty"`

	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`

	ApiPrefix string `json:"apiPrefix,omitempty"`

	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`

	Load int32 `json:"load,omitempty"`

	LoadTimeStamp time.Time `json:"loadTimeStamp,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	ChfServiceInfo ChfServiceInfo `json:"chfServiceInfo,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	NfServiceSetIdList []string `json:"nfServiceSetIdList,omitempty"`

	SNssais []Snssai `json:"sNssais,omitempty"`

	PerPlmnSnssaiList []PlmnSnssai `json:"perPlmnSnssaiList,omitempty"`

	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId string `json:"vendorId,omitempty"`

	SupportedVendorSpecificFeatures map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty"`

	Oauth2Required bool `json:"oauth2Required,omitempty"`
}
