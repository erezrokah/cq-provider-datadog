package provider_test

import (
	"testing"

	"github.com/andrewthetechie/cq-provider-datadog/resources/provider"

	"github.com/cloudquery/cq-provider-sdk/migration"
)

func TestMigrations(t *testing.T) {
	migration.RunMigrationsTest(t, provider.Provider(), []string{"latest"})
}
