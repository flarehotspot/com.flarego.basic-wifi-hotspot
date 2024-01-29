//go:build !mono

package main

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/web/navs"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func main() {}

func Init(api plugin.IPluginApi) {
	routes.PortalRoutes(api)
	routes.AdminRoutes(api)

	navs.SetPortalItems(api)
	navs.AdminWifiRates(api)
}
