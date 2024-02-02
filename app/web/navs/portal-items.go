package navs

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	sdkpayments "github.com/flarehotspot/core/sdk/api/payments"
	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func SetPortalItems(api sdkplugin.PluginApi) {

	vrouter := api.Http().VueRouter()

	vrouter.RegisterPortalRoutes([]sdkhttp.VuePortalRoute{
		{
			RouteName: "portal.insert-coin",
			RoutePath: "/insert-coin",
			Component: "portal/InsertCoin.vue",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				p := sdkpayments.PurchaseRequest{
					Sku:                  "wifi-connection",
					Name:                 "WiFi Connection",
					Description:          "Basic Wifi Hotspot",
					AnyPrice:             true,
					CallbackVueRouteName: "portal.purchase-callback",
				}
				api.Payments().Checkout(w, r, p)
			},
		},
		{
			RouteName: "portal.purchase-callback",
			RoutePath: "/purchase-callback",
			Component: "portal/PurchaseCallback.vue",
		},
	}...)

	vrouter.PortalItemsFunc(func(r *http.Request) []sdkhttp.VuePortalItem {
		navs := []sdkhttp.VuePortalItem{}
		navs = append(navs, sdkhttp.VuePortalItem{
			IconPath:  "images/wifi-logo.png",
			Label:     "insert_coin",
			RouteName: "portal.insert-coin",
		})
		return navs
	})

	// api.HttpApi().VueRouter().PortalRoutes(func(r *http.Request) []router.VuePortalRoute {
	// 	return []router.VuePortalRoute{
	// 		{RouteName: "sample", RoutePath: "/sample", ComponentPath: "components/portal/Sample.vue"},
	// 	}
	// })

	// api.HttpApi().VueRouter().PortalItems(func(r *http.Request) []router.VuePortalItem {
	// 	navs := []router.VuePortalItem{}
	// 	navs = append(navs, router.VuePortalItem{
	// 		TranslateLabel: "sample",
	// 		RouteName:      "sample",
	// 	})
	// 	return navs
	// })

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
