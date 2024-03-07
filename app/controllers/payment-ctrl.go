package controllers

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
)

func PaymentRecevied(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		purchase, err := api.Payments().GetPendingPurchase(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		err = purchase.Confirm()
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		err = api.SessionsMgr().CreateSession(r.Context(), clnt.Id(), 0, 60, 60, nil, 10, 10, false)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		res.SendFlashMsg(w, "success", "Payment received", http.StatusOK)
	}
}

func StartSession(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		err = api.SessionsMgr().Connect(r.Context(), clnt)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		res.SetFlashMsg("success", "Session started")
		res.RedirectToPortal(w)
	}
}
