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
		navs := []navigation.IPortalItem{inscoin}

		startSession := NewSessionBtnNav(api, r)
		if clnt, err := startSession.client(); err == nil {
			if clnt.HasSession() {
				navs = append(navs, startSession)
			}
		}

		return navs
	})
}
