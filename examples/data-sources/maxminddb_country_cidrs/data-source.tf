data "maxminddb_country_cidrs" "foo" {
  country {
    iso_code     = "GB"
    subdivisions = ["ENG"]
  }

  country {
    iso_code = "NL"
  }
}
