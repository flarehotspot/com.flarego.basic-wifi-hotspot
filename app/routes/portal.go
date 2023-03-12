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

	rtr.PluginRouter().Route("/portal", func(r router.IRouter) {
		r.Get("/insert-coin", portalCtrl.GetInsertCoin).Name(names.RouteInsertCoin)
		r.Get("/payment/received", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Payment received!"))
		}).Name(names.RoutePaymentReceived)
	})
}
