// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resource Domain-defined abstraction of a deposited 'work'. Resource's abstractions are describable for our domain’s purposes, i.e. for management needs within our system.
// swagger:model Resource
type Resource struct {

	// URI for the JSON-LD context definitions.
	// Required: true
	AtContext *string `json:"@context"`

	// The content type of the resource. Selected from an established set of values.
	// Required: true
	// Enum: [http://sdr.sul.stanford.edu/models/sdr3-object.jsonld http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld http://sdr.sul.stanford.edu/models/sdr3-book.jsonld http://sdr.sul.stanford.edu/models/sdr3-document.jsonld http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld http://sdr.sul.stanford.edu/models/sdr3-image.jsonld http://sdr.sul.stanford.edu/models/sdr3-page.jsonld http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld http://sdr.sul.stanford.edu/models/sdr3-map.jsonld http://sdr.sul.stanford.edu/models/sdr3-media.jsonld http://sdr.sul.stanford.edu/models/sdr3-track.jsonld http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld]
	AtType *string `json:"@type"`

	// access
	// Required: true
	Access *ResourceAccess `json:"access"`

	// administrative
	// Required: true
	Administrative *ResourceAdministrative `json:"administrative"`

	// Identifier retrieved from identification.sourceId that stands for analog or source identifier that this resource is a digital representation of.
	DedupeIdentifier string `json:"dedupeIdentifier,omitempty"`

	// The Agent (User, Group, Application, Department, other) that deposited the resource into TAQUITO.
	Depositor *Agent `json:"depositor,omitempty"`

	// Following version for the resource within TAQUITO.
	FollowingVersion string `json:"followingVersion,omitempty"`

	// identification
	// Required: true
	Identification *ResourceIdentification `json:"identification"`

	// Identifier for the resource within TAQUITO. UUID, unique for each new version of a TAQUITO resource.
	// Format: uuid
	Identifier strfmt.UUID `json:"identifier,omitempty"`

	// Primary processing label (can be same as title) for a resource.
	// Required: true
	Label *string `json:"label"`

	// permissions
	Permissions *ResourcePermissions `json:"permissions,omitempty"`

	// Preceding version for the resource within TAQUITO.
	PrecedingVersion string `json:"precedingVersion,omitempty"`

	// structural
	// Required: true
	Structural *ResourceStructural `json:"structural"`

	// Version for the resource within TAQUITO.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAtType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdministrative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepositor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStructural(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateAtContext(formats strfmt.Registry) error {

	if err := validate.Required("@context", "body", m.AtContext); err != nil {
		return err
	}

	return nil
}

var resourceTypeAtTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http://sdr.sul.stanford.edu/models/sdr3-object.jsonld","http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld","http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld","http://sdr.sul.stanford.edu/models/sdr3-book.jsonld","http://sdr.sul.stanford.edu/models/sdr3-document.jsonld","http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld","http://sdr.sul.stanford.edu/models/sdr3-image.jsonld","http://sdr.sul.stanford.edu/models/sdr3-page.jsonld","http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld","http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld","http://sdr.sul.stanford.edu/models/sdr3-map.jsonld","http://sdr.sul.stanford.edu/models/sdr3-media.jsonld","http://sdr.sul.stanford.edu/models/sdr3-track.jsonld","http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld","http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeAtTypePropEnum = append(resourceTypeAtTypePropEnum, v)
	}
}

const (

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ObjectJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-object.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ObjectJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-object.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr33dJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr33dJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-3d.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3AgreementJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3AgreementJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-agreement.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3BookJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-book.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3BookJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-book.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3DocumentJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-document.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3DocumentJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-document.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3GeoJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3GeoJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-geo.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ImageJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-image.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ImageJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-image.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PageJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-page.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PageJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-page.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PhotographJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3PhotographJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-photograph.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ManuscriptJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3ManuscriptJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-manuscript.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MapJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-map.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MapJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-map.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MediaJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-media.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3MediaJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-media.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3TrackJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-track.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3TrackJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-track.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveBinaryJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveBinaryJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-webarchive-binary.jsonld"

	// ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveSeedJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld"
	ResourceAtTypeHTTPSdrSulStanfordEduModelsSdr3WebarchiveSeedJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-webarchive-seed.jsonld"
)

