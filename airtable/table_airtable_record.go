package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAirtableRecord(tableName string) *plugin.Table {
	return &plugin.Table{
		Name:        toTableName(tableName),
		Description: "The " + tableName + " table.",
		List: &plugin.ListConfig{
			Hydrate:           listRecord(tableName),
			KeyColumns:        plugin.OptionalColumns([]string{"filter_formula"}),
			ShouldIgnoreError: isNotFoundError,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getRecord(tableName),
			ShouldIgnoreError: isNotFoundError,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The record ID of the row."},
			{Name: "created_time", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the record was created."},
			{Name: "fields", Type: proto.ColumnType_JSON, Description: "The fields of the row."},
			{Name: "filter_formula", Type: proto.ColumnType_STRING, Description: "The formula used to filter records. For more information see https://support.airtable.com/hc/en-us/articles/203255215.", Transform: transform.FromQual("filter_formula")},
		},
	}
}

func listRecord(tableName string) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("airtable_table.listTable", "connection_error", err)
			return nil, err
		}
		airtableConfig := GetConfig(d.Connection)
		table := client.GetTable(*airtableConfig.DatabaseID, tableName)
		offset := ""
		filterFormula := ""
		var maxResult *int64 = nil
		limit := d.QueryContext.Limit

		if limit != nil {
			if *limit < int64(100) {
				maxResult = limit
			}
		}

		if d.KeyColumnQuals["filter_formula"] != nil {
			filterFormula = d.KeyColumnQuals["filter_formula"].GetStringValue()
		}

		for {
			query := table.GetRecords().WithFilterFormula(filterFormula).WithOffset(offset)
			if maxResult != nil {
				query = query.MaxRecords(int(*maxResult))
			}

			records, err := query.Do()
			if err != nil {
				plugin.Logger(ctx).Error("airtable_table.listTable", err)
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

func getRecord(tableName string) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("airtable_table.getTable", "connection_error", err)
			return nil, err
		}
		airtableConfig := GetConfig(d.Connection)
		table := client.GetTable(*airtableConfig.DatabaseID, tableName)
		quals := d.KeyColumnQuals
		id := quals["id"].GetStringValue()
		record, err := table.GetRecord(id)
		if err != nil {
			plugin.Logger(ctx).Error("airtable_table.getTable", err)
			return nil, err
		}
		return record, nil
	}
}
