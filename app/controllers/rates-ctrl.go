package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flarehotspot/sdk/api/config"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/constants"
	"github.com/flarehotspot/sdk/utils/slices"
	"github.com/flarehotspot/sdk/utils/strings"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

type WifiRatesCtrl struct {
	api plugin.IPluginApi
}

func NewWifiRatesCtrl(api plugin.IPluginApi) *WifiRatesCtrl {
	return &WifiRatesCtrl{api}
}

func (self *WifiRatesCtrl) Index(w http.ResponseWriter, r *http.Request) {
	rates, err := self.api.ConfigApi().WifiRates().All()
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
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
	self.api.HttpApi().Respond().AdminView(w, r, "rates/index.html", data)
}

func (self *WifiRatesCtrl) Save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	rates, err := self.api.ConfigApi().WifiRates().All()
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
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
		self.api.HttpApi().Respond().Error(w, r, err)
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

	_, err = self.api.ConfigApi().WifiRates().Write(newRates)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	self.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeSuccess, "Wifi rate saved successfully.")
	http.Redirect(w, r, self.api.HttpApi().Router().UrlForRoute(names.RouteAdminRatesIndex), http.StatusSeeOther)
}

func (self *WifiRatesCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	uuid := self.api.HttpApi().MuxVars(r)["uuid"]
	rates, err := self.api.ConfigApi().WifiRates().All()
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	rates = slices.Filter(rates, func(r *config.WifiRate) bool {
		return r.Uuid != uuid
	})

	_, err = self.api.ConfigApi().WifiRates().Write(rates)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	self.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeInfo, "Wifi rate deleted successfully.")
	ratesUrl := self.api.HttpApi().Router().UrlForRoute(names.RouteAdminRatesIndex)
	http.Redirect(w, r, ratesUrl, http.StatusSeeOther)
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
