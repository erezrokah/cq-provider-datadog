package monitor

import (
	"context"
	"fmt"

	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:        "datadog_monitors",
		Description: "Datadog Montiors",
		Resolver:    fetchDatadogMonitors,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"meta_id"}},
		Columns: []schema.Column{
			{
				Name:        "meta_id",
				Description: "The Meta_ID of this monitor, combines datadog account name from config with ID and monitor name to provide a unique PK.",
				Type:        schema.TypeString,
				Resolver:    resolveMetaID,
			},
			{
				Name:        "id",
				Description: "The ID of this monitor.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "The name of this monitor.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of this monitor.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created",
				Description: "Timestamp of when the monitor was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deleted",
				Description: "Timestamp of when the monitor was deleted",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "modified",
				Description: "Timestamp of when the monitor was last modified",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "creator_email",
				Description: "Email of the monitor's creator",
				Type:        schema.TypeString,
			},
			{
				Name:        "creator_handle",
				Description: "Handle of the monitor's creator",
				Type:        schema.TypeString,
			},
			{
				Name:        "creator_name",
				Description: "Name of the monitor's creator",
				Type:        schema.TypeString,
			},
			{
				Name:        "message",
				Description: "Monitor's message",
				Type:        schema.TypeString,
			},
			{
				Name:        "multi",
				Description: "If a monitor is a multi alert",
				Type:        schema.TypeBool,
			},
			{
				Name:        "options",
				Description: "Options for this monitor as a json blob",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "overall_state",
				Description: "Overall State of the monitor, string summary",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "State of the monitor as json",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "priority",
				Description: "Priority of the monitor",
				Type:        schema.TypeString,
			},
			{
				Name:        "query",
				Description: "Query of the monitor",
				Type:        schema.TypeString,
			},
			{
				Name:        "restricted_roles",
				Description: "Restricted roles of the monitor",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "Tags of the monitor",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchDatadogMonitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return nil
}

func resolveMetaID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	// TODO: Figure out multiplexing, for now, we're just getting one account working
	c := meta.(*client.Client)
	return diag.WrapError(resource.Set(col.Name, fmt.Sprintf("%s-%s-%s", c.Accounts[0].Name, resource.Get("id"), resource.Get("name"))))
}
