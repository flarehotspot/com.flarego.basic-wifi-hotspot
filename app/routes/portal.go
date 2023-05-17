package routes

import (
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/router"
	"github.com/flarehotspot/wifi-hotspot/app/controllers"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

func SetupRoutes(api plugin.IPluginApi) {
	rtr := api.HttpApi().Router()
	portalCtrl := controllers.NewPortalCtrl(api)
	deviceMw := api.HttpApi().Middlewares().Device()

	rtr.PluginRouter().Group("/portal", func(subrouter router.IRouter) {
		subrouter.Use(deviceMw)
		subrouter.Get("/insert-coin", portalCtrl.GetInsertCoin).Name(names.RouteInsertCoin)
	})

	rtr.PluginRouter().Group("/payments", func(subrouter router.IRouter) {
    subrouter.Use(deviceMw)
		subrouter.Get("/received", func(w http.ResponseWriter, r *http.Request) {
			paymt, err := api.PaymentsApi().ParsePaymentInfo(r)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
        return
			}
			log.Printf("Payment Received: \n%+v", paymt.Purchase)
			w.WriteHeader(http.StatusOK)
		}).Name(names.RoutePaymentReceived)
	})
}
