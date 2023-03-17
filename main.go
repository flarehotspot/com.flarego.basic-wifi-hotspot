//go:build !dev

package main

import (
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/routes"
	"github.com/flarehotspot/wifi-hotspot/app/web/navs"
)

func main() {}

func Init(api plugin.IPluginApi) {
	routes.SetupRoutes(api)
	navs.SetPortalItems(api)
}
