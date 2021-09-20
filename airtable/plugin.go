package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-airtable",
		DefaultTransform: transform.FromGo(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMapFunc: PluginTables,
	}

	return p
}

func PluginTables(p *plugin.Plugin) (map[string]*plugin.Table, error) {
	airtableConfig := GetConfig(p.Connection)

	tableMap := map[string]*plugin.Table{}

	for _, table := range airtableConfig.Tables {
		tableMap["airtable_"+ toTableName(table)] = tableAirtableTable(table)
	}

	return tableMap, nil
}
