package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/models"
	"github.com/flarehotspot/sdk/api/payments"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/contexts"
)

type PortalCtrl struct {
	api plugin.IPluginApi
}

func (ctrl *PortalCtrl) GetInsertCoin(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(contexts.DeviceCtxKey).(models.IDeviceInstance)
	log.Println("Insert coin device mac: ", device.MacAddress())

	item := &payments.PurchaseItem{
		Sku:         "some-sku",
		Name:        "Wifi Connection",
		Description: "Purchase for wifi connection",
	}

	params := &payments.PurchaseParams{
		Items:         []*payments.PurchaseItem{item},
		ReturnRoute:   "payment:received",
	}

	ctrl.api.HttpApi().Respond().RequestPayment(w, r, params)
}

func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
	return &PortalCtrl{api}
}
