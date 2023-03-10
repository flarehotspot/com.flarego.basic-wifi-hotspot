package navs

import (
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/navigation"
	"github.com/flarehotspot/wifi-hotspot/app/routes"
)

func SetPortalItems(api plugin.IPluginApi) {
	portalItem := navigation.PortalItem{
		Label:     "insert_coin",
		RouteName: routes.RouteInsertCoin,
	}
	api.NavApi().NewPortalNav(&portalItem)
}
