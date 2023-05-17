package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
)

type PaymentCtrl struct {
	api plugin.IPluginApi
}

func (self *PaymentCtrl) PaymentRecevied(w http.ResponseWriter, r *http.Request) {
	info, err := self.api.PaymentsApi().ParsePaymentInfo(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = info.Purchase.Confirm(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Payment Received: \n%+v", info.Purchase)
	w.WriteHeader(http.StatusOK)

}

func NewPaymentCtrl(api plugin.IPluginApi) *PaymentCtrl {
	return &PaymentCtrl{api}
}
