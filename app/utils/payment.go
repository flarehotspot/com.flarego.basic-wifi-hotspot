package utils

// import (
	// "net/http"

	// "github.com/flarehotspot/core/sdk/api/connmgr"
	// "github.com/flarehotspot/core/sdk/api/plugin"
	// "github.com/flarehotspot/core/sdk/utils/http/req"
// )

// func MakeSession(api plugin.IPluginApi, clnt connmgr.IClientDevice, amount float64) error {
	// rates, err := api.ConfigApi().WifiRates().All()
	// if err != nil {
		// return err
	// }
// }

// func PaymentReceived(api plugin.IPluginApi, r *http.Request, amount float64) error {
	// clnt, err := req.ClientDevice(r)
	// if err != nil {
		// return err
	// }

	// info, err := api.PaymentsApi().ParsePaymentInfo(r)
	// if err != nil {
		// return err
	// }

	// pur := info.Purchase

	// ctx := r.Context()
	// tx, err := api.Db().BeginTx(ctx, nil)
	// if err != nil {
		// return err
	// }
	// defer tx.Rollback()

	// return tx.Commit()
// }
