package controllers

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
)

type WifiRatesCtrl struct {
	api plugin.IPluginApi
}

func NewWifiRatesCtrl(api plugin.IPluginApi) *WifiRatesCtrl {
	return &WifiRatesCtrl{api}
}

func (self *WifiRatesCtrl) Index(w http.ResponseWriter, r *http.Request) {
	self.api.HttpApi().Respond().AdminView(w, r, "rates/index.html", nil)
}
