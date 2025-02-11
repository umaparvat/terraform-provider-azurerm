package dns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"

// AaaaRecord an AAAA record.
type AaaaRecord struct {
	// Ipv6Address - The IPv6 address of this AAAA record.
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}

// ARecord an A record.
type ARecord struct {
	// Ipv4Address - The IPv4 address of this A record.
	Ipv4Address *string `json:"ipv4Address,omitempty"`
}

// CaaRecord a CAA record.
type CaaRecord struct {
	// Flags - The flags for this CAA record as an integer between 0 and 255.
	Flags *int32 `json:"flags,omitempty"`
	// Tag - The tag for this CAA record.
	Tag *string `json:"tag,omitempty"`
	// Value - The value for this CAA record.
	Value *string `json:"value,omitempty"`
}

// CloudError an error response from the service.
type CloudError struct {
	// Error - Cloud error body.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from the service.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// CnameRecord a CNAME record.
type CnameRecord struct {
	// Cname - The canonical name for this CNAME record.
	Cname *string `json:"cname,omitempty"`
}

// MxRecord an MX record.
type MxRecord struct {
	// Preference - The preference value for this MX record.
	Preference *int32 `json:"preference,omitempty"`
	// Exchange - The domain name of the mail host for this MX record.
	Exchange *string `json:"exchange,omitempty"`
}

// NsRecord an NS record.
type NsRecord struct {
	// Nsdname - The name server name for this NS record.
	Nsdname *string `json:"nsdname,omitempty"`
}

// PtrRecord a PTR record.
type PtrRecord struct {
	// Ptrdname - The PTR target domain name for this PTR record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// RecordSet describes a DNS record set (a collection of DNS records with the same name and type).
type RecordSet struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The ID of the record set.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the record set.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the record set.
	Type *string `json:"type,omitempty"`
	// Etag - The etag of the record set.
	Etag *string `json:"etag,omitempty"`
	// RecordSetProperties - The properties of the record set.
	*RecordSetProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for RecordSet.
func (rs RecordSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rs.Etag != nil {
		objectMap["etag"] = rs.Etag
	}
	if rs.RecordSetProperties != nil {
		objectMap["properties"] = rs.RecordSetProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for RecordSet struct.
func (rs *RecordSet) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rs.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rs.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rs.Type = &typeVar
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				rs.Etag = &etag
			}
		case "properties":
			if v != nil {
				var recordSetProperties RecordSetProperties
				err = json.Unmarshal(*v, &recordSetProperties)
				if err != nil {
					return err
				}
				rs.RecordSetProperties = &recordSetProperties
			}
		}
	}

	return nil
}

