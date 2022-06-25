package intranet_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kamilturek/intranet"
	"github.com/kamilturek/terraform-provider-intranet/intranet/acctest"
)

func TestAccHourEntry_basic(t *testing.T) {
	rName := "intranet_hour_entry.test"
	now := time.Now().Format(intranet.DateFormat)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckHourEntryDestroy(now),
		Steps: []resource.TestStep{
			{
				Config: testAccHourEntry_basic(now),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHourEntryExists(rName, now),
					resource.TestCheckResourceAttr(rName, "date", now),
					resource.TestCheckResourceAttr(rName, "description", "test description"),
					resource.TestCheckResourceAttr(rName, "project_id", "422"),
					resource.TestCheckNoResourceAttr(rName, "ticket_id"),
					resource.TestCheckResourceAttr(rName, "time", "1.5"),
				),
			},
		},
	})
}

func testAccCheckHourEntryDestroy(date string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acctest.Provider.Meta().(*intranet.Client)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "intranet_hour_entry" {
				continue
			}

			input := &intranet.GetHourEntriesInput{
				Date: date,
			}

			output, err := client.GetHourEntries(input)
			if err != nil {
				return err
			}

			for _, e := range output.Entries {
				if strconv.Itoa(e.ID) == rs.Primary.ID {
					return fmt.Errorf("Hour Entry (%s) still exists.", rs.Primary.ID)
				}
			}
		}

		return nil
	}
}

func testAccCheckHourEntryExists(resourceName, date string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Hour Entry ID is not set.")
		}

		client := acctest.Provider.Meta().(*intranet.Client)

		input := &intranet.GetHourEntriesInput{
			Date: date,
		}

		output, err := client.GetHourEntries(input)
		if err != nil {
			return err
		}

		for _, e := range output.Entries {
			if strconv.Itoa(e.ID) == rs.Primary.ID {
				return nil
			}
		}

		return fmt.Errorf("Hour Entry (%s) not found.", rs.Primary.ID)
	}
}

func testAccHourEntry_basic(date string) string {
	return fmt.Sprintf(`
resource "intranet_hour_entry" "test" {
  date        = %[1]q
  description = "test description"
  project_id  = 422
  time        = 1.5
}
`, date)
}
