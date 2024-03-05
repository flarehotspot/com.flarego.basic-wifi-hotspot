//go:build !mono

package main

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/web/navs"
	plugin "github.com/flarehotspot/sdk/api/plugin"
)

func main() {}

func Init(api plugin.PluginApi) {
	routes.PortalRoutes(api)
	routes.AdminRoutes(api)
    navs.SetAdminNavs(api)
	navs.SetPortalItems(api)
}
