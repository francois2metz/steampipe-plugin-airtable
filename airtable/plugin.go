package airtable

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	tableMap := map[string]*plugin.Table{}
	tableAirtable := []string{"Domaines", "Competences", "Acquis", "Epreuves", "Tubes"}
	for _, table := range tableAirtable {
		tableMap["airtable_"+strings.ToLower(table)] = tableAirtableTable(table)
	}
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-airtable",
		DefaultTransform: transform.FromGo(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: tableMap,
	}

	return p
}
