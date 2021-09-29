package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableAirtableTable(tableName string) *plugin.Table {
	return &plugin.Table{
		Name:        toTableName(tableName),
		Description: "The " + tableName + " table.",
		List: &plugin.ListConfig{
			Hydrate: listTable(tableName),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The record ID of the row."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the record was created."},
			{Name: "fields", Type: proto.ColumnType_JSON, Description: "The fields of the row."},
		},
	}
}

func listTable(tableName string) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		airtableConfig := GetConfig(d.Connection)
		table := client.GetTable(*airtableConfig.DatabaseID, tableName)
		if err != nil {
			return nil, err
		}
		offset := ""
		for {
			records, err := table.GetRecords().WithOffset(offset).Do()
			if err != nil {
				return nil, err
			}
			for _, record := range records.Records {
				d.StreamListItem(ctx, record)
			}
			offset = records.Offset
			if offset == "" {
				break
			}
		}
		return nil, nil
	}
}
