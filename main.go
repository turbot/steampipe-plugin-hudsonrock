package main

import (
	"github.com/turbot/steampipe-plugin-hudsonrock/hudsonrock"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hudsonrock.Plugin})
}
