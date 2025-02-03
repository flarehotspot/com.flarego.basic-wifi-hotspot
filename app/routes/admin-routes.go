package routes

import (
	sdkapi "sdk/api"
)

func AdminRoutes(api sdkapi.IPluginApi) {
	// adminR := api.Http().HttpRouter().AdminRouter()

	// adminR.Group("/payment-settings", func(subrouter sdkhttp.HttpRouterInstance) {
	// 	subrouter.
	// 		Get("/", controllers.GetPaymentSettings(api)).
	// 		Name("admin.payment-settings.get")

	// 	subrouter.
	// 		Post("/payment-settings/save", controllers.SavePaymentSettings(api)).
	// 		Name("admin.payment-settings.save")
	// })

	// api.Http().VueRouter().RegisterAdminRoutes(sdkhttp.VueAdminRoute{
	// 	RouteName: "admin.payment-settings",
	// 	RoutePath: "/payment-settings",
	// 	Component: "admin/payment-settings.vue",
	// })

	//    api.Http().HttpRouter().PluginRouter().Group("/payments")
}
