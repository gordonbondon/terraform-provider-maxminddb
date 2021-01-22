package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"db_path": {
					Description: "path to mmdb file. Can be also set via `MAXMIND_DB_PATH` env variable",
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("MAXMIND_DB_PATH", ""),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"maxminddb_country_cidrs": dataCountryCIDRs(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type maxMindDB struct {
	DBPath string
}

func configure(version string, p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		dbPath := d.Get("db_path").(string)

		return &maxMindDB{DBPath: dbPath}, nil
	}
}
