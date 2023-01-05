package airtable

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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

func PluginTables(ctx context.Context, d *plugin.TableMapData) (map[string]*plugin.Table, error) {
	tableMap := map[string]*plugin.Table{}
	client, err := rawConnect(ctx, d.Connection, d.ConectionCache)
	if err != nil {
		plugin.Logger(ctx).Error("airtable.init", "connection_error", err)
		return nil, err
	}

	bases := client.GetBases()
	offset := ""
	for {
		resultBase, err := bases.WithOffset(offset).Do()
		if err != nil {
			plugin.Logger(ctx).Error("airtable.init_bases", err)
			return nil, err
		}
		for _, base := range resultBase.Bases {
			result, err := client.GetBaseSchema(base.ID).Do()
			if err != nil {
				plugin.Logger(ctx).Error("airtable.init_getbaseschema", err)
				return nil, err
			}
			for _, table := range result.Tables {
				tableMap[toTableName(base.ID, table.Name)] = tableAirtableRecord(ctx, base, table)
			}
		}
		offset = resultBase.Offset
		if offset == "" {
			break
		}
	}

	tableMap["airtable_base"] = tableAirtableBase()
	tableMap["airtable_table"] = tableAirtableTable()

	return tableMap, nil
}
