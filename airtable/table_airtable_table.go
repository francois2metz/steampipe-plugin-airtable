package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAirtableTable() *plugin.Table {
	return &plugin.Table{
		Name:        "airtable_table",
		Description: "The list of tables of a base.",
		List: &plugin.ListConfig{
			Hydrate:    listTable,
			KeyColumns: plugin.SingleColumn("base_id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "base_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("base_id"),
				Description: "ID of the base.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "A unique identifier for the table.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the table.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of the table.",
			},
			{
				Name:        "fields",
				Type:        proto.ColumnType_JSON,
				Description: "Fields of the table.",
			},
		},
	}
}

func listTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("airtable_table.listTable", "connection_error", err)
		return nil, err
	}
	baseID := d.KeyColumnQuals["base_id"].GetStringValue()
	baseSchema := client.GetBaseSchema(baseID)

	result, err := baseSchema.Do()
	if err != nil {
		plugin.Logger(ctx).Error("airtable_table.listTable", err)
		return nil, err
	}
	for _, table := range result.Tables {
		d.StreamListItem(ctx, table)
	}
	return nil, nil
}
