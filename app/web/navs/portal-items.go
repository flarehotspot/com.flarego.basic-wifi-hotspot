package navs

import (
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/navigation/navgen"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

func SetPortalItems(api plugin.IPluginApi) {
	portalItem := navgen.NewPortalItem(api, "insert_coin", names.RouteInsertCoin)
	api.NavApi().NewPortalNav(portalItem)
}
