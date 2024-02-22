package controllers

// import (
// 	"errors"
// 	"log"
// 	"net/http"

// 	"github.com/flarehotspot/sdk/api/http"
// 	"github.com/flarehotspot/sdk/api/plugin"
// )

// type PaymentCtrl struct {
// 	api      sdkplugin.PluginApi
// }

// func NewPaymentCtrl(api sdkplugin.PluginApi) *PaymentCtrl {
// 	return &PaymentCtrl{api}
// }

// func (ctrl *PaymentCtrl) PaymentRecevied(w http.ResponseWriter, r *http.Request) {
// 	info, err := ctrl.api.PaymentsApi().ParsePaymentInfo(r)
// 	if err != nil {
// 		log.Println(err)
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	ctx := r.Context()
// 	tx, err := ctrl.api.DbApi().BeginTx(ctx, nil)
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}
// 	defer tx.Rollback()

// 	clnt, err := ctrl.api.ClientReg().CurrentClient(r)
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	amount, err := info.Purchase.PaymentsTotalTx(tx, ctx)
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	devId := clnt.Id()
// 	t := models.SessionTypeTime

// 	result, err := ctrl.api.ConfigApi().WifiRates().ComputeSession(clnt.IpAddr(), amount, t)
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	net, err := ctrl.api.NetworkApi().FindByIp(clnt.IpAddr())
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	speed, ok := ctrl.api.ConfigApi().Bandwidth().GetConfig(net.Ifname())
// 	if !ok {
// 		err = errors.New("unable to get bandwidth config for " + net.Ifname())
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	minutes := result.TimeMins
// 	mbytes := result.DataMbytes
// 	exp := ctrl.api.ConfigApi().Sessions().ComputeExpDays(minutes, mbytes)

// 	_, err = ctrl.api.ModelsApi().Session().CreateTx(tx, ctx, devId, t.ToUint8(), minutes, float64(mbytes), &exp, speed.UserDownMbits, speed.UserUpMbits, speed.UseGlobal)
// 	if err != nil {
// 		log.Println("Error creating session: ", err)
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	err = info.Purchase.Confirm(r.Context())
// 	if err != nil {
// 		log.Println(err)
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		ctrl.errRoute.Redirect(w, r, err)
// 		return
// 	}

// 	log.Printf("Payment Received: \n%+v", info.Purchase)
// 	flash.SetFlashMsg(w, flash.Success, "Session created successfully.")
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }
