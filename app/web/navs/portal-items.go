package navs

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkpayments "github.com/flarehotspot/sdk/api/payments"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func TextBtn(api sdkplugin.PluginApi, r *http.Request) string {
	clnt, err := api.Http().GetClientDevice(r)
	if err != nil {
		return err.Error()
	}

	if api.SessionsMgr().IsConnected(clnt) {
		return "Pause Session"
	}

	if api.SessionsMgr().HasSession(r.Context(), clnt.Id()) {
		return "Start Session"
	}

	return ""
}

func SessionRoute(api sdkplugin.PluginApi, r *http.Request) string {
	clnt, err := api.Http().GetClientDevice(r)
	if err != nil {
		return err.Error()
	}

	if api.SessionsMgr().HasSession(r.Context(), clnt.Id()) {
		if api.SessionsMgr().IsConnected(clnt) {
			return "portal.pause-session"
		} else {
			return "portal.start-session"
		}
	} else {
		return "/"
	}
}

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
		{
			RouteName:   "portal.pause-session",
			RoutePath:   "/pause-session",
			HandlerFunc: controllers.PauseSession(api),
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
			Label:     TextBtn(api, r),
			RouteName: SessionRoute(api, r),
		})

		if TextBtn(api, r) == "" {
			navs = append(navs[:1], navs[2:]...)
		}

		return navs
	})

}
