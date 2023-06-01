package navs

import (
	"github.com/flarehotspot/sdk/api/http/navigation/navgen"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

func SetPortalItems(api plugin.IPluginApi) {
	inscoin := navgen.NewPortalItem(api, "insert_coin", true, names.RouteInsertCoin)
	startSession := NewSessionBtnNav(api, nil)

	api.NavApi().NewPortalNav(inscoin)
	api.NavApi().NewPortalNav(startSession)
}
