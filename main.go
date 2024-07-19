//go:build !mono

package main

import (
	"com.flarego.basic-wifi-hotspot/app/routes"
	"com.flarego.basic-wifi-hotspot/app/web/navs"
	plugin "sdk/api/plugin"
)

func main() {}

func Init(api plugin.PluginApi) {
	if err := api.Migrate(); err != nil {
		api.Logger().Error(err.Error())
		return
	}

	routes.PortalRoutes(api)
	routes.AdminRoutes(api)
	navs.SetAdminNavs(api)
	navs.SetPortalItems(api)
}
