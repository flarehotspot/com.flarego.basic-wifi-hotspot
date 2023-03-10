//go:build dev

package wifihotspot

import (
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/routes"
	"github.com/flarehotspot/wifi-hotspot/app/web/navs"
)

func Init(api plugin.IPluginApi) {
	routes.SetupRoutes(api)
	navs.SetPortalItems(api)
}