// RecordSetListResult the response to a record set List operation.
type RecordSetListResult struct {
	autorest.Response `json:"-"`
	// Value - Information about the record sets in the response.
	Value *[]RecordSet `json:"value,omitempty"`
	// NextLink - READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for RecordSetListResult.
func (rslr RecordSetListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rslr.Value != nil {
		objectMap["value"] = rslr.Value
	}
	return json.Marshal(objectMap)
}

// RecordSetListResultIterator provides access to a complete listing of RecordSet values.
type RecordSetListResultIterator struct {
	i    int
	page RecordSetListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RecordSetListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *RecordSetListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RecordSetListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RecordSetListResultIterator) Response() RecordSetListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RecordSetListResultIterator) Value() RecordSet {
	if !iter.page.NotDone() {
		return RecordSet{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RecordSetListResultIterator type.
func NewRecordSetListResultIterator(page RecordSetListResultPage) RecordSetListResultIterator {
	return RecordSetListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rslr RecordSetListResult) IsEmpty() bool {
	return rslr.Value == nil || len(*rslr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rslr RecordSetListResult) hasNextLink() bool {
	return rslr.NextLink != nil && len(*rslr.NextLink) != 0
}

// recordSetListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rslr RecordSetListResult) recordSetListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !rslr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rslr.NextLink)))
}

// RecordSetListResultPage contains a page of RecordSet values.
type RecordSetListResultPage struct {
	fn   func(context.Context, RecordSetListResult) (RecordSetListResult, error)
	rslr RecordSetListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RecordSetListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rslr)
		if err != nil {
			return err
		}
		page.rslr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RecordSetListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RecordSetListResultPage) NotDone() bool {
	return !page.rslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RecordSetListResultPage) Response() RecordSetListResult {
	return page.rslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RecordSetListResultPage) Values() []RecordSet {
	if page.rslr.IsEmpty() {
		return nil
	}
	return *page.rslr.Value
}

// Creates a new instance of the RecordSetListResultPage type.
func NewRecordSetListResultPage(cur RecordSetListResult, getNextPage func(context.Context, RecordSetListResult) (RecordSetListResult, error)) RecordSetListResultPage {
	return RecordSetListResultPage{
		fn:   getNextPage,
		rslr: cur,
	}
}

// RecordSetProperties represents the properties of the records in the record set.
type RecordSetProperties struct {
	// Metadata - The metadata attached to the record set.
	Metadata map[string]*string `json:"metadata"`
	// TTL - The TTL (time-to-live) of the records in the record set.
	TTL *int64 `json:"TTL,omitempty"`
	// Fqdn - READ-ONLY; Fully qualified domain name of the record set.
	Fqdn *string `json:"fqdn,omitempty"`
	// ProvisioningState - READ-ONLY; provisioning State of the record set.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// TargetResource - A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResource `json:"targetResource,omitempty"`
	// ARecords - The list of A records in the record set.
	ARecords *[]ARecord `json:"ARecords,omitempty"`
	// AaaaRecords - The list of AAAA records in the record set.
	AaaaRecords *[]AaaaRecord `json:"AAAARecords,omitempty"`
	// MxRecords - The list of MX records in the record set.
	MxRecords *[]MxRecord `json:"MXRecords,omitempty"`
	// NsRecords - The list of NS records in the record set.
	NsRecords *[]NsRecord `json:"NSRecords,omitempty"`
	// PtrRecords - The list of PTR records in the record set.
	PtrRecords *[]PtrRecord `json:"PTRRecords,omitempty"`
	// SrvRecords - The list of SRV records in the record set.
	SrvRecords *[]SrvRecord `json:"SRVRecords,omitempty"`
	// TxtRecords - The list of TXT records in the record set.
	TxtRecords *[]TxtRecord `json:"TXTRecords,omitempty"`
	// CnameRecord - The CNAME record in the  record set.
	CnameRecord *CnameRecord `json:"CNAMERecord,omitempty"`
	// SoaRecord - The SOA record in the record set.
	SoaRecord *SoaRecord `json:"SOARecord,omitempty"`
	// CaaRecords - The list of CAA records in the record set.
	CaaRecords *[]CaaRecord `json:"caaRecords,omitempty"`
}

// MarshalJSON is the custom marshaler for RecordSetProperties.
func (rsp RecordSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rsp.Metadata != nil {
		objectMap["metadata"] = rsp.Metadata
	}
	if rsp.TTL != nil {
		objectMap["TTL"] = rsp.TTL
	}
	if rsp.TargetResource != nil {
		objectMap["targetResource"] = rsp.TargetResource
	}
	if rsp.ARecords != nil {
		objectMap["ARecords"] = rsp.ARecords
	}
	if rsp.AaaaRecords != nil {
		objectMap["AAAARecords"] = rsp.AaaaRecords
	}
	if rsp.MxRecords != nil {
		objectMap["MXRecords"] = rsp.MxRecords
	}
	if rsp.NsRecords != nil {
		objectMap["NSRecords"] = rsp.NsRecords
	}
	if rsp.PtrRecords != nil {
		objectMap["PTRRecords"] = rsp.PtrRecords
	}
	if rsp.SrvRecords != nil {
		objectMap["SRVRecords"] = rsp.SrvRecords
	}
	if rsp.TxtRecords != nil {
		objectMap["TXTRecords"] = rsp.TxtRecords
	}
	if rsp.CnameRecord != nil {
		objectMap["CNAMERecord"] = rsp.CnameRecord
	}
	if rsp.SoaRecord != nil {
		objectMap["SOARecord"] = rsp.SoaRecord
	}
	if rsp.CaaRecords != nil {
		objectMap["caaRecords"] = rsp.CaaRecords
	}
	return json.Marshal(objectMap)
}

// RecordSetUpdateParameters parameters supplied to update a record set.
type RecordSetUpdateParameters struct {
	// RecordSet - Specifies information about the record set being updated.
	RecordSet *RecordSet `json:"RecordSet,omitempty"`
}

// Resource common properties of an Azure Resource Manager resource
type Resource struct {
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceReference represents a single Azure resource and its referencing DNS records.
type ResourceReference struct {
	// DNSResources - A list of dns Records
	DNSResources *[]SubResource `json:"dnsResources,omitempty"`
	// TargetResource - A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResource `json:"targetResource,omitempty"`
}

// ResourceReferenceRequest represents the properties of the Dns Resource Reference Request.
type ResourceReferenceRequest struct {
	// ResourceReferenceRequestProperties - The properties of the Resource Reference Request.
	*ResourceReferenceRequestProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceReferenceRequest.
func (rrr ResourceReferenceRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rrr.ResourceReferenceRequestProperties != nil {
		objectMap["properties"] = rrr.ResourceReferenceRequestProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceReferenceRequest struct.
func (rrr *ResourceReferenceRequest) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var resourceReferenceRequestProperties ResourceReferenceRequestProperties
				err = json.Unmarshal(*v, &resourceReferenceRequestProperties)
				if err != nil {
					return err
				}
				rrr.ResourceReferenceRequestProperties = &resourceReferenceRequestProperties
			}
		}
	}

	return nil
}

// ResourceReferenceRequestProperties represents the properties of the Dns Resource Reference Request.
type ResourceReferenceRequestProperties struct {
	// TargetResources - A list of references to azure resources for which referencing dns records need to be queried.
	TargetResources *[]SubResource `json:"targetResources,omitempty"`
}

// ResourceReferenceResult represents the properties of the Dns Resource Reference Result.
type ResourceReferenceResult struct {
	autorest.Response `json:"-"`
	// ResourceReferenceResultProperties - The result of dns resource reference request. Returns a list of dns resource references for each of the azure resource in the request.
	*ResourceReferenceResultProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceReferenceResult.
func (rrr ResourceReferenceResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rrr.ResourceReferenceResultProperties != nil {
		objectMap["properties"] = rrr.ResourceReferenceResultProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceReferenceResult struct.
func (rrr *ResourceReferenceResult) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var resourceReferenceResultProperties ResourceReferenceResultProperties
				err = json.Unmarshal(*v, &resourceReferenceResultProperties)
				if err != nil {
					return err
				}
				rrr.ResourceReferenceResultProperties = &resourceReferenceResultProperties
			}
		}
	}

	return nil
}

// ResourceReferenceResultProperties the result of dns resource reference request. Returns a list of dns
// resource references for each of the azure resource in the request.
type ResourceReferenceResultProperties struct {
	// DNSResourceReferences - The result of dns resource reference request. A list of dns resource references for each of the azure resource in the request
	DNSResourceReferences *[]ResourceReference `json:"dnsResourceReferences,omitempty"`
}

// SoaRecord an SOA record.
type SoaRecord struct {
	// Host - The domain name of the authoritative name server for this SOA record.
	Host *string `json:"host,omitempty"`
	// Email - The email contact for this SOA record.
	Email *string `json:"email,omitempty"`
	// SerialNumber - The serial number for this SOA record.
	SerialNumber *int64 `json:"serialNumber,omitempty"`
	// RefreshTime - The refresh value for this SOA record.
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// RetryTime - The retry time for this SOA record.
	RetryTime *int64 `json:"retryTime,omitempty"`
	// ExpireTime - The expire time for this SOA record.
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// MinimumTTL - The minimum value for this SOA record. By convention this is used to determine the negative caching duration.
	MinimumTTL *int64 `json:"minimumTTL,omitempty"`
}

// SrvRecord an SRV record.
type SrvRecord struct {
	// Priority - The priority value for this SRV record.
	Priority *int32 `json:"priority,omitempty"`
	// Weight - The weight value for this SRV record.
	Weight *int32 `json:"weight,omitempty"`
	// Port - The port value for this SRV record.
	Port *int32 `json:"port,omitempty"`
	// Target - The target domain name for this SRV record.
	Target *string `json:"target,omitempty"`
}

// SubResource a reference to a another resource
type SubResource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
}

