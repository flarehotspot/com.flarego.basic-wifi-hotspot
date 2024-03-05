package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

type PaymentSettings []struct {
	Price int `json:"price"`
	Data  int `json:"dataAlloc"`
	Time  int `json:"timeInMins"`
}

func Payments(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}

		//reading the request body using io.readall instead of r.postformvalue
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var userData PaymentSettings
		//unmarshalling the json request body into the userinput stru
		err = json.Unmarshal(body, &userData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		api.Http().VueResponse().Json(w, userData, 200)
		api.Config().Plugin().WriteJson(&userData)
	}
}
