package geostore

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
)

// == CLASS =================================================================

type Country struct {
	dataobject.DataObject
}

// == CONSTRUCTORS ==========================================================

func NewCountry() *Country {
	country := &Country{}
	country.SetID(uid.HumanUid())
	country.SetStatus(COUNTRY_STATUS_ACTIVE)
	country.SetName("")
	country.SetIsoCode2("")
	country.SetIsoCode3("")
	country.SetContinent("")
	country.SetPhonePrefix("")
	country.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	country.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	country.SetDeletedAt(sb.NULL_DATETIME)
	return country
}

func NewCountryFromExistingData(data map[string]string) *Country {
	country := &Country{}
	country.Hydrate(data)
	return country
}

// == SETTERS AND GETTERS ===================================================

func (o *Country) Continent() string {
	return o.Get("continent")
}

func (o *Country) SetContinent(continent string) *Country {
	o.Set("continent", continent)
	return o
}

func (o *Country) CreatedAt() string {
	return o.Get("created_at")
}

func (o *Country) SetCreatedAt(createdAt string) *Country {
	o.Set("created_at", createdAt)
	return o
}

func (o *Country) DeletedAt() string {
	return o.Get("deleted_at")
}

func (o *Country) SetDeletedAt(deletedAt string) *Country {
	o.Set("deleted_at", deletedAt)
	return o
}

func (o *Country) IsoCode2() string {
	return o.Get("iso2_code")
}

func (o *Country) SetIsoCode2(isoCode2 string) *Country {
	o.Set("iso2_code", isoCode2)
	return o
}

func (o *Country) IsoCode3() string {
	return o.Get("iso3_code")
}

func (o *Country) SetIsoCode3(isoCode3 string) *Country {
	o.Set("iso3_code", isoCode3)
	return o
}

func (o *Country) Name() string {
	return o.Get("name")
}

func (o *Country) SetName(name string) *Country {
	o.Set("name", name)
	return o
}

func (o *Country) PhonePrefix() string {
	return o.Get("phone_prefix")
}

func (o *Country) SetPhonePrefix(phonePrefix string) *Country {
	o.Set("phone_prefix", phonePrefix)
	return o
}

func (o *Country) Status() string {
	return o.Get("status")
}

func (o *Country) SetStatus(status string) *Country {
	o.Set("status", status)
	return o
}

func (o *Country) UpdatedAt() string {
	return o.Get("updated_at")
}

func (o *Country) SetUpdatedAt(updatedAt string) *Country {
	o.Set("updated_at", updatedAt)
	return o
}
