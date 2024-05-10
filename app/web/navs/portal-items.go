package navs

import (
	"context"
	"net/http"

	// "github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkpayments "github.com/flarehotspot/sdk/api/payments"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func SetPortalItems(api sdkplugin.PluginApi) {
	portalR := api.Http().HttpRouter().PluginRouter()

	portalR.Group("/portal", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/insert-coin", func(w http.ResponseWriter, r *http.Request) {
			p := sdkpayments.PurchaseRequest{
				Sku:                  "wifi-connection",
				Name:                 "WiFi Connection",
				Description:          "Basic Wifi Hotspot",
				AnyPrice:             true,
				CallbackVueRouteName: "portal.purchase-callback",
			}
			api.Payments().Checkout(w, r, p)
		})
	})

	vrouter := api.Http().VueRouter()
	vrouter.RegisterPortalRoutes([]sdkhttp.VuePortalRoute{
		{
			RouteName: "portal.insert-coin",
			RoutePath: "/insert-coin",
			Component: "portal/InsertCoin.vue",
		},
		{
			RouteName: "portal.purchase-callback",
			RoutePath: "/purchase-callback",
			Component: "portal/PurchaseCallback.vue",
		},
		{
			RouteName: "portal.start-session",
			RoutePath: "/start-session",
			Component: "portal/StartSession.vue",
		},
		{
			RouteName: "portal.pause-session",
			RoutePath: "/pause-session",
			Component: "portal/PauseSession.vue",
		},
	}...)

	vrouter.PortalItemsFunc(func(clnt sdkconnmgr.ClientDevice) []sdkhttp.VuePortalItem {
		navs := []sdkhttp.VuePortalItem{}

		navs = append(navs, sdkhttp.VuePortalItem{
			IconPath:  "images/wifi-logo.png",
			Label:     "Insert Coin",
			RouteName: "portal.insert-coin",
		})

		if api.SessionsMgr().IsConnected(clnt) {
			navs = append(navs, sdkhttp.VuePortalItem{
				IconPath:  "images/wifi-logo.png",
				Label:     "Pause Session",
				RouteName: "portal.pause-session",
			})
		} else if _, err := api.SessionsMgr().GetSession(context.Background(), clnt); err == nil {
			navs = append(navs, sdkhttp.VuePortalItem{
				IconPath:  "images/wifi-logo.png",
				Label:     "Start Session",
				RouteName: "portal.start-session",
			})
		}

		// reload portal when session connected or disconnected
		go utils.PortalReload(api, clnt)

		return navs
	})

}
