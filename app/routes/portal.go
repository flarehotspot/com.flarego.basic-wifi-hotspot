package routes

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/router"
	"github.com/flarehotspot/wifi-hotspot/app/controllers"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

func SetupRoutes(api plugin.IPluginApi) {
	rtr := api.HttpApi().Router()
	portalCtrl := controllers.NewPortalCtrl(api)

	rtr.PluginRouter().Group("/portal", func(subrouter router.IRouter) {
		subrouter.Get("/insert-coin", portalCtrl.GetInsertCoin).Name(names.RouteInsertCoin)
		subrouter.Get("/payment/received", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Payment received!"))
		}).Name(names.RoutePaymentReceived)
	})
}
