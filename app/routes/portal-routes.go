package routes

import (
	"com.flarego.basic-wifi-hotspot/app/controllers"
	sdkhttp "sdk/api/http"
	plugin "sdk/api/plugin"
)

func PortalRoutes(api plugin.PluginApi) {
	portalRouter := api.Http().HttpRouter().PluginRouter()
	portalRouter.Group("/purchase", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/wifi", controllers.PurchaseWifiSession(api)).Name("portal.purchase.wifi")
		subrouter.Get("/callback", controllers.PaymentRecevied(api)).Name("portal.purchase.callback")
	})

	portalRouter.Group("/sessions", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Post("/start", controllers.StartSession(api)).Name("portal.sessions.start")
		subrouter.Post("/stop", controllers.PauseSession(api)).Name("portal.sessions.stop")
	})
}
