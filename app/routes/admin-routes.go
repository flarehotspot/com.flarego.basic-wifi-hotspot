package routes

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/v1.0.0/api/plugin"
)

func AdminRoutes(api plugin.IPluginApi) {
	rtr := api.HttpApi().Router().AdminRouter()
	ratesCtrl := controllers.NewWifiRatesCtrl(api)

	rtr.Get("/rates", ratesCtrl.Index).Name(names.RouteAdminRatesIndex)
	rtr.Post("/rates/save", ratesCtrl.Save).Name(names.RouteAdminRatesSave)
	rtr.Get("/rates/{uuid}/delete", ratesCtrl.Delete).Name(names.RouteAdminRatesDelete)
}
