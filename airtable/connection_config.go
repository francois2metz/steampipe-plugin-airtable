package airtable

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type airtableConfig struct {
	Token      *string `cty:"token"`
	DatabaseID *string `cty:"database_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"database_id": {
		Type: schema.TypeString,
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
