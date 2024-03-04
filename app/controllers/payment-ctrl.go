package controllers

import (
	"log"
	"net/http"
	"strconv"

	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func divideIntoTensFivesOnes(n int) (int, int, int) {
	tens := n / 10
	remainder := n % 10
	fives := remainder / 5
	ones := remainder % 5
	return tens, fives, ones
}

func PaymentReceived(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetDevice(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var savedData UserInput

		err = api.Config().Plugin().ReadJson(&savedData)
		if err != nil {
			http.Error(w, "Unable to read JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		purchase, err := api.Payments().GetPendingPurchase(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		timeSamples := make([]int, 3)
		dataPoints := make([]int, 3)

		for i, out := range savedData.Out {
			timeSamples[i], _ = strconv.Atoi(out.Time)
			dataPoints[i] = out.Data
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

		res.FlashMsg("success", "Payment received")
		res.Json(w, nil, 200)
	}
}
func StartSession(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetDevice(r)
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
