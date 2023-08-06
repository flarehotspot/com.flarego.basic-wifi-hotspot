package routes

import (
	"github.com/flarehotspot/sdk/api/http/router"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
)

func PortalRoutes(api plugin.IPluginApi) {
	rtr := api.HttpApi().Router()
	portalCtrl := controllers.NewPortalCtrl(api)
	paymentsCtrl := controllers.NewPaymentCtrl(api)
	deviceMw := api.HttpApi().Middlewares().Device()

	rtr.PluginRouter().Group("/portal", func(subrouter router.IRouter) {
		subrouter.Use(deviceMw)
		subrouter.Get("/insert-coin", portalCtrl.GetInsertCoin).Name(names.RouteInsertCoin)
	})

	rtr.PluginRouter().Group("/payments", func(subrouter router.IRouter) {
		subrouter.Use(deviceMw)
		subrouter.Get("/received", paymentsCtrl.PaymentRecevied).Name(names.RoutePaymentReceived)
	})

	rtr.PluginRouter().Group("/session", func(subrouter router.IRouter) {
		subrouter.Use(deviceMw)
		subrouter.Get("/start", portalCtrl.StartSession).Name(names.RouteStartSession)
		subrouter.Get("/stop", portalCtrl.StopSession).Name(names.RouteStopSession)
	})
}
