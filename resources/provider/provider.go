package provider

import (
	"embed"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/andrewthetechie/cq-provider-datadog/resources/services/monitors"
)

var (
	//go:embed migrations/*/*.sql
	providerMigrations embed.FS
	Version            = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version: Version,
		// CHANGEME: Change to your provider name
		Name:      "YourProviderName",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"datadog_monitors": monitors.Monitors(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
