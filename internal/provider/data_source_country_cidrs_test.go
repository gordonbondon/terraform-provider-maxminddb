package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCountryCIDRs_Basic(t *testing.T) {
	resName := "data.maxminddb_country_cidrs.foo"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCountryCIDRs_Create(testDB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "cidrs.#", "26"),
					resource.TestCheckResourceAttr(
						resName, "cidrs.0", "2.125.160.216/29"),
				),
			},
			{
				Config: testAccDataSourceCountryCIDRs_Subdivisions(testDB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "cidrs.#", "5"),
					resource.TestCheckResourceAttr(
						resName, "cidrs.0", "2.125.160.216/29"),
				),
			},
			{
				Config: testAccDataSourceCountryCIDRs_Multiple(testDB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "cidrs.#", "10"),
					resource.TestCheckResourceAttr(
						resName, "cidrs.0", "216.160.83.56/29"),
				),
			},
			{
				Config: testAccDataSourceCountryCIDRs_IPv6(testDB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "cidrs.#", "21"),
					resource.TestCheckResourceAttr(
						resName, "cidrs.0", "2a02:d3c0::/29"),
				),
			},
		},
	})
}

func testAccDataSourceCountryCIDRs_Create(dbPath string) string {
	return fmt.Sprintf(`
provider "maxminddb" {
  db_path = "%s"
}

data "maxminddb_country_cidrs" "foo" {
  country {
    iso_code = "GB"
  }
}`, dbPath)
}

func testAccDataSourceCountryCIDRs_Subdivisions(dbPath string) string {
	return fmt.Sprintf(`
provider "maxminddb" {
  db_path = "%s"
}

data "maxminddb_country_cidrs" "foo" {
  country {
    iso_code = "GB"
    subdivisions = ["ENG"]
  }
}`, dbPath)
}

func testAccDataSourceCountryCIDRs_Multiple(dbPath string) string {
	return fmt.Sprintf(`
provider "maxminddb" {
  db_path = "%s"
}

data "maxminddb_country_cidrs" "foo" {
  country {
    iso_code = "US"
  }

  country {
    iso_code = "NL"
  }
}`, dbPath)
}

func testAccDataSourceCountryCIDRs_IPv6(dbPath string) string {
	return fmt.Sprintf(`
provider "maxminddb" {
  db_path = "%s"
}

data "maxminddb_country_cidrs" "foo" {
  ip_address_version = "IPV6"
  country {
    iso_code = "GB"
  }
}`, dbPath)
}
