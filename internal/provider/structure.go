package provider

import (
	"github.com/gordonbondon/maxminddb-cidrs/pkg/cidrs"
)

func expandCountry(config []interface{}) []cidrs.Country {
	result := make([]cidrs.Country, 0)

	for _, i := range config {
		c := i.(map[string]interface{})
		country := cidrs.Country{ISOCode: c["iso_code"].(string)}

		if v := c["subdivisions"].([]interface{}); len(v) > 0 {
			sub := make([]string, len(v))
			for i, s := range v {
				sub[i] = s.(string)
			}
			country.Subdivisions = sub
		}

		result = append(result, country)
	}

	return result
}
