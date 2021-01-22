// +build tools

package tools

import (
	// document generation
	_ "github.com/bflad/tfproviderdocs"
	_ "github.com/bflad/tfproviderlint/cmd/tfproviderlint"
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/katbyte/terrafmt"
)
