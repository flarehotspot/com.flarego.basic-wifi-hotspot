package utils

type PaymentSettings []struct {
	Amount   float64 `json:"amount"`
	DataMb   float64 `json:"data_mb"`
	TimeMins float64 `json:"time_mins"`
}

var DefaultPaymentSettings = PaymentSettings{}
