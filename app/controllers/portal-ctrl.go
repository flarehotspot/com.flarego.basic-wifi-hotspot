package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/payments"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/constants"
	"github.com/flarehotspot/sdk/utils/http/req"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

type PortalCtrl struct {
	api plugin.IPluginApi
}

func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
	return &PortalCtrl{api}
}

func (ctrl *PortalCtrl) GetInsertCoin(w http.ResponseWriter, r *http.Request) {
	clnt, ok := r.Context().Value(constants.ClientCtxKey).(connmgr.IClientDevice)
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
			CallbackUrl: ctrl.api.HttpApi().Router().UrlForRoute(names.RoutePaymentReceived),
		}

		ctrl.api.PaymentsApi().Checkout(w, r, params)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ctrl *PortalCtrl) StartSession(w http.ResponseWriter, r *http.Request) {
	clnt, err := req.ClientDevice(r)

	if err != nil {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeError, err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if ctrl.api.ClientMgr().IsConnected(clnt) {
		msg := "Client device is already connected."
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.TranslateInfo, msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = ctrl.api.ClientMgr().Connect(clnt)
	if err != nil {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeError, err.Error())
	} else {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeInfo, "You are now connected to internet.")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (ctrl *PortalCtrl) StopSession(w http.ResponseWriter, r *http.Request) {
	clnt, err := req.ClientDevice(r)

	if err != nil {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.TranslateInfo, err.Error())
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if !ctrl.api.ClientMgr().IsConnected(clnt) {
		msg := "Client device is not connected."
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.TranslateInfo, msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

  msg := "You are now disconnected from internet."
	err = ctrl.api.ClientMgr().Disconnect(clnt, errors.New(msg))
	if err != nil {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeError, err.Error())
	} else {
		ctrl.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeError, msg)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
