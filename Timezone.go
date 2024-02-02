package geostore

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

// == CLASS =================================================================

type Timezone struct {
	dataobject.DataObject
}

// == CONSTRUCTORS ==========================================================

func NewTimezone() *Timezone {
	tz := &Timezone{}
	tz.SetID(uid.HumanUid())
	tz.SetStatus(TIMEZONE_STATUS_ACTIVE)
	tz.SetTimezone("")
	tz.SetZoneName("")
	tz.SetGlobalName("")
	tz.SetCountryCode("")
	tz.SetOffset("")
	tz.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	tz.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	tz.SetDeletedAt(sb.NULL_DATETIME)
	return tz
}

func NewTimezoneFromExistingData(data map[string]string) *Timezone {
	tz := &Timezone{}
	tz.Hydrate(data)
	return tz
}

// == SETTERS AND GETTERS ===================================================

func (o *Timezone) CreatedAt() string {
	return o.Get("created_at")
}

func (o *Timezone) CreatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *Timezone) SetCreatedAt(createdAt string) *Timezone {
	o.Set("created_at", createdAt)
	return o
}

func (o *Timezone) CountryCode() string {
	return o.Get("country_code")
}

func (o *Timezone) SetCountryCode(countryCode string) *Timezone {
	o.Set("country_code", countryCode)
	return o
}

func (o *Timezone) DeletedAt() string {
	return o.Get("deleted_at")
}

func (o *Timezone) SetDeletedAt(deletedAt string) *Timezone {
	o.Set("deleted_at", deletedAt)
	return o
}

func (o *Timezone) GlobalName() string {
	return o.Get("global_name")
}

func (o *Timezone) SetGlobalName(globalName string) *Timezone {
	o.Set("global_name", globalName)
	return o
}

func (o *Timezone) Offset() string {
	return o.Get("offset")
}

func (o *Timezone) SetOffset(offset string) *Timezone {
	o.Set("offset", offset)
	return o
}

func (o *Timezone) Status() string {
	return o.Get("status")
}

func (o *Timezone) SetStatus(status string) *Timezone {
	o.Set("status", status)
	return o
}

func (o *Timezone) Timezone() string {
	return o.Get("timezone")
}

func (o *Timezone) SetTimezone(timezone string) *Timezone {
	o.Set("timezone", timezone)
	return o
}

func (o *Timezone) UpdatedAt() string {
	return o.Get("updated_at")
}

func (o *Timezone) UpdatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.UpdatedAt(), carbon.UTC)
}

func (o *Timezone) SetUpdatedAt(updatedAt string) *Timezone {
	o.Set("updated_at", updatedAt)
	return o
}

func (o *Timezone) ZoneName() string {
	return o.Get("zone_name")
}

func (o *Timezone) SetZoneName(zoneName string) *Timezone {
	o.Set("zone_name", zoneName)
	return o
}
