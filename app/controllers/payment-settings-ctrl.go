package controllers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func GetPaymentSettings(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().VueResponse()

		settings, err := utils.GetPaymentConfig(api)
		if err != nil {
			res.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, settings, http.StatusOK)
	}
}

func SavePaymentSettings(api sdkplugin.PluginApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var settings utils.PaymentSettings
		err := json.NewDecoder(r.Body).Decode(&settings)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		sort.Slice(settings, func(i, j int) bool {
			return settings[i].Amount > settings[j].Amount
		})

		err = api.Config().Custom("default").Save(&settings)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := api.Http().VueResponse()

		res.SendFlashMsg(w, "success", "Settings saved successfully", http.StatusOK)
	})
}
