package controllers

import (
	"github.com/flarehotspot/sdk/api/currencies"
	"github.com/flarehotspot/sdk/api/models/device"
	"github.com/flarehotspot/sdk/api/payments"
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

		item := &payments.PaymentRequestItem{
			Sku:         "some-sku",
			Name:        "Wifi Connection",
			Description: "Purchase for wifi connection",
			UnitAmount: &payments.UnitAmount{
				CurrencyCode:   currencies.CurrencyPhilippinePeso,
				VariableAmount: true,
			},
		}

		params := &payments.PaymentRequestParams{
			Items:       []*payments.PaymentRequestItem{item},
			CallbackUrl: ctrl.api.HttpApi().Router().UrlForRoute(names.RoutePaymentReceived),
		}

		ctrl.api.PaymentsApi().RequestPayment(w, r, params)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
	return &PortalCtrl{api}
}
