package routes

import (
	sdkapi "sdk/api"

	"com.flarego.basic-wifi-hotspot/app/controllers"
)

func PortalRoutes(api sdkapi.IPluginApi) {
	portalRouter := api.Http().HttpRouter().PluginRouter()
	portalRouter.Group("/purchase", func(subrouter sdkapi.IHttpRouterInstance) {
		subrouter.Get("/wifi", controllers.PurchaseWifiSession(api)).Name("purchase.wifi")
		subrouter.Get("/callback", controllers.PaymentRecevied(api)).Name("purchase.wifi.callback")
	})

	portalRouter.Group("/sessions", func(subrouter sdkapi.IHttpRouterInstance) {
		subrouter.Get("/start", controllers.StartSession(api)).Name("portal.sessions.start")
		subrouter.Get("/stop", controllers.PauseSession(api)).Name("portal.sessions.stop")
	})
}
