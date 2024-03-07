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

		var paymentData utils.PaymentSettings
		err = api.Config().Plugin().ReadJson(&paymentData)
		if err != nil {
			res.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		purchase, err := api.Payments().GetPendingPurchase(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}
		timeSamples := make([]int, 5)
		dataPoints := make([]int, 5)

		for i, value := range paymentData {
			timeSamples[i] = value.TimeMins
			dataPoints[i] = value.DataMb
		}

		purchaseState, err := purchase.State()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if purchaseState.TotalPayment > 0 {
			tens, fives, ones := divideIntoTensFivesOnes(int(purchaseState.TotalPayment))
			totaldata := tens*dataPoints[2] + fives*dataPoints[1] + ones*dataPoints[0]
			totalamount := tens*timeSamples[2] + fives*timeSamples[1] + ones*timeSamples[0]

			err = api.SessionsMgr().CreateSession(r.Context(), clnt.Id(), 0, uint(float64(totalamount)), float64(totaldata), nil, 10, 10, false)
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
