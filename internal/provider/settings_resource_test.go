// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/supabase/cli/pkg/api"
	"github.com/supabase/terraform-provider-supabase/examples"
	"gopkg.in/h2non/gock.v1"
)

func TestAccSettingsResource(t *testing.T) {
	// Setup mock api
	defer gock.OffAll()
	// Step 1: create
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("10s"),
		})
	gock.New("https://api.supabase.com").
		Put("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("10s"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusOK).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs:   Ptr([]string{"0.0.0.0/0"}),
				DbAllowedCidrsV6: Ptr([]string{"::/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Post("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusCreated).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs:   Ptr([]string{"0.0.0.0/0"}),
				DbAllowedCidrsV6: Ptr([]string{"::/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           1000,
		})
	gock.New("https://api.supabase.com").
		Patch("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           1000,
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
		})
	gock.New("https://api.supabase.com").
		Patch("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/pgbouncer").
		Reply(http.StatusOK).
		JSON(api.V1PgbouncerConfigResponse{
			ConnectionString:        Ptr("postgresql://user:pass@pooler.supabase.com:6543/postgres"),
			DefaultPoolSize:         Ptr(float32(15)),
			IgnoreStartupParameters: Ptr(""),
			MaxClientConn:           Ptr(float32(200)),
			PoolMode:                Ptr(api.Transaction),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/pgbouncer").
		Reply(http.StatusOK).
		JSON(api.V1PgbouncerConfigResponse{
			ConnectionString:        Ptr("postgresql://user:pass@pooler.supabase.com:6543/postgres"),
			DefaultPoolSize:         Ptr(float32(15)),
			IgnoreStartupParameters: Ptr(""),
			MaxClientConn:           Ptr(float32(200)),
			PoolMode:                Ptr(api.Transaction),
		})
	// Step 2: read
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("10s"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("10s"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusOK).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs:   Ptr([]string{"0.0.0.0/0"}),
				DbAllowedCidrsV6: Ptr([]string{"::/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusOK).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs:   Ptr([]string{"0.0.0.0/0"}),
				DbAllowedCidrsV6: Ptr([]string{"::/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           1000,
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           1000,
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/pgbouncer").
		Reply(http.StatusOK).
		JSON(api.V1PgbouncerConfigResponse{
			ConnectionString:        Ptr("postgresql://user:pass@pooler.supabase.com:6543/postgres"),
			DefaultPoolSize:         Ptr(float32(15)),
			IgnoreStartupParameters: Ptr(""),
			MaxClientConn:           Ptr(float32(200)),
			PoolMode:                Ptr(api.Transaction),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/pgbouncer").
		Reply(http.StatusOK).
		JSON(api.V1PgbouncerConfigResponse{
			ConnectionString:        Ptr("postgresql://user:pass@pooler.supabase.com:6543/postgres"),
			DefaultPoolSize:         Ptr(float32(15)),
			IgnoreStartupParameters: Ptr(""),
			MaxClientConn:           Ptr(float32(200)),
			PoolMode:                Ptr(api.Transaction),
		})
	// Step 3: update
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("10s"),
		})
	gock.New("https://api.supabase.com").
		Put("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("20s"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/postgres").
		Reply(http.StatusOK).
		JSON(api.PostgresConfigResponse{
			StatementTimeout: Ptr("20s"),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusOK).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs:   Ptr([]string{"0.0.0.0/0"}),
				DbAllowedCidrsV6: Ptr([]string{"::/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Post("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusCreated).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs: Ptr([]string{"0.0.0.0/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/network-restrictions").
		Reply(http.StatusOK).
		JSON(api.NetworkRestrictionsResponse{
			Config: api.NetworkRestrictionsRequest{
				DbAllowedCidrs: Ptr([]string{"0.0.0.0/0"}),
			},
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           1000,
		})
	gock.New("https://api.supabase.com").
		Patch("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           100,
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/postgrest").
		Reply(http.StatusOK).
		JSON(api.V1PostgrestConfigResponse{
			DbExtraSearchPath: "public,extensions",
			DbSchema:          "public,storage,graphql_public",
			MaxRows:           100,
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
			JwtExp:  Ptr(float32(3600)),
		})
	gock.New("https://api.supabase.com").
		Patch("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
			JwtExp:  Ptr(float32(1800)),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/auth").
		Reply(http.StatusOK).
		JSON(api.AuthConfigResponse{
			SiteUrl: Ptr("http://localhost:3000"),
			JwtExp:  Ptr(float32(1800)),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/pgbouncer").
		Reply(http.StatusOK).
		JSON(api.V1PgbouncerConfigResponse{
			ConnectionString:        Ptr("postgresql://user:pass@pooler.supabase.com:6543/postgres"),
			DefaultPoolSize:         Ptr(float32(15)),
			IgnoreStartupParameters: Ptr(""),
			MaxClientConn:           Ptr(float32(200)),
			PoolMode:                Ptr(api.Transaction),
		})
	gock.New("https://api.supabase.com").
		Get("/v1/projects/mayuaycdtijbctgqbycg/config/database/pgbouncer").
		Reply(http.StatusOK).
		JSON(api.V1PgbouncerConfigResponse{
			ConnectionString:        Ptr("postgresql://user:pass@pooler.supabase.com:6543/postgres"),
			DefaultPoolSize:         Ptr(float32(15)),
			IgnoreStartupParameters: Ptr(""),
			MaxClientConn:           Ptr(float32(200)),
			PoolMode:                Ptr(api.Transaction),
		})
	// Run test
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: examples.SettingsResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("supabase_settings.production", "id", "mayuaycdtijbctgqbycg"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "supabase_settings.production",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccSettingsResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("supabase_settings.production", "api"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

const testAccSettingsResourceConfig = `
resource "supabase_settings" "production" {
  project_ref = "mayuaycdtijbctgqbycg"

  database = jsonencode({
    statement_timeout = "20s"
  })

  network = jsonencode({
	restrictions = ["0.0.0.0/0"]
  })

  api = jsonencode({
	db_schema            = "public,storage,graphql_public"
    db_extra_search_path = "public,extensions"
	max_rows             = 100
  })

  auth = jsonencode({
    site_url = "http://localhost:3000"
    jwt_exp  = 1800
  })

  # storage = jsonencode({
  #   file_size_limit = "50MB"
  # })

  # pooler = jsonencode({
  #   default_pool_size         = 15
  #   ignore_startup_parameters = ""
  #   max_client_conn           = 200
  #   pool_mode                 = "transaction"
  # })
}
`
