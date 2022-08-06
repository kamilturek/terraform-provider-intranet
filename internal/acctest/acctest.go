package acctest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	intranet "github.com/kamilturek/terraform-provider-intranet/internal"
)

func PreCheck(t *testing.T) {
	t.Helper()

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
