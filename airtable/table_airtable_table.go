package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAirtableTable(tableName string) *plugin.Table {
	return &plugin.Table{
		Name:        toTableName(tableName),
		Description: "The " + tableName + " table.",
		List: &plugin.ListConfig{
			Hydrate: listTable(tableName),
			KeyColumns: plugin.OptionalColumns([]string{"query"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTable(tableName),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The record ID of the row."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the record was created."},
			{Name: "fields", Type: proto.ColumnType_JSON, Description: "The fields of the row."},
			{Name: "query", Type: proto.ColumnType_STRING, Description: "Filter string to [filterWithFormula](https://support.airtable.com/hc/en-us/articles/203255215).", Transform: transform.FromQual("query")},
		},
	}
}

func listTable(tableName string) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			return nil, err
		}
		airtableConfig := GetConfig(d.Connection)
		table := client.GetTable(*airtableConfig.DatabaseID, tableName)
		offset := ""
		queryFilter := ""

		if d.KeyColumnQuals["query"] != nil {
			queryFilter = d.KeyColumnQuals["query"].GetStringValue()
		}

		for {
			records, err := table.GetRecords().WithFilterFormula(queryFilter).WithOffset(offset).Do()
			if err != nil {
				if is404Error(err) {
					return nil, nil
				}
				return nil, err
			}
			for _, record := range records.Records {
				d.StreamListItem(ctx, record)
			}
			if d.QueryStatus.RowsRemaining(ctx) <= 0 {
				break
			}
			offset = records.Offset
			if offset == "" {
				break
			}
		}
		return nil, nil
	}
}

func getTable(tableName string) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		airtableConfig := GetConfig(d.Connection)
		table := client.GetTable(*airtableConfig.DatabaseID, tableName)
		if err != nil {
			return nil, err
		}
		quals := d.KeyColumnQuals
		id := quals["id"].GetStringValue()
		record, err := table.GetRecord(id)
		if err != nil {
			if is404Error(err) {
				return nil, nil
			}
			return nil, err
		}
		return record, nil
	}
}
