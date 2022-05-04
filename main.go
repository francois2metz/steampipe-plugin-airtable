package main

import (
	"github.com/francois2metz/steampipe-plugin-airtable/airtable"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: airtable.Plugin})
}
