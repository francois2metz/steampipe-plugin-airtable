package airtable

import (
	"context"

	"github.com/mehanizm/airtable"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAirtableRecord(ctx context.Context, base *airtable.Base, table *airtable.TableSchema) *plugin.Table {
	columns := []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "The record ID of the row",
		},
		{
			Name:        "created_time",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Time when the record was created.",
		},
		{
			Name:        "filter_formula",
			Type:        proto.ColumnType_STRING,
			Description: "The formula used to filter records. For more information see https://support.airtable.com/hc/en-us/articles/203255215.",
			Transform:   transform.FromQual("filter_formula"),
		},
	}
	for _, field := range table.Fields {
		columnName := findColumnNameAvailable(columns, field.Name)
		columns = append(columns, &plugin.Column{
			Name:        columnName,
			Type:        airtableFieldTypeToSteampipeType(field.Type),
			Description: field.Description,
			Transform:   transform.FromField("Fields." + field.Name),
		})
	}
	return &plugin.Table{
		Name:        toTableName(base.ID, table.Name),
		Description: "The " + table.Name + " table from the base " + base.Name + ".",
		List: &plugin.ListConfig{
			Hydrate:           listRecord(base.ID, table),
			KeyColumns:        plugin.OptionalColumns([]string{"filter_formula"}),
			ShouldIgnoreError: isNotFoundError,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getRecord(base.ID, table),
			ShouldIgnoreError: isNotFoundError,
		},
		Columns: columns,
	}
}

func findColumnNameAvailable(columns []*plugin.Column, fieldName string) string {
	if hasThisColumnNameAlready(columns, fieldName) {
		return findColumnNameAvailable(columns, "_"+ fieldName)
	}
	return fieldName
}

func hasThisColumnNameAlready(columns []*plugin.Column, fieldName string) bool {
	for _, column := range columns {
		if column.Name == fieldName {
			return true
		}
	}
	return false
}

func airtableFieldTypeToSteampipeType(airtableType string) proto.ColumnType {
	switch airtableType {
	case "number", "autoNumber", "count":
		return proto.ColumnType_DOUBLE
	case "date", "lastModifiedTime", "createdTime":
		return proto.ColumnType_TIMESTAMP
	case "multipleLookupValues", "multipleRecordLinks":
		return proto.ColumnType_JSON
	default:
		return proto.ColumnType_STRING
	}
}

func listRecord(databaseID string, table *airtable.TableSchema) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("airtable_record.listRecord", "connection_error", err)
			return nil, err
		}
		table := client.GetTable(databaseID, table.Name)
		offset := ""
		filterFormula := ""
		var maxResult *int64 = nil
		limit := d.QueryContext.Limit

		if limit != nil {
			if *limit < int64(100) {
				maxResult = limit
			}
		}

		if d.EqualsQuals["filter_formula"] != nil {
			filterFormula = d.EqualsQuals["filter_formula"].GetStringValue()
		}

		for {
			query := table.GetRecords().WithFilterFormula(filterFormula).WithOffset(offset)
			if maxResult != nil {
				query = query.MaxRecords(int(*maxResult))
			}

			records, err := query.Do()
			if err != nil {
				plugin.Logger(ctx).Error("airtable_record.listRecord", err)
				return nil, err
			}
			for _, record := range records.Records {
				d.StreamListItem(ctx, record)
			}
			if d.RowsRemaining(ctx) <= 0 {
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

func getRecord(databaseID string, table *airtable.TableSchema) func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("airtable_record.getRecord", "connection_error", err)
			return nil, err
		}
		table := client.GetTable(databaseID, table.Name)
		quals := d.EqualsQuals
		id := quals["id"].GetStringValue()
		record, err := table.GetRecord(id)
		if err != nil {
			plugin.Logger(ctx).Error("airtable_record.getRecord", err)
			return nil, err
		}
		return record, nil
	}
}
