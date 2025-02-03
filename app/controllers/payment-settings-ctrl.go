package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	sdkplugin "sdk/api"

	"com.flarego.basic-wifi-hotspot/app/utils"
)

func GetPaymentSettings(api sdkplugin.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().HttpResponse()

		settings, err := utils.GetPaymentConfig(api)
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		w.Write([]byte(fmt.Sprintf("Must show payment settings: %+v\n", settings)))
	}
}

func SavePaymentSettings(api sdkplugin.IPluginApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().HttpResponse()
		var settings utils.PaymentSettings
		err := json.NewDecoder(r.Body).Decode(&settings)
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		sort.Slice(settings, func(i, j int) bool {
			return settings[i].Amount > settings[j].Amount
		})

		w.Write([]byte("must show payment settings"))

		// err = api.Config().Custom("default").Save(&settings)
		// if err != nil {
		// 	api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// res := api.Http().VueResponse()

		// res.SendFlashMsg(w, "success", "Settings saved successfully", http.StatusOK)
	})
}
