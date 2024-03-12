package utils

import "sort"

type PaymentSettings []struct {
	Amount   float64 `json:"amount"`
	DataMb   int     `json:"data_mb"`
	TimeMins int     `json:"time_mins"`
}

var DefaultPaymentSettings = PaymentSettings{}

/*
Calculates the breakdown of time and data based on the payment amount and the provided
payment settings. It iterates over the payment settings in reverse order, starting from the highest
denomination, and deducts the amount from the payment until it can't be deducted anymore. Then, it
accumulates the time and data accordingly.
*/
func DivideIntoTimeData(paymentAmount float64, paymentSettings PaymentSettings) (int, int) {
	var totalTime, totalData int

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
			totalTime += times * time
			totalData += times * data
		}
	}

	return totalTime, totalData
}
