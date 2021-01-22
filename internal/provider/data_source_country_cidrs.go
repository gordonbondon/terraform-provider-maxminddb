package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/gordonbondon/maxminddb-cidrs/pkg/cidrs"
)

func dataCountryCIDRs() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieving network lists based on [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) country codes and subdivision codes.",

		ReadContext: dataCountryCIDRsRead,

		Schema: map[string]*schema.Schema{
			"country": {
				Description: "Country to find CIDRs for",
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_code": {
							Description: "`ISO 3166-2` country code",
							Type:        schema.TypeString,
							Required:    true,
						},
						"subdivisions": {
							Description: "List of `ISO 3166-2` subdivision codes for selected country",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"ip_address_version": {
				Description: "IP version, one of `IPV4` or `IPV6`",
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{
					"IPV4",
					"IPV6",
				}, false),
			},
			"cidrs": {
				Description: "List of network CIDRs retrieved for specified countries and subdivisions",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataCountryCIDRsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	dbPath := meta.(*maxMindDB).DBPath

	var diags diag.Diagnostics

	options := &cidrs.ListOptions{
		DBPath:    dbPath,
		Countries: expandCountry(d.Get("country").(*schema.Set).List()),
	}

	if v := d.Get("ip_address_version").(string); v != "" {
		switch v {
		case "IPV4":
			options.IPv4 = true
		case "IPV6":
			options.IPv6 = true
		}
	}

	cidrs, err := cidrs.List(options)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(cidrs) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "No networks found for provided country list",
			AttributePath: cty.Path{cty.GetAttrStep{Name: "country"}},
		})

		return diags
	}

	sort.Strings(cidrs)

	d.SetId(fmt.Sprintf("%d", hash(strings.Join(cidrs, ""))))
	if err := d.Set("cidrs", cidrs); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Failed setting cidrs value",
			AttributePath: cty.Path{cty.GetAttrStep{Name: "cidrs"}},
		})

		return diags
	}

	return diags
}
