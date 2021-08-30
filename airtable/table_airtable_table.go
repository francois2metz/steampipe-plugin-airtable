package airtable

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableAirtableTable(tableName string) *plugin.Table {
	return &plugin.Table{
		Name:        "airtable_"+  strings.ToLower(tableName),
		Description: "Table " + tableName,
		List: &plugin.ListConfig{
			Hydrate: listTable(tableName),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "the record ID"},
			{Name: "fields", Type: proto.ColumnType_JSON, Description: "fields of the table"},
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
