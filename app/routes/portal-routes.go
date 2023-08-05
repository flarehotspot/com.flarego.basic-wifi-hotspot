package routes

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/v1.0.0/api"
	"github.com/flarehotspot/sdk/v1.0.0/api/http/router"
)

func PortalRoutes(API api.IPluginApi) {
	rtr := API.HttpApi().Router()
	portalCtrl := controllers.NewPortalCtrl(API)
	paymentsCtrl := controllers.NewPaymentCtrl(API)
	deviceMw := API.HttpApi().Middlewares().Device()

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
