package routes

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/router"
	"github.com/flarehotspot/wifi-hotspot/app/controllers"
)

const (
	RouteInsertCoin router.PluginRouteName = "insert-coin"
)

func SetupRoutes(api plugin.IPluginApi) {
	rtr := api.HttpApi().Router()
	portalCtrl := controllers.NewPortalCtrl(api)

	rtr.PluginRouter().Route("/portal", func(r router.IRouter) {
		r.Get("/insert-coin", portalCtrl.GetInsertCoin).Name(RouteInsertCoin)
		r.Get("/payment/received", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Payment received!"))
		}).Name("payment:received")
	})
}
