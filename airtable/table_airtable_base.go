package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAirtableBase() *plugin.Table {
	return &plugin.Table{
		Name:        "airtable_base",
		Description: "The list of airtable bases.",
		List: &plugin.ListConfig{
			Hydrate: listBase,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "A unique identifier for a base.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the base.",
			},
			{
				Name:        "permission_level",
				Type:        proto.ColumnType_STRING,
				Description: "Your permissions: none, read, comment, edit, create.",
			},
		},
	}
}

func listBase(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("airtable_base.listBase", "connection_error", err)
		return nil, err
	}
	bases := client.GetBases()
	offset := ""

	for {
		query := bases.WithOffset(offset)

		result, err := query.Do()
		if err != nil {
			plugin.Logger(ctx).Error("airtable_base.listBase", err)
			return nil, err
		}
		for _, base := range result.Bases {
			d.StreamListItem(ctx, base)
		}
		if d.RowsRemaining(ctx) <= 0 {
			break
		}
		offset = result.Offset
		if offset == "" {
			break
		}
	}
	return nil, nil
}
