package geostore

import "github.com/gouniverse/sb"

// SQLCreateTable returns a SQL string for creating the country table
func (st *Store) sqlCountryTableCreate() string {
	sql := sb.NewBuilder(st.dbDriverName).
		Table(st.countryTableName).
		Column(sb.Column{
			Name:       "id",
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   "status",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "iso2_code",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 2,
		}).
		Column(sb.Column{
			Name:   "iso3_code",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 3,
		}).
		Column(sb.Column{
			Name:   "name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 255,
		}).
		Column(sb.Column{
			Name:   "continent",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "phone_prefix",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name: "created_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "updated_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "deleted_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}

func (st *Store) sqlTimezoneTableCreate() string {
	sql := sb.NewBuilder(st.dbDriverName).
		Table(st.timezoneTableName).
		Column(sb.Column{
			Name:       "id",
			Type:       sb.COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		}).
		Column(sb.Column{
			Name:   "status",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 20,
		}).
		Column(sb.Column{
			Name:   "timezone",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "zone_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "global_name",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 100,
		}).
		Column(sb.Column{
			Name:   "country_code",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name:   "offset",
			Type:   sb.COLUMN_TYPE_STRING,
			Length: 50,
		}).
		Column(sb.Column{
			Name: "created_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "updated_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		Column(sb.Column{
			Name: "deleted_at",
			Type: sb.COLUMN_TYPE_DATETIME,
		}).
		CreateIfNotExists()

	return sql
}
