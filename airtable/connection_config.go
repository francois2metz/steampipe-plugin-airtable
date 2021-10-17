package airtable

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type airtableConfig struct {
	Token      *string  `cty:"token"`
	DatabaseID *string  `cty:"database_id"`
	Tables     []string `cty:"tables"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"database_id": {
		Type: schema.TypeString,
	},
	"tables": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &airtableConfig{}
}

func GetConfig(connection *plugin.Connection) airtableConfig {
	if connection == nil || connection.Config == nil {
		return airtableConfig{}
	}
	config, _ := connection.Config.(airtableConfig)
	return config
}
