package airtable

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/schema"
)

type airtableConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &airtableConfig{}
}

func GetConfig(connection *plugin.Connection) airtableConfig {
	if connection == nil || connection.GetConfig() == nil {
		return airtableConfig{}
	}
	config, _ := connection.GetConfig().(airtableConfig)
	return config
}
