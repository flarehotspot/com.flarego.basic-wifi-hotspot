package routes

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"

	// "github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
)

func AdminRoutes(api plugin.PluginApi) {

	api.Http().VueRouter().RegisterAdminRoutes(sdkhttp.VueAdminRoute{
		RouteName: "payment",
		RoutePath: "/payment",
		Component: "admin/PaymentDashboard.vue",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			var savedData controllers.UserInput

			// err := api.Config().Plugin().ReadJson(&savedData)
			api.Http().VueResponse().Json(w, savedData, 200)
			// if err != nil {
			// 	http.Error(w, "Unable to read JSON: "+err.Error(), http.StatusBadRequest)
			// 	return
			// }
		},
	})

	api.Http().VueRouter().AdminNavsFunc(func(r *http.Request) []sdkhttp.VueAdminNav {
		return []sdkhttp.VueAdminNav{
			{
				Category:  sdkhttp.NavCategoryPayments,
				Label:     "Payment",
				RouteName: "payment",
			},
		}
	})

	adminrouter := api.Http().HttpRouter().AdminRouter()
	adminrouter.Group("/save", func(subrouter sdkhttp.RouterInstance) {
		subrouter.Post("/settings", controllers.Payments(api)).Name(names.RoutePayments)
	})
	// rtr := api.HttpApi().HttpRouter().AdminRouter()
	// ratesCtrl := controllers.NewWifiRatesCtrl(api)

	// rtr.Get("/rates", ratesCtrl.Index).Name(names.RouteAdminRatesIndex)
	// rtr.Post("/rates/save", ratesCtrl.Save).Name(names.RouteAdminRatesSave)
	// rtr.Get("/rates/{uuid}/delete", ratesCtrl.Delete).Name(names.RouteAdminRatesDelete)
}
