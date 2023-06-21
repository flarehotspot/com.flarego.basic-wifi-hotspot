package routes

import (
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/controllers"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

func AdminRoutes(api plugin.IPluginApi) {
	rtr := api.HttpApi().Router().AdminRouter()
	ratesCtrl := controllers.NewWifiRatesCtrl(api)

	rtr.Get("/rates", ratesCtrl.Index).Name(names.RouteAdminRatesIndex)
}
