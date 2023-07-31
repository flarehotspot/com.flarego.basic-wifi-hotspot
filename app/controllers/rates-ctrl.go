package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/api/config"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/errutil"
	"github.com/flarehotspot/sdk/utils/flash"
	"github.com/flarehotspot/sdk/utils/slices"
	"github.com/flarehotspot/sdk/utils/strings"
)

type WifiRatesCtrl struct {
	api plugin.IPluginApi
}

func NewWifiRatesCtrl(api plugin.IPluginApi) *WifiRatesCtrl {
	return &WifiRatesCtrl{api}
}

func (ctrl *WifiRatesCtrl) Index(w http.ResponseWriter, r *http.Request) {
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

func (ctrl *WifiRatesCtrl) Save(w http.ResponseWriter, r *http.Request) {
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

	newRates := []*config.WifiRate{}
	for _, newRate := range ratesData {
		prevRate := findRate(rates, newRate.Uuid)
		if prevRate != nil {
			rate := &config.WifiRate{
				Uuid:       prevRate.Uuid,
				Network:    network,
				Amount:     newRate.Amount,
				TimeMins:   newRate.TimeMins,
				DataMbytes: newRate.DataMbytes,
			}
			newRates = append(newRates, rate)
		} else {
			if createRate {
				rate := &config.WifiRate{
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

func (ctrl *WifiRatesCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := ctrl.api.HttpApi().MuxVars(r)["uuid"]
	rates, err := ctrl.api.ConfigApi().WifiRates().All()
	if err != nil {
		ctrl.Error(w, r, err)
		return
	}

	rates = slices.Filter(rates, func(r *config.WifiRate) bool {
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

func (ctrl *WifiRatesCtrl) Error(w http.ResponseWriter, r *http.Request, err error) {
	errRoute := errutil.NewErrRedirect(ctrl.indexUrl())
	errRoute.Redirect(w, r, err)
}

func (ctrl *WifiRatesCtrl) indexUrl() string {
	return ctrl.api.HttpApi().Router().UrlForRoute(names.RouteAdminRatesIndex)
}

func findRate(rates []*config.WifiRate, uuid string) *config.WifiRate {
	for _, r := range rates {
		if r.Uuid == uuid {
			return r
		}
	}
	return nil
}

func formToRates(formData []map[string]string) ([]*config.WifiRate, error) {
	rates := []*config.WifiRate{}
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

		rate := &config.WifiRate{
			Uuid:       uuid,
			Amount:     amount,
			TimeMins:   uint(mins),
			DataMbytes: uint(mbytes),
		}

		rates = append(rates, rate)
	}

	return rates, nil
}
