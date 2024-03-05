package controllers

import (
	"net/http"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func PaymentDashboard(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var savedData UserInput

		res := api.Http().VueResponse()
		res.Json(w, savedData, 200)
	}
}
