package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testDB = "../../test-data/test-data/GeoIP2-City-Test.mmdb"

var providerFactories = map[string]func() (*schema.Provider, error){
	"maxminddb": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	t.Helper()

	_, err := os.Stat(testDB)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("test mmdb %s file does not exist, update git submodules", testDB)
		} else {
			t.Fatalf("err: %s", err)
		}
	}
}
