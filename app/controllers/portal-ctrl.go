package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/http/contexts"
	"github.com/flarehotspot/sdk/api/http/flash"
	"github.com/flarehotspot/sdk/api/payments"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/api/utils"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
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
	clnt := clntSym.(connmgr.IClientDevice)

	if clnt.IsConnected() {
		msg := "Client device is already connected."
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, utils.MsgTypeInfo, msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := clnt.Connect()
	if err != nil {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, flash.FlashTypeError, err.Error())
	} else {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, flash.FlashTypeInfo, "You are now connected to internet.")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (ctrl *PortalCtrl) StopSession(w http.ResponseWriter, r *http.Request) {
	clntSym := r.Context().Value(contexts.ClientCtxKey)
	clnt := clntSym.(connmgr.IClientDevice)

	if !clnt.IsConnected() {
		msg := "Client device is not connected."
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, utils.MsgTypeInfo, msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := clnt.Disconnect("Session paused.")
	if err != nil {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, flash.FlashTypeError, err.Error())
	} else {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, flash.FlashTypeError, "You are now disconnected from internet.")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
