package routes

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	sdkhttp "sdk/api/http"
	plugin "sdk/api/plugin"
)

func AdminRoutes(api plugin.PluginApi) {
	adminR := api.Http().HttpRouter().AdminRouter()

	adminR.Group("/payment-settings", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.
			Get("/", controllers.GetPaymentSettings(api)).
			Name("admin.payment-settings.get")

		subrouter.
			Post("/payment-settings/save", controllers.SavePaymentSettings(api)).
			Name("admin.payment-settings.save")
	})

	api.Http().VueRouter().RegisterAdminRoutes(sdkhttp.VueAdminRoute{
		RouteName: "admin.payment-settings",
		RoutePath: "/payment-settings",
		Component: "admin/payment-settings.vue",
	})
}
