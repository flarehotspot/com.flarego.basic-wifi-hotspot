package controllers

import (
	"sort"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/utils"
)

/*
Calculates the breakdown of time and data based on the payment amount and the provided
payment settings. It iterates over the payment settings in reverse order, starting from the highest
denomination, and deducts the amount from the payment until it can't be deducted anymore. Then, it
accumulates the time and data accordingly.
*/
func divideIntoTimeData(paymentAmount float64, paymentSettings utils.PaymentSettings) (float64, float64) {
	var totalTime, totalData float64

	// Sort paymentSettings by amount in descending order
	sort.Slice(paymentSettings, func(i, j int) bool {
		return paymentSettings[i].Amount > paymentSettings[j].Amount
	})

	for i := range paymentSettings {
		amount := paymentSettings[i].Amount
		time := paymentSettings[i].TimeMins
		data := paymentSettings[i].DataMb

		// Calculate how many times the current amount fits into the remaining payment
		times := int(paymentAmount / amount)

		// Ensure the times don't exceed the available amount
		if times > 0 && float64(times)*amount <= paymentAmount {
			// Deduct the amount used from the total payment
			paymentAmount -= float64(times) * amount

			// Accumulate time and data based on the calculated times
			totalTime += float64(times) * time
			totalData += float64(times) * data
		}
	}

	return totalTime, totalData
}
