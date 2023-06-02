package controllers

import (
	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/http/contexts"
	"github.com/flarehotspot/sdk/api/payments"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
	"log"
	"net/http"
)

type PortalCtrl struct {
	api plugin.IPluginApi
}

func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
	return &PortalCtrl{api}
}

func (ctrl *PortalCtrl) GetInsertCoin(w http.ResponseWriter, r *http.Request) {
	client, ok := r.Context().Value(contexts.ClientCtxKey).(connmgr.IClientDevice)
	if ok && client != nil {
		device := client.Device()
		log.Println("Insert coin device mac: ", device.MacAddress())

		item := &payments.PurchaseItem{
			Sku:         "some-sku",
			Name:        "Wifi Connection",
			Description: "Purchase for wifi connection",
			Price:       11.1,
		}

		params := &payments.PurchaseRequest{
			Items:       []*payments.PurchaseItem{item},
			VarPrice:    true,
			CallbackUrl: ctrl.api.HttpApi().Router().UrlForRoute(names.RoutePaymentReceived),
		}

		ctrl.api.PaymentsApi().Checkout(w, r, params)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ctrl *PortalCtrl) StartSession(w http.ResponseWriter, r *http.Request) {
	clntSym := r.Context().Value(contexts.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.IClientDevice)
	if !ok {
		http.Error(w, "Could not determine client device", http.StatusInternalServerError)
		return
	}

	_, err := clnt.HasValidSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := clnt.NextSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Session Type:", s.SessionType())
	log.Println("Time:", s.TimeSecs())
	log.Println("Data:", s.DataMbyte())
	log.Println("Session Expiration:", s.ExpiresAt())

	w.WriteHeader(200)
}
