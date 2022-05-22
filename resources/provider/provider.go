package provider

import (
	"embed"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/andrewthetechie/cq-provider-datadog/resources/services/monitor"
)

var (
	//go:embed migrations/*/*.sql
	providerMigrations embed.FS
	Version            = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:   Version,
		Name:      "Datadog",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"datadog_monitors": monitor.Monitors(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
