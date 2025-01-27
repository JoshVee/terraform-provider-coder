package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"

	"github.com/coder/terraform-provider-coder/provider"
)

func TestWorkspace(t *testing.T) {
	t.Setenv("CODER_WORKSPACE_OWNER", "owner123")
	t.Setenv("CODER_WORKSPACE_OWNER_EMAIL", "owner123@example.com")
	t.Setenv("CODER_WORKSPACE_OWNER_SESSION_TOKEN", "abc123")
	t.Setenv("CODER_WORKSPACE_TEMPLATE_ID", "templateID")
	t.Setenv("CODER_WORKSPACE_TEMPLATE_NAME", "template123")

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"coder": provider.New(),
		},
		IsUnitTest: true,
		Steps: []resource.TestStep{{
			Config: `
			provider "coder" {
				url = "https://example.com:8080"
			}
			data "coder_workspace" "me" {
			}`,
			Check: func(state *terraform.State) error {
				require.Len(t, state.Modules, 1)
				require.Len(t, state.Modules[0].Resources, 1)
				resource := state.Modules[0].Resources["data.coder_workspace.me"]
				require.NotNil(t, resource)

				attribs := resource.Primary.Attributes
				value := attribs["transition"]
				require.NotNil(t, value)
				t.Log(value)
				require.Equal(t, "8080", attribs["access_port"])
				require.Equal(t, "owner123", attribs["owner"])
				require.Equal(t, "owner123@example.com", attribs["owner_email"])
				require.Equal(t, "abc123", attribs["owner_session_token"])
				require.Equal(t, "templateID", attribs["template_id"])
				require.Equal(t, "template123", attribs["template_name"])
				return nil
			},
		}},
	})
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"coder": provider.New(),
		},
		IsUnitTest: true,
		Steps: []resource.TestStep{{
			Config: `
			provider "coder" {
				url = "https://example.com:8080"
			}
			data "coder_workspace" "me" {
			}`,
			Check: func(state *terraform.State) error {
				require.Len(t, state.Modules, 1)
				require.Len(t, state.Modules[0].Resources, 1)
				resource := state.Modules[0].Resources["data.coder_workspace.me"]
				require.NotNil(t, resource)

				attribs := resource.Primary.Attributes
				value := attribs["transition"]
				require.NotNil(t, value)
				t.Log(value)
				require.Equal(t, "https://example.com:8080", attribs["access_url"])
				require.Equal(t, "owner123", attribs["owner"])
				require.Equal(t, "owner123@example.com", attribs["owner_email"])
				require.Equal(t, "templateID", attribs["template_id"])
				require.Equal(t, "template123", attribs["template_name"])
				return nil
			},
		}},
	})
}
