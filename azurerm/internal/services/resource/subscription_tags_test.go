package resource_test

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
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
	// "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
	//"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
)

type SubscriptionTags struct {
}

func TestSubscriptionTag_basic(t *testing.T) {
	if os.Getenv("ARM_SUBSCRIPTION_ID") == "" {
		t.Skip("skipping tests - no subscription ID data provided")
	}
	data := acceptance.BuildTestData(t, "azurerm_subscription_tags", "test")
	r := SubscriptionTags{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicConfig(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.ImportStep(),
	})
}

func TestSubscriptionTag_requiresTagImport(t *testing.T) {
	if os.Getenv("ARM_SUBSCRIPTION_ID") == "" {
		t.Skip("skipping tests - no subscription ID data provided")
	}
	data := acceptance.BuildTestData(t, "azurerm_subscription_tags", "test")
	r := SubscriptionTags{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicConfig(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r)),
		},
		data.RequiresImportErrorStep(r.requiresImportConfig),
	})
}

func TestSubscriptionTags_updateWithTags(t *testing.T) {
	if os.Getenv("ARM_SUBSCRIPTION_ID") == "" {
		t.Skip("skipping tests - no subscription ID data provided")
	}
	data := acceptance.BuildTestData(t, "azurerm_subscription_tags", "test")

	r := SubscriptionTags{}
	assert := check.That(data.ResourceName)
	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicConfig(data),
			Check: resource.ComposeTestCheckFunc(
				assert.ExistsInAzure(r),
				assert.Key("tags.%").HasValue("2"),
				assert.Key("tags.cost_center").HasValue("MSFT"),
				assert.Key("tags.environment").HasValue("Production"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withTagsUpdatedConfig(data),
			Check: resource.ComposeTestCheckFunc(
				assert.ExistsInAzure(r),
				assert.Key("tags.%").HasValue("1"),
				assert.Key("tags.environment").HasValue("staging"),
			),
		},
		data.ImportStep(),
	})
}

// func (t SubscriptionTags) Destroy(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
// 	subscriptionId := state.Attributes["subscription_id"]
// 	atags := state.Attributes["tags"]
// 	resource_tags := resources.Tags{
// 		Tags: tags.Expand(atags),
// 	}
// 	parameters := resources.TagsPatchResource{Operation: "Delete", Properties: &resource_tags}

// 	tagsClient := client.Resource.TagsClient
// 	deleteFuture, err := tagsClient.UpdateAtScope(ctx, "subscriptions/"+subscriptionId, parameters)
// 	if err != nil {
// 		return nil, fmt.Errorf("deleting Resource Group %q: %+v", subscriptionId, err)
// 	}

// 	if deleteFuture.IsHTTPStatus(200) != true {
// 		return nil, fmt.Errorf("waiting for deletion of subscriptionTags %q: %+v", subscriptionId, err)
// 	}

// 	return utils.Bool(true), nil
// }

func (t SubscriptionTags) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	subscriptionId := state.Attributes["subscription_id"]
	atags := state.Attributes["tags"]
	fmt.Printf("tags: %T\n %s\n", atags, atags)
	fmt.Println("subscription id", subscriptionId, "tags", atags, "\n ")
	resp, err := client.Resource.TagsClient.GetAtScope(ctx, "subscriptions/"+subscriptionId)
	if err != nil {
		return nil, fmt.Errorf("retrieving tags from subscription %q: %+v", subscriptionId, err)
	}
	fmt.Println(resp.Properties)
	return utils.Bool(resp.Properties != nil), nil
}

func (t SubscriptionTags) basicConfig(data acceptance.TestData) string {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_subscription_tags" "test" {
  subscription_id = "%s"
  tags = {
    environment = "Production"
    cost_center = "MSFT"
  }
}
`, subscriptionId)
}

func (r SubscriptionTags) requiresImportConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_subscription_tags" "import" {
  subscription_id = azurerm_subscription_tags.test.subscription_id
  tags = azurerm_subscription_tags.test.tags
}
`, r.basicConfig(data))
}

func (t SubscriptionTags) withTagsUpdatedConfig(data acceptance.TestData) string {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_subscription_tags" "test" {
  subscription_id = "%s"

  tags = {
    environment = "staging"
  }
}
`, subscriptionId)
}
