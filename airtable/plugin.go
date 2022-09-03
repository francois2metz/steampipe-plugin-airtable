package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-airtable",
		DefaultTransform: transform.FromGo(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		SchemaMode:   plugin.SchemaModeDynamic,
		TableMapFunc: PluginTables,
	}

	return p
}

func PluginTables(ctx context.Context, connection *plugin.Connection) (map[string]*plugin.Table, error) {
	airtableConfig := GetConfig(connection)

	client, err := rawConnect(ctx, connection)
	if err != nil {
		plugin.Logger(ctx).Error("airtable.init", "connection_error", err)
		return nil, err
	}

	result, err := client.GetBase(*airtableConfig.DatabaseID).Do()
	if err != nil {
		plugin.Logger(ctx).Error("airtable.init", err)
		return nil, err
	}
	tableMap := map[string]*plugin.Table{}
	for _, table := range result.Tables {
		tableMap[toTableName(table.Name)] = tableAirtableRecord(ctx, table)
	}

	tableMap["airtable_base"] = tableAirtableBase()
	tableMap["airtable_table"] = tableAirtableTable()

	return tableMap, nil
}
