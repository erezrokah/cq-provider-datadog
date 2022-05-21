package monitors

import (
	"context"

	// "github.com/andrewthetechie/cq-provider-datadog/client"
	// "github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:                 "datadog_monitors",
		Description:          "Datadog Montiors",
		Resolver:             fetchDatadogMonitors,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "The ID of this monitor.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner",
				Description: "The AWS account ID of the topic's owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy",
				Description: "The JSON serialization of the topic's access control policy.",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "delivery_policy",
				Description:   "The JSON serialization of the topic's delivery policy.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "display_name",
				Description: "The human-readable name used in the From field for notifications to email and email-json endpoints.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subscriptions_confirmed",
				Description: "The number of confirmed subscriptions for the topic.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subscriptions_deleted",
				Description: "The number of deleted subscriptions for the topic.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subscriptions_pending",
				Description: "The number of subscriptions pending confirmation for the topic.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "effective_delivery_policy",
				Description: "The JSON serialization of the effective delivery policy, taking system defaults into account.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "fifo_topic",
				Description: "When this is set to true, a FIFO topic is created.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "content_based_deduplication",
				Description: "Enables content-based deduplication for FIFO topics.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "kms_master_key_id",
				Description: "The ID of an AWS managed customer master key (CMK) for Amazon SNS or a custom CMK",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The topic's ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TopicArn"),
			},
			{
				Name:        "tags",
				Description: "Topic tags.",
				Type:        schema.TypeJSON,
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