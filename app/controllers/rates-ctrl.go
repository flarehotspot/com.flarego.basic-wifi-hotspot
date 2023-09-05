package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/core/sdk/api/config"
	"github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/sdk/utils/flash"
	"github.com/flarehotspot/core/sdk/utils/slices"
	"github.com/flarehotspot/core/sdk/utils/strings"
)

type SessionRatesCtrl struct {
	api plugin.IPluginApi
}

func NewWifiRatesCtrl(api plugin.IPluginApi) *SessionRatesCtrl {
	return &SessionRatesCtrl{api}
}

func (ctrl *SessionRatesCtrl) Index(w http.ResponseWriter, r *http.Request) {
	rates, err := ctrl.api.ConfigApi().WifiRates().All()
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	network := "10.0.0.0/20"
	if len(rates) > 0 {
		network = rates[0].Network
	}

	data := map[string]any{
		"rates":   rates,
		"network": network,
	}
	ctrl.api.HttpApi().Respond().AdminView(w, r, "rates/index.html", data)
}

func (ctrl *SessionRatesCtrl) Save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	rates, err := ctrl.api.ConfigApi().WifiRates().All()
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	network := r.FormValue("network")
	createRate := r.FormValue("new_rate") == "on"
	formData := []map[string]string{}

	count := 0
	for key, vals := range r.PostForm {
		if key == "uuid" || key == "amount" || key == "time_mins" || key == "data_mbytes" {
			for ii, v := range vals {
				if count == 0 {
					data := map[string]string{key: v}
					formData = append(formData, data)
				} else {
					formData[ii][key] = v
				}
			}

			count += 1
		}
	}

	ratesData, err := formToRates(formData)
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	newRates := []*config.SessionRate{}
	for _, newRate := range ratesData {
		prevRate := findRate(rates, newRate.Uuid)
		if prevRate != nil {
			rate := &config.SessionRate{
				Uuid:       prevRate.Uuid,
				Network:    network,
				Amount:     newRate.Amount,
				TimeMins:   newRate.TimeMins,
				DataMbytes: newRate.DataMbytes,
			}
			newRates = append(newRates, rate)
		} else {
			if createRate {
				rate := &config.SessionRate{
					Uuid:       strings.Rand(8),
					Network:    network,
					Amount:     newRate.Amount,
					TimeMins:   newRate.TimeMins,
					DataMbytes: newRate.DataMbytes,
				}
				newRates = append(newRates, rate)
			}
		}
	}

	log.Println("Network: ", network)

	_, err = ctrl.api.ConfigApi().WifiRates().Write(newRates)
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	flash.SetFlashMsg(w, flash.Success, "Wifi rate saved successfully.")
	http.Redirect(w, r, ctrl.indexUrl(), http.StatusSeeOther)
}

func (ctrl *SessionRatesCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := ctrl.api.HttpApi().MuxVars(r)["uuid"]
	rates, err := ctrl.api.ConfigApi().WifiRates().All()
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	rates = slices.Filter(rates, func(r *config.SessionRate) bool {
		return r.Uuid != uuid
	})

	_, err = ctrl.api.ConfigApi().WifiRates().Write(rates)
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	flash.SetFlashMsg(w, flash.Info, "Wifi rate deleted successfully.")
	http.Redirect(w, r, ctrl.indexUrl(), http.StatusSeeOther)
}

func (ctrl *SessionRatesCtrl) Error(w http.ResponseWriter, r *http.Request, err error) {
	e := ctrl.api.HttpApi().Respond().NewErrRoute(names.RouteAdminRatesIndex)
	e.Redirect(w, r, err)
}

func (ctrl *SessionRatesCtrl) indexUrl() string {
	return ctrl.api.HttpApi().Router().UrlForRoute(names.RouteAdminRatesIndex)
}

func findRate(rates []*config.SessionRate, uuid string) *config.SessionRate {
	for _, r := range rates {
		if r.Uuid == uuid {
			return r
		}
	}
	return nil
}

func formToRates(formData []map[string]string) ([]*config.SessionRate, error) {
	rates := []*config.SessionRate{}
	for _, data := range formData {
		uuid := data["uuid"]
		amount, err := strconv.ParseFloat(data["amount"], 64)
		if err != nil {
			return nil, err
		}

		mins, err := strconv.Atoi(data["time_mins"])
		if err != nil {
			return nil, err
		}

		mbytes, err := strconv.Atoi(data["data_mbytes"])
		if err != nil {
			return nil, err
		}

		rate := &config.SessionRate{
			Uuid:       uuid,
			Amount:     amount,
			TimeMins:   uint(mins),
			DataMbytes: uint(mbytes),
		}

		rates = append(rates, rate)
	}

	return rates, nil
}
