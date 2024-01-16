package navs

import (
	"net/http"

	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/plugin"
	// "github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
)

func SetPortalItems(api plugin.IPluginApi) {

	api.HttpApi().VueRouter().PortalRoutes(func(r *http.Request) []router.VuePortalRoute {
		return []router.VuePortalRoute{
			{RouteName: "sample", RoutePath: "/sample", ComponentPath: "components/portal/Sample.vue"},
		}
	})

	api.HttpApi().VueRouter().PortalItems(func(r *http.Request) []router.VuePortalItem {
		navs := []router.VuePortalItem{}
		navs = append(navs, router.VuePortalItem{
			TranslateLabel: "sample",
			RouteName:      "sample",
		})
		return navs
	})

	// api.NavApi().PortalNavsFn(func(r *http.Request) []navigation.Portal {
	// 	inscoin := navigation.NewPortalItem("", "Insert Coin", "", api.HttpApi().Router().UrlForRoute(names.RouteInsertCoin))
	// 	navs := []navigation.IPortalItem{inscoin}

	// 	startSession := NewSessionBtnNav(api, r)
	// 	if clnt, err := startSession.client(); err == nil {
	// 		if clnt.HasSession(r.Context()) {
	// 			navs = append(navs, startSession)
	// 		}
	// 	}

	// 	return navs
	// })
}
