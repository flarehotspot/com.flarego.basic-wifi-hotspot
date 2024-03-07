package navs

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkpayments "github.com/flarehotspot/sdk/api/payments"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
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
			RouteName:   "portal.purchase-callback",
			RoutePath:   "/purchase-callback",
			Component:   "portal/PurchaseCallback.vue",
			HandlerFunc: controllers.PaymentRecevied(api),
		},
		{
			RouteName:   "portal.start-session",
			RoutePath:   "/start-session",
			HandlerFunc: controllers.StartSession(api),
		},
	}...)

	vrouter.PortalItemsFunc(func(r *http.Request) []sdkhttp.VuePortalItem {
		navs := []sdkhttp.VuePortalItem{}
		navs = append(navs, sdkhttp.VuePortalItem{
			IconPath:  "images/wifi-logo.png",
			Label:     "Insert Coin",
			RouteName: "portal.insert-coin",
		})

		navs = append(navs, sdkhttp.VuePortalItem{
			IconPath:  "images/wifi-logo.png",
			Label:     "Start Session",
			RouteName: "portal.start-session",
		})

		return navs
	})

}
