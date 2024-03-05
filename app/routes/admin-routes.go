package routes

import (
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	plugin "github.com/flarehotspot/sdk/api/plugin"

	// "github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/controllers"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
)

func AdminRoutes(api plugin.PluginApi) {
	adminrouter := api.Http().HttpRouter().AdminRouter()
	adminrouter.Group("/save", func(subrouter sdkhttp.HttpRouterInstance) {
		subrouter.Post("/settings", controllers.Payments(api)).Name(names.RouteSaveSettings)
	})
}
