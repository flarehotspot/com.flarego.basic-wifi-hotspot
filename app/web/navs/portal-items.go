package navs

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/v1.0.0/api"
	"github.com/flarehotspot/sdk/v1.0.0/api/http/navigation"
)

func SetPortalItems(API api.IPluginApi) {
	API.NavApi().PortalNavsFn(func(r *http.Request) []navigation.IPortalItem {
		inscoin := navigation.NewPortalItem("", "Insert Coin", "", API.HttpApi().Router().UrlForRoute(names.RouteInsertCoin))
		navs := []navigation.IPortalItem{inscoin}

		startSession := NewSessionBtnNav(API, r)
		if clnt, err := startSession.client(); err == nil {
			if clnt.HasSession(r.Context()) {
				navs = append(navs, startSession)
			}
		}

		return navs
	})
}
