package geostore

type StoreInterface interface {
	CountryCreate(country *Country) error
	CountryDelete(country *Country) error
	CountryDeleteByID(countryID string) error
	CountryFindByID(countryID string) (*Country, error)
	CountryFindByIso2(iso2Code string) (*Country, error)
	CountryList(options CountryQueryOptions) ([]Country, error)
	CountrySoftDelete(discount *Country) error
	CountrySoftDeleteByID(discountID string) error
	CountryUpdate(country *Country) error
	TimezoneCreate(timezone *Timezone) error
	TimezoneList(options TimezoneQueryOptions) ([]Timezone, error)
}
