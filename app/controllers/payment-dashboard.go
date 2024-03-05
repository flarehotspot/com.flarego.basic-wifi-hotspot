package controllers

import (
	"encoding/json"
	"net/http"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func PaymentDashboard(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method. ", http.StatusBadRequest)
			return
		}

		var savedData UserInput

		err := json.NewDecoder(r.Body).Decode(&savedData)
		if err != nil {
			http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		if savedData.Out == nil || len(savedData.Out) != 3 {
			http.Error(w, "Invalid input format.", http.StatusBadRequest)
			return
		}

		for i := 0; i < 3; i++ {
			if savedData.Out[i].Price <= 0 || savedData.Out[i].Data <= 0 || savedData.Out[i].Time <= "" {
				http.Error(w, "Invalid input. ", http.StatusBadRequest)
				return
			}
		}

		err = api.Config().Plugin().ReadJson(savedData)
		if err != nil {
			http.Error(w, "Unable to read JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		res := api.Http().VueResponse()
		res.Json(w, savedData, 200)
	}
}