// prop value enum
func (m *Resource) validateAtTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceTypeAtTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateAtType(formats strfmt.Registry) error {

	if err := validate.Required("@type", "body", m.AtType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAtTypeEnum("@type", "body", *m.AtType); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	if m.Access != nil {
		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateAdministrative(formats strfmt.Registry) error {

	if err := validate.Required("administrative", "body", m.Administrative); err != nil {
		return err
	}

	if m.Administrative != nil {
		if err := m.Administrative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("administrative")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateDepositor(formats strfmt.Registry) error {

	if swag.IsZero(m.Depositor) { // not required
		return nil
	}

	if m.Depositor != nil {
		if err := m.Depositor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("depositor")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.Identifier) { // not required
		return nil
	}

	if err := validate.FormatOf("identifier", "body", "uuid", m.Identifier.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateStructural(formats strfmt.Registry) error {

	if err := validate.Required("structural", "body", m.Structural); err != nil {
		return err
	}

	if m.Structural != nil {
		if err := m.Structural.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("structural")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceAccess Access Metadata for the resource.
// swagger:model ResourceAccess
type ResourceAccess struct {

	// Access level for the resource.
	// Required: true
	// Enum: [world stanford location-based citation-only dark]
	Access *string `json:"access"`

	// The human readable copyright statement that applies to the resource.
	Copyright string `json:"copyright,omitempty"`

	// Date when the resource is released from an embargo, if an embargo exists.
	// Format: date-time
	EmbargoReleaseDate strfmt.DateTime `json:"embargoReleaseDate,omitempty"`

	// The license governing reuse of the resource. Should be an IRI for known licenses (i.e. CC, RightsStatement.org URI, etc.).
	License string `json:"license,omitempty"`

	// The human readable reuse and reproduction statement that applies to the resource.
	ReuseAndReproductionStatement string `json:"reuseAndReproductionStatement,omitempty"`

	// License or terms of use governing reuse of the resource. Should be a text statement.
	TermsOfUse string `json:"termsOfUse,omitempty"`

	// Percentage of the resource that is visibility during an embargo period
	Visibility int64 `json:"visibility,omitempty"`
}

// Validate validates this resource access
func (m *ResourceAccess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmbargoReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceAccessTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["world","stanford","location-based","citation-only","dark"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceAccessTypeAccessPropEnum = append(resourceAccessTypeAccessPropEnum, v)
	}
}

const (

	// ResourceAccessAccessWorld captures enum value "world"
	ResourceAccessAccessWorld string = "world"

	// ResourceAccessAccessStanford captures enum value "stanford"
	ResourceAccessAccessStanford string = "stanford"

	// ResourceAccessAccessLocationBased captures enum value "location-based"
	ResourceAccessAccessLocationBased string = "location-based"

	// ResourceAccessAccessCitationOnly captures enum value "citation-only"
	ResourceAccessAccessCitationOnly string = "citation-only"

	// ResourceAccessAccessDark captures enum value "dark"
	ResourceAccessAccessDark string = "dark"
)

// prop value enum
func (m *ResourceAccess) validateAccessEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceAccessTypeAccessPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceAccess) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access"+"."+"access", "body", m.Access); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccessEnum("access"+"."+"access", "body", *m.Access); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAccess) validateEmbargoReleaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EmbargoReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("access"+"."+"embargoReleaseDate", "body", "date-time", m.EmbargoReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceAccess) UnmarshalBinary(b []byte) error {
	var res ResourceAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceAdministrative Administrative metadata for the resource.
// swagger:model ResourceAdministrative
type ResourceAdministrative struct {

	// When the resource in TAQUITO was created.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// If the resource has been deleted (but not purged).
	Deleted bool `json:"deleted,omitempty"`

	// Message describing why the object was deleted.
	GravestoneMessage string `json:"gravestoneMessage,omitempty"`

	// Pointer to the PURL/XML file that is a traditional representation of the metadata for the resource.
	IsDescribedBy string `json:"isDescribedBy,omitempty"`

	// When the resource in TAQUITO was last updated.
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// Administrative or Internal project this resource is a part of.
	PartOfProject string `json:"partOfProject,omitempty"`

	// The Agent (User, Group, Application, Department, other) that remediated a resource in TAQUITO.
	RemediatedBy []*Agent `json:"remediatedBy"`

	// If this resource should be sent to Preservation.
	// Required: true
	SdrPreserve *bool `json:"sdrPreserve"`
}

// Validate validates this resource administrative
func (m *ResourceAdministrative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemediatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSdrPreserve(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceAdministrative) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("administrative"+"."+"created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAdministrative) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("administrative"+"."+"lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAdministrative) validateRemediatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.RemediatedBy) { // not required
		return nil
	}

	for i := 0; i < len(m.RemediatedBy); i++ {
		if swag.IsZero(m.RemediatedBy[i]) { // not required
			continue
		}

		if m.RemediatedBy[i] != nil {
			if err := m.RemediatedBy[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("administrative" + "." + "remediatedBy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceAdministrative) validateSdrPreserve(formats strfmt.Registry) error {

	if err := validate.Required("administrative"+"."+"sdrPreserve", "body", m.SdrPreserve); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceAdministrative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceAdministrative) UnmarshalBinary(b []byte) error {
	var res ResourceAdministrative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceIdentification Identifying information for the resource.
// swagger:model ResourceIdentification
type ResourceIdentification struct {

	// Links to catalog records for the resource that the resource represents in whole or part.
	CatalogLinks []*ResourceIdentificationCatalogLinksItems0 `json:"catalogLinks"`

	// Digital Object Identifier (DOI) for the resource within this repository.
	Doi string `json:"doi,omitempty"`

	// Another object, either external or internal to the system (if duplication occurs) that is the same digital resource as this resource.
	SameAs string `json:"sameAs,omitempty"`

	// A source resource or object (perhaps but not necessarily analog or physical) that the resource is a digital representation of.
	SourceID string `json:"sourceId,omitempty"`
}

// Validate validates this resource identification
func (m *ResourceIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceIdentification) validateCatalogLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.CatalogLinks) { // not required
		return nil
	}

	for i := 0; i < len(m.CatalogLinks); i++ {
		if swag.IsZero(m.CatalogLinks[i]) { // not required
			continue
		}

		if m.CatalogLinks[i] != nil {
			if err := m.CatalogLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identification" + "." + "catalogLinks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceIdentification) UnmarshalBinary(b []byte) error {
	var res ResourceIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceIdentificationCatalogLinksItems0 resource identification catalog links items0
// swagger:model ResourceIdentificationCatalogLinksItems0
type ResourceIdentificationCatalogLinksItems0 struct {

	// Catalog that is the source of the linked record.
	// Required: true
	Catalog *string `json:"catalog"`

	// Record identifier that is unique within the context of the linked record's catalog.
	// Required: true
	CatalogRecordID *string `json:"catalogRecordId"`

	// If the linked record should be automatically updated when the resource descriptive metadata changes.
	DeliverMetadata bool `json:"deliverMetadata,omitempty"`

	// If the resource descriptive metadata should be automatically updated when the linked record changes.
	DeriveMetadata bool `json:"deriveMetadata,omitempty"`

	// Metadata schema of the linked record.
	RecordSchema string `json:"recordSchema,omitempty"`

	// Whether the linked record describes a resource that is broader than, equivalent to, or narrower than the resource the resource represents.
	// Enum: [broader equivalent narrower]
	RecordScope string `json:"recordScope,omitempty"`
}

// Validate validates this resource identification catalog links items0
func (m *ResourceIdentificationCatalogLinksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalogRecordID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceIdentificationCatalogLinksItems0) validateCatalog(formats strfmt.Registry) error {

	if err := validate.Required("catalog", "body", m.Catalog); err != nil {
		return err
	}

	return nil
}

func (m *ResourceIdentificationCatalogLinksItems0) validateCatalogRecordID(formats strfmt.Registry) error {

	if err := validate.Required("catalogRecordId", "body", m.CatalogRecordID); err != nil {
		return err
	}

	return nil
}

var resourceIdentificationCatalogLinksItems0TypeRecordScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["broader","equivalent","narrower"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceIdentificationCatalogLinksItems0TypeRecordScopePropEnum = append(resourceIdentificationCatalogLinksItems0TypeRecordScopePropEnum, v)
	}
}

const (

	// ResourceIdentificationCatalogLinksItems0RecordScopeBroader captures enum value "broader"
	ResourceIdentificationCatalogLinksItems0RecordScopeBroader string = "broader"

	// ResourceIdentificationCatalogLinksItems0RecordScopeEquivalent captures enum value "equivalent"
	ResourceIdentificationCatalogLinksItems0RecordScopeEquivalent string = "equivalent"

	// ResourceIdentificationCatalogLinksItems0RecordScopeNarrower captures enum value "narrower"
	ResourceIdentificationCatalogLinksItems0RecordScopeNarrower string = "narrower"
)

// prop value enum
func (m *ResourceIdentificationCatalogLinksItems0) validateRecordScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceIdentificationCatalogLinksItems0TypeRecordScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceIdentificationCatalogLinksItems0) validateRecordScope(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecordScopeEnum("recordScope", "body", m.RecordScope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceIdentificationCatalogLinksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceIdentificationCatalogLinksItems0) UnmarshalBinary(b []byte) error {
	var res ResourceIdentificationCatalogLinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourcePermissions Permissions Metadata for the resource.
// swagger:model ResourcePermissions
type ResourcePermissions struct {

	// Indicates if approval is required to deposit or manage the resource in TAQUITO.
	ApprovalRequired bool `json:"approvalRequired,omitempty"`

	// Agents who are required to approve deposit or management of this resource in TAQUITO.
	Approvers []*Agent `json:"approvers"`
}

// Validate validates this resource permissions
func (m *ResourcePermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprovers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcePermissions) validateApprovers(formats strfmt.Registry) error {

	if swag.IsZero(m.Approvers) { // not required
		return nil
	}

	for i := 0; i < len(m.Approvers); i++ {
		if swag.IsZero(m.Approvers[i]) { // not required
			continue
		}

		if m.Approvers[i] != nil {
			if err := m.Approvers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + "approvers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcePermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePermissions) UnmarshalBinary(b []byte) error {
	var res ResourcePermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceStructural Structural metadata for the resource.
// swagger:model ResourceStructural
type ResourceStructural struct {

	// Filesets that contain the digital representations (Files) of the resource.
	Contains []string `json:"contains"`

	// Agreement that covers the deposit of the resource into TAQUITO.
	// Required: true
	HasAgreement *string `json:"hasAgreement"`

	// has member
	HasMember *ResourceStructuralHasMemberTuple0 `json:"hasMember,omitempty"`

	// has member orders
	HasMemberOrders *ResourceStructuralHasMemberOrdersTuple0 `json:"hasMemberOrders,omitempty"`

	// An Annotation instance that applies to the resource.
	IsTargetOf string `json:"isTargetOf,omitempty"`
}

// Validate validates this resource structural
func (m *ResourceStructural) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasMemberOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceStructural) validateHasAgreement(formats strfmt.Registry) error {

	if err := validate.Required("structural"+"."+"hasAgreement", "body", m.HasAgreement); err != nil {
		return err
	}

	return nil
}

func (m *ResourceStructural) validateHasMember(formats strfmt.Registry) error {

	if swag.IsZero(m.HasMember) { // not required
		return nil
	}

	if m.HasMember != nil {
		if err := m.HasMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("structural" + "." + "hasMember")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceStructural) validateHasMemberOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.HasMemberOrders) { // not required
		return nil
	}

	if m.HasMemberOrders != nil {
		if err := m.HasMemberOrders.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("structural" + "." + "hasMemberOrders")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceStructural) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceStructural) UnmarshalBinary(b []byte) error {
	var res ResourceStructural
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceStructuralHasMemberOrdersTuple0 ResourceStructuralHasMemberOrdersTuple0 a representation of an anonymous Tuple type
// swagger:model ResourceStructuralHasMemberOrdersTuple0
type ResourceStructuralHasMemberOrdersTuple0 struct {

	// p0
	// Required: true
	P0 *Sequence `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (m *ResourceStructuralHasMemberOrdersTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements

	if len(stage1) > 0 {
		var dataP0 Sequence
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		m.P0 = &dataP0

	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (m ResourceStructuralHasMemberOrdersTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		m.P0,
	}

	return json.Marshal(data)
}

// Validate validates this resource structural has member orders tuple0
func (m *ResourceStructuralHasMemberOrdersTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceStructuralHasMemberOrdersTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", m.P0); err != nil {
		return err
	}

	if m.P0 != nil {
		if err := m.P0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("P0")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceStructuralHasMemberOrdersTuple0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceStructuralHasMemberOrdersTuple0) UnmarshalBinary(b []byte) error {
	var res ResourceStructuralHasMemberOrdersTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceStructuralHasMemberTuple0 ResourceStructuralHasMemberTuple0 a representation of an anonymous Tuple type
// swagger:model ResourceStructuralHasMemberTuple0
type ResourceStructuralHasMemberTuple0 struct {

	// p0
	// Required: true
	P0 *string `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (m *ResourceStructuralHasMemberTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements

	if len(stage1) > 0 {
		var dataP0 string
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		m.P0 = &dataP0

	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (m ResourceStructuralHasMemberTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		m.P0,
	}

	return json.Marshal(data)
}

// Validate validates this resource structural has member tuple0
func (m *ResourceStructuralHasMemberTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceStructuralHasMemberTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", m.P0); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceStructuralHasMemberTuple0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceStructuralHasMemberTuple0) UnmarshalBinary(b []byte) error {
	var res ResourceStructuralHasMemberTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}