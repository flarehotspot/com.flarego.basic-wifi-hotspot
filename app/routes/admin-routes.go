package routes

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/v1/api"
)

func AdminRoutes(API api.IPluginApi) {
	rtr := API.HttpApi().Router().AdminRouter()
	ratesCtrl := controllers.NewWifiRatesCtrl(API)

	rtr.Get("/rates", ratesCtrl.Index).Name(names.RouteAdminRatesIndex)
	rtr.Post("/rates/save", ratesCtrl.Save).Name(names.RouteAdminRatesSave)
	rtr.Get("/rates/{uuid}/delete", ratesCtrl.Delete).Name(names.RouteAdminRatesDelete)
}
