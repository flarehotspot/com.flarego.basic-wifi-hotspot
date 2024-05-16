package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
	sdkpayments "github.com/flarehotspot/sdk/api/payments"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func PurchaseWifiSession(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := sdkpayments.PurchaseRequest{
			Sku:                  "wifi-connection",
			Name:                 "WiFi Connection",
			Description:          "Basic Wifi Hotspot",
			AnyPrice:             true,
			CallbackVueRouteName: "portal.purchase-callback",
		}
		api.Payments().Checkout(w, r, p)
	}
}

func PaymentRecevied(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		paymentSettings, err := utils.GetPaymentConfig(api)
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
			totalSecs, totalMbytes := utils.DivideIntoTimeData(float64(purchaseState.TotalPayment), paymentSettings)
			err = api.SessionsMgr().CreateSession(r.Context(), clnt.Id(), 0, uint(totalSecs), float64(totalMbytes), nil, 10, 10, false)
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
