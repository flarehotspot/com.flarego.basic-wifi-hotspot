package navs

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/http/navigation"
	"github.com/flarehotspot/sdk/api/http/navigation/navgen"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

func SetPortalItems(api plugin.IPluginApi) {
	api.NavApi().PortalNavsFn(func(r *http.Request) []navigation.IPortalItem {
		inscoin := navgen.NewPortalItem(api, "insert_coin", true, names.RouteInsertCoin)
		startSession := NewSessionBtnNav(api, r)

		return []navigation.IPortalItem{
			inscoin,
			startSession,
		}
	})
}
