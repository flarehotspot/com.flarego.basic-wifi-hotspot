package controllers

import (
	curr "github.com/flarehotspot/sdk/api/currencies"
	"github.com/flarehotspot/sdk/api/models/device"
	pymnt "github.com/flarehotspot/sdk/api/payments"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/web/contexts"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
	"log"
	"net/http"
)

type PortalCtrl struct {
	api plugin.IPluginApi
}

func (ctrl *PortalCtrl) GetInsertCoin(w http.ResponseWriter, r *http.Request) {
	device, ok := r.Context().Value(contexts.DeviceCtxKey).(device.IDeviceInstance)
	if ok && device != nil {
		log.Println("Insert coin device mac: ", device.MacAddress())

		item := &pymnt.PaymentRequestItem{
			Sku:         "some-sku",
			Name:        "Wifi Connection",
			Description: "Purchase for wifi connection",
			UnitAmount: &pymnt.UnitAmount{
				CurrencyCode: curr.CurrencyPhilippinePeso,
				Value:        11.0,
			},
		}

		params := &pymnt.PaymentRequestParams{
			Items:       []*pymnt.PaymentRequestItem{item},
			ReturnRoute: names.RoutePaymentReceived,
		}

		ctrl.api.HttpApi().Respond().RequestPayment(w, r, params)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
	return &PortalCtrl{api}
}
