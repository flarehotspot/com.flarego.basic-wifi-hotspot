package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func SavePaymentSettings(api sdkplugin.PluginApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var settings utils.PaymentSettings
		err := json.NewDecoder(r.Body).Decode(&settings)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		err = api.Config().Plugin().WriteJson(&settings)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := api.Http().VueResponse()
		res.FlashMsg("success", "Settings saved successfully")
		res.Json(w, settings, http.StatusOK)
	})
}
