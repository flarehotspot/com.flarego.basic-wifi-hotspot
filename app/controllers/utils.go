package controllers

import "github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"

/*
Calculates the breakdown of time and data based on the payment amount and the provided
payment settings. It iterates over the payment settings in reverse order, starting from the highest
denomination, and deducts the amount from the payment until it can't be deducted anymore. Then, it
accumulates the time and data accordingly.
*/
func divideIntoTimeData(paymentAmount float64, paymentSettings utils.PaymentSettings) (float64, float64) {
	var totalTime, totalData float64

	for i := len(paymentSettings) - 1; i >= 0; i-- {
		amount := paymentSettings[i].Amount
		time := paymentSettings[i].TimeMins
		data := paymentSettings[i].DataMb

		for paymentAmount >= amount {
			totalTime += time
			totalData += data
			paymentAmount -= amount
		}
	}

	return totalTime, totalData
}
