package subscription_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/subscription/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type SubscriptionResourceTag struct{}

func TestAccSubscriptionResourceTag_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_subscription", "test")
	r := SubscriptionResourceTag{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicEnrollmentAccount(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.ImportStep(),
	})
}

func TestAccSubscriptionResourceTag_basicWithTag(t *testing.T) {

	data := acceptance.BuildTestData(t, "azurerm_subscription", "test")
	r := SubscriptionResourceTag{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.withTagsConfig(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.ImportStep(),
	})
}

func TestAccSubscriptionResourceTag_requiresTagImport(t *testing.T) {

	data := acceptance.BuildTestData(t, "azurerm_subscription", "test")
	r := SubscriptionResourceTag{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicEnrollmentAccount(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccSubscriptionResourceTag_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_subscription", "test")
	r := SubscriptionResourceTag{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicEnrollmentAccount(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.ImportStep(),
		{
			Config: r.basicEnrollmentAccountUpdate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.ImportStep(),
	})
}
func TestAccSubscriptionResourceTag_updateWithTags(t *testing.T) {

	data := acceptance.BuildTestData(t, "azurerm_subscription", "test")
	r := SubscriptionResourceTag{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.withTagsConfig(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("tags.%").HasValue("2"),
				check.That(data.ResourceName).Key("tags.cost_center").HasValue("MSFT"),
				check.That(data.ResourceName).Key("tags.environment").HasValue("Production"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withTagsUpdatedConfig(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("tags.%").HasValue("1"),
				check.That(data.ResourceName).Key("tags.environment").HasValue("staging"),
			),
		},
		data.ImportStep(),
	})
}
func (SubscriptionResourceTag) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.SubscriptionAliasID(state.ID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Subscription.AliasClient.Get(ctx, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving Subscription Alias %q: %+v", id.Name, err)
	}

	return utils.Bool(true), nil
}

// TODO - Need Env vars in CI for Billing Account and Enrollment Account - Testing disabled for now
func (SubscriptionResourceTag) basicEnrollmentAccount(data acceptance.TestData) string {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_subscription" "test" {
  alias             = "testAcc-%[1]d"
  subscription_name = "testAccSubscription %[1]d"
  subscription_id = "%s"

}
`, data.RandomInteger, subscriptionId)
}

func (SubscriptionResourceTag) basicEnrollmentAccountUpdate(data acceptance.TestData) string {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_subscription" "test" {
  alias             = "testAcc-%[1]d"
  subscription_name = "testAccSubscription Renamed %[1]d"
  subscription_id = "%s"

}
`, data.RandomInteger, subscriptionId)
}

func (r SubscriptionResourceTag) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_subscription" "import" {
  alias             = azurerm_subscription.test.alias
  subscription_name = azurerm_subscription.test.subscription_name
  subscription_id = azurerm_subscription.test.subscription_id
}
`, r.basicEnrollmentAccount(data))
}

func (t SubscriptionResourceTag) withTagsConfig(data acceptance.TestData) string {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_subscription" "test" {
  alias             = "testAcc-%[1]d"
  subscription_name = "testAccSubscription %[1]d"
  subscription_id = "%s"
  tags = {
    environment = "Production"
    cost_center = "MSFT"
  }
}
`, data.RandomInteger, subscriptionId)
}

func (t SubscriptionResourceTag) withTagsUpdatedConfig(data acceptance.TestData) string {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_subscription" "test" {
  alias             = "testAcc-%[1]d"
  subscription_name = "testAccSubscription Renamed %[1]d"
  subscription_id = "%s"



  tags = {
    environment = "staging"
  }
}
`, data.RandomInteger, subscriptionId)
}
