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

	tableMap := map[string]*plugin.Table{}

	for _, table := range airtableConfig.Tables {
		tableMap[toTableName(table)] = tableAirtableRecord(table)
	}

	return tableMap, nil
}
