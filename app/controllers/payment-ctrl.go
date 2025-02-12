package controllers

import (
	"fmt"
	"log"
	"net/http"

	sdkapi "sdk/api"

	"com.flarego.basic-wifi-hotspot/app/utils"
)

func PurchaseWifiSession(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := sdkapi.PurchaseRequest{
			Sku:           "wifi-connection",
			Name:          "WiFi Connection",
			Description:   "Basic Wifi Hotspot",
			AnyPrice:      true,
			CallbackRoute: "purchase.wifi.callback",
		}
		api.Payments().Checkout(w, r, p)
	}
}

func PaymentRecevied(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().Response()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		paymentSettings, err := utils.GetPaymentConfig(api)
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		purchase, err := api.Payments().GetPurchaseRequest(r)
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
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
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		if purchaseState.TotalPayment > 0 {
			totalSecs, totalMbytes := utils.DivideIntoTimeData(float64(purchaseState.TotalPayment), paymentSettings)
			fmt.Printf("\n*******************\nSession Total time: %d, Total data: %d\n", totalSecs, totalMbytes)
			err = api.SessionsMgr().CreateSession(r.Context(), clnt.Id(), sdkapi.SessionTypeTime, totalSecs, float64(totalMbytes), nil, 10, 10, false)
			if err != nil {
				res.Error(w, r, err, http.StatusInternalServerError)
				return
			}
		} else {
			log.Println("null")
			return
		}

		err = purchase.Confirm()
		if err != nil {
			res.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		api.Http().Response().FlashMsg(w, r, "Payment successful.", sdkapi.FlashMsgSuccess)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
