package controllers

// import (
// 	"errors"
// 	"log"
// 	"net/http"

// 	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
// 	"github.com/flarehotspot/core/sdk/api/connmgr"
// 	"github.com/flarehotspot/core/sdk/api/http/response"
// 	"github.com/flarehotspot/core/sdk/api/payments"
// 	"github.com/flarehotspot/core/sdk/api/plugin"
// 	"github.com/flarehotspot/core/sdk/utils/contexts"
// 	"github.com/flarehotspot/core/sdk/utils/flash"
// )

// type PortalCtrl struct {
// 	api      plugin.IPluginApi
// 	errRoute response.IErrorRedirect
// }

// func NewPortalCtrl(api plugin.IPluginApi) *PortalCtrl {
// 	errRoute := api.HttpApi().Respond().NewErrUrl("/")
// 	return &PortalCtrl{api, errRoute}
// }

// func (ctrl *PortalCtrl) GetInsertCoin(w http.ResponseWriter, r *http.Request) {
// 	clnt, ok := r.Context().Value(contexts.ClientCtxKey).(connmgr.IClientDevice)
// 	if ok && clnt != nil {
// 		log.Println("Insert coin device mac: ", clnt.MacAddr())

// 		item := &payments.PurchaseItem{
// 			Sku:         "wifi-connection",
// 			Name:        "Wifi Connection",
// 			Description: "Purchase for wifi connection",
// 			Price:       11.1,
// 		}

// 		params := &payments.PurchaseRequest{
// 			Items:       []*payments.PurchaseItem{item},
// 			VarPrice:    true,
// 			CallbackUrl: ctrl.api.HttpApi().HttpRouter().UrlForRoute(names.RoutePaymentReceived),
// 		}

// 		ctrl.api.PaymentsApi().Checkout(w, r, params)
// 	} else {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }

// func (ctrl *PortalCtrl) StartSession(w http.ResponseWriter, r *http.Request) {
// 	clnt, err := ctrl.api.ClientReg().CurrentClient(r)

// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	if ctrl.api.ClientMgr().IsConnected(clnt) {
// 		msg := "Client device is already connected."
// 		ctrl.errRoute.Redirect(w, r, errors.New(msg))
// 		return
// 	}

// 	err = ctrl.api.ClientMgr().Connect(clnt)
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 	} else {
// 		flash.SetFlashMsg(w, flash.Success, "You are now connected to internet.")
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

// func (ctrl *PortalCtrl) StopSession(w http.ResponseWriter, r *http.Request) {
// 	clnt, err := ctrl.api.ClientReg().CurrentClient(r)

// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	if !ctrl.api.ClientMgr().IsConnected(clnt) {
// 		msg := "Client device is not connected."
// 		flash.SetFlashMsg(w, flash.Info, msg)
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	msg := "You are now disconnected from internet."
// 	err = ctrl.api.ClientMgr().Disconnect(clnt, errors.New(msg))
// 	if err != nil {
// 		flash.SetFlashMsg(w, flash.Error, err.Error())
// 	} else {
// 		flash.SetFlashMsg(w, flash.Error, msg)
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }
