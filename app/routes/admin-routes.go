package routes

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	plugin "github.com/flarehotspot/sdk/api/plugin"
)

func AdminRoutes(api plugin.PluginApi) {
	adminR := api.Http().HttpRouter().AdminRouter()

	adminR.Group("/payment-settings", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			res := api.Http().VueResponse()

			var settings utils.PaymentSettings
			err := api.Config().Plugin("default").Get(&settings)
			if err != nil {
				res.Json(w, utils.DefaultPaymentSettings, http.StatusOK)
				return
			}

			res.Json(w, settings, http.StatusOK)
		}).Name("admin.payment-settings.get")

		subrouter.
			Post("/payment-settings/save", controllers.SavePaymentSettings(api)).
			Name("admin.payment-settings.save")
	})

	api.Http().VueRouter().RegisterAdminRoutes(sdkhttp.VueAdminRoute{
		RouteName: "admin.payment-settings",
		RoutePath: "/payment-settings",
		Component: "admin/payment-settings.vue",
		// HandlerFunc: ,
	})
}