// TxtRecord a TXT record.
type TxtRecord struct {
	// Value - The text value of this TXT record.
	Value *[]string `json:"value,omitempty"`
}

// Zone describes a DNS zone.
type Zone struct {
	autorest.Response `json:"-"`
	// Etag - The etag of the zone.
	Etag *string `json:"etag,omitempty"`
	// ZoneProperties - The properties of the zone.
	*ZoneProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Zone.
func (z Zone) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if z.Etag != nil {
		objectMap["etag"] = z.Etag
	}
	if z.ZoneProperties != nil {
		objectMap["properties"] = z.ZoneProperties
	}
	if z.Location != nil {
		objectMap["location"] = z.Location
	}
	if z.Tags != nil {
		objectMap["tags"] = z.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Zone struct.
func (z *Zone) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				z.Etag = &etag
			}
		case "properties":
			if v != nil {
				var zoneProperties ZoneProperties
				err = json.Unmarshal(*v, &zoneProperties)
				if err != nil {
					return err
				}
				z.ZoneProperties = &zoneProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				z.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				z.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				z.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				z.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				z.Tags = tags
			}
		}
	}

	return nil
}

// ZoneListResult the response to a Zone List or ListAll operation.
type ZoneListResult struct {
	autorest.Response `json:"-"`
	// Value - Information about the DNS zones.
	Value *[]Zone `json:"value,omitempty"`
	// NextLink - READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ZoneListResult.
func (zlr ZoneListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if zlr.Value != nil {
		objectMap["value"] = zlr.Value
	}
	return json.Marshal(objectMap)
}

// ZoneListResultIterator provides access to a complete listing of Zone values.
type ZoneListResultIterator struct {
	i    int
	page ZoneListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ZoneListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ZoneListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ZoneListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ZoneListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ZoneListResultIterator) Response() ZoneListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ZoneListResultIterator) Value() Zone {
	if !iter.page.NotDone() {
		return Zone{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ZoneListResultIterator type.
func NewZoneListResultIterator(page ZoneListResultPage) ZoneListResultIterator {
	return ZoneListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (zlr ZoneListResult) IsEmpty() bool {
	return zlr.Value == nil || len(*zlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (zlr ZoneListResult) hasNextLink() bool {
	return zlr.NextLink != nil && len(*zlr.NextLink) != 0
}

// zoneListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (zlr ZoneListResult) zoneListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !zlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(zlr.NextLink)))
}

// ZoneListResultPage contains a page of Zone values.
type ZoneListResultPage struct {
	fn  func(context.Context, ZoneListResult) (ZoneListResult, error)
	zlr ZoneListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ZoneListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ZoneListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.zlr)
		if err != nil {
			return err
		}
		page.zlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ZoneListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ZoneListResultPage) NotDone() bool {
	return !page.zlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ZoneListResultPage) Response() ZoneListResult {
	return page.zlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ZoneListResultPage) Values() []Zone {
	if page.zlr.IsEmpty() {
		return nil
	}
	return *page.zlr.Value
}

// Creates a new instance of the ZoneListResultPage type.
func NewZoneListResultPage(cur ZoneListResult, getNextPage func(context.Context, ZoneListResult) (ZoneListResult, error)) ZoneListResultPage {
	return ZoneListResultPage{
		fn:  getNextPage,
		zlr: cur,
	}
}

// ZoneProperties represents the properties of the zone.
type ZoneProperties struct {
	// MaxNumberOfRecordSets - READ-ONLY; The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int64 `json:"maxNumberOfRecordSets,omitempty"`
	// MaxNumberOfRecordsPerRecordSet - READ-ONLY; The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordsPerRecordSet *int64 `json:"maxNumberOfRecordsPerRecordSet,omitempty"`
	// NumberOfRecordSets - READ-ONLY; The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets *int64 `json:"numberOfRecordSets,omitempty"`
	// NameServers - READ-ONLY; The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers *[]string `json:"nameServers,omitempty"`
	// ZoneType - The type of this DNS zone (Public or Private). Possible values include: 'Public', 'Private'
	ZoneType ZoneType `json:"zoneType,omitempty"`
	// RegistrationVirtualNetworks - A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
	RegistrationVirtualNetworks *[]SubResource `json:"registrationVirtualNetworks,omitempty"`
	// ResolutionVirtualNetworks - A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
	ResolutionVirtualNetworks *[]SubResource `json:"resolutionVirtualNetworks,omitempty"`
}

// MarshalJSON is the custom marshaler for ZoneProperties.
func (zp ZoneProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if zp.ZoneType != "" {
		objectMap["zoneType"] = zp.ZoneType
	}
	if zp.RegistrationVirtualNetworks != nil {
		objectMap["registrationVirtualNetworks"] = zp.RegistrationVirtualNetworks
	}
	if zp.ResolutionVirtualNetworks != nil {
		objectMap["resolutionVirtualNetworks"] = zp.ResolutionVirtualNetworks
	}
	return json.Marshal(objectMap)
}

// ZonesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type ZonesDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ZonesClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *ZonesDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for ZonesDeleteFuture.Result.
func (future *ZonesDeleteFuture) result(client ZonesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.ZonesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("dns.ZonesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ZoneUpdate describes a request to update a DNS zone.
type ZoneUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ZoneUpdate.
func (zu ZoneUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if zu.Tags != nil {
		objectMap["tags"] = zu.Tags
	}
	return json.Marshal(objectMap)
}
