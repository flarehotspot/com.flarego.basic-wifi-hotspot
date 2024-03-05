package utils

type PaymentSettings []struct {
	Amount   int `json:"amount"`
	DataMb   int `json:"data_mb"`
	TimeMins int `json:"time_mins"`
}

var DefaultPaymentSettings = PaymentSettings{}
