//go:build !mono

package main

import (
	plugin "sdk/api"

	"com.flarego.basic-wifi-hotspot/app/routes"
	"com.flarego.basic-wifi-hotspot/app/web/navs"
)

func main() {}

func Init(api plugin.IPluginApi) {
	routes.PortalRoutes(api)
	routes.AdminRoutes(api)
	navs.SetAdminNavs(api)
	navs.SetPortalItems(api)
}
