package acctest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kamilturek/terraform-provider-intranet/intranet"
)

func PreCheck(t *testing.T) {
	if v := os.Getenv("INTRANET_SESSION_ID"); v == "" {
		t.Fatal("INTRANET_SESSION_ID must be set for acceptance tests")
	}
}

var Provider *schema.Provider
var Providers map[string]*schema.Provider

func init() {
	Provider = intranet.Provider()
	Providers = map[string]*schema.Provider{
		"intranet": Provider,
	}
}
