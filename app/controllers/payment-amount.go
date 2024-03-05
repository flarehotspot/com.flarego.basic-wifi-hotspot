package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

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

		for i := 0; i < 3; i++ {
			price, err := strconv.Atoi(strconv.Itoa(userData[i].Price))
			if err != nil {
				http.Error(w, "Price parsing error: "+err.Error(), http.StatusBadRequest)
				return
			}
			data, err := strconv.Atoi(strconv.Itoa(userData[i].Data))
			if err != nil {
				http.Error(w, "Data parsing error:"+err.Error(), http.StatusBadRequest)
				return

			}

			userData[i].Price = price
			userData[i].Data = data
		}

		api.Http().VueResponse().Json(w, userData, 200)
		api.Config().Plugin().WriteJson(&userData)
	}
}
