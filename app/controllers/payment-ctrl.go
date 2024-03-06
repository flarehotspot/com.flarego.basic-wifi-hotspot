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

		//read the json
		var paymentData utils.PaymentSettings
		err = api.Config().Plugin().ReadJson(&paymentData)
		if err != nil {
			res.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		//test the json
		log.Printf("%+v", paymentData)

		//iterate the value
		timeSlice := make([]int, 5)
		dataSlice := make([]int, 5)

		for k, v := range paymentData {
			timeSlice[k] = v.TimeMins
			dataSlice[k] = v.DataMb
		}

		purchase, err := api.Payments().GetPendingPurchase(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		purchaseState, err := purchase.State()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if purchaseState.TotalPayment > 0 {
			tens, fives, ones := divideIntoTensFivesOnes(int(purchaseState.TotalPayment))
			totaldata := tens*dataSlice[2] + fives*dataSlice[1] + ones*dataSlice[0]
			totalamount := tens*timeSlice[2] + fives*timeSlice[1] + ones*timeSlice[0]

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

		res.FlashMsg("success", "Payment received")
		res.Json(w, nil, 200)
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

		err = api.SessionsMgr().Connect(clnt)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		res.FlashMsg("success", "Session started")
		res.RedirectToPortal(w)
	}
}
