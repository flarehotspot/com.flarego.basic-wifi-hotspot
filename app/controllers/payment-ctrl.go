package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func PaymentRecevied(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res := api.Http().VueResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		var paymentSettings utils.PaymentSettings

		err = api.Config().Plugin().Get(&paymentSettings)
		if err != nil {
			res.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		purchase, err := api.Payments().GetPendingPurchase(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		timeSamples := make([]int, len(paymentSettings))
		dataPoints := make([]int, len(paymentSettings))

		for i, value := range paymentSettings {
			timeSamples[i] = value.TimeMins
			dataPoints[i] = value.DataMb
		}

		purchaseState, err := purchase.State()
		if err != nil {
			res.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if purchaseState.TotalPayment > 0 {

			totalData, totalTime := utils.DivideIntoTimeData(float64(purchaseState.TotalPayment), paymentSettings)

			err = api.SessionsMgr().CreateSession(r.Context(), clnt.Id(), 0, uint(totalTime), float64(totalData), nil, 10, 10, false)
			if err != nil {
				res.Error(w, err.Error(), 500)
				return
			}
		} else {
			log.Println("null")
			return
		}

		err = purchase.Confirm()
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

		err = api.SessionsMgr().Connect(r.Context(), clnt, "Session started")
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		res.RedirectToPortal(w)
	}
}
