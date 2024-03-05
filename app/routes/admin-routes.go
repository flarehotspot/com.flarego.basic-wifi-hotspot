package routes

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	plugin "github.com/flarehotspot/sdk/api/plugin"
)

func AdminRoutes(api plugin.PluginApi) {
	api.Http().HttpRouter().AdminRouter().
		Post("/payment-settings/save", controllers.SavePaymentSettings(api)).
		Name("admin.payment-settings.save")

	api.Http().VueRouter().RegisterAdminRoutes(sdkhttp.VueAdminRoute{
		RouteName: "admin.payment-settings",
		RoutePath: "/payment-settings",
		Component: "admin/payment-settings.vue",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			res := api.Http().VueResponse()

			var settings utils.PaymentSettings
			err := api.Config().Plugin().ReadJson(&settings)
			if err != nil {
				res.Json(w, utils.DefaultPaymentSettings, http.StatusOK)
				return
			}

			res.Json(w, settings, http.StatusOK)
		},
	})
}
