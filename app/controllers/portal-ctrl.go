package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/payments"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/contexts"
	"github.com/flarehotspot/sdk/utils/flash"
)

type PortalCtrl struct {
	api plugin.IPluginApi
}

func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
	return &PortalCtrl{api}
}

func (self *PortalCtrl) GetInsertCoin(w http.ResponseWriter, r *http.Request) {
	clnt, ok := r.Context().Value(contexts.ClientCtxKey).(connmgr.IClientDevice)
	if ok && clnt != nil {
		log.Println("Insert coin device mac: ", clnt.MacAddr())

		item := &payments.PurchaseItem{
			Sku:         "wifi-connection",
			Name:        "Wifi Connection",
			Description: "Purchase for wifi connection",
			Price:       11.1,
		}

		params := &payments.PurchaseRequest{
			Items:       []*payments.PurchaseItem{item},
			VarPrice:    true,
			CallbackUrl: self.api.HttpApi().Router().UrlForRoute(names.RoutePaymentReceived),
		}

		self.api.PaymentsApi().Checkout(w, r, params)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (self *PortalCtrl) StartSession(w http.ResponseWriter, r *http.Request) {
	clnt, err := self.api.ClientReg().CurrentClient(r)

	if err != nil {
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Error, err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if self.api.ClientMgr().IsConnected(clnt) {
		msg := "Client device is already connected."
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Warning, msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = self.api.ClientMgr().Connect(clnt)
	if err != nil {
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Error, err.Error())
	} else {
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Success, "You are now connected to internet.")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (self *PortalCtrl) StopSession(w http.ResponseWriter, r *http.Request) {
	clnt, err := self.api.ClientReg().CurrentClient(r)

	if err != nil {
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Info, err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if !self.api.ClientMgr().IsConnected(clnt) {
		msg := "Client device is not connected."
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Info, msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	msg := "You are now disconnected from internet."
	err = self.api.ClientMgr().Disconnect(clnt, errors.New(msg))
	if err != nil {
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Error, err.Error())
	} else {
		self.api.HttpApi().Respond().SetFlashMsg(w, flash.Error, msg)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
