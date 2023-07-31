package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/models"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/flash"
)

type PaymentCtrl struct {
	api plugin.IPluginApi
}

func (self *PaymentCtrl) PaymentRecevied(w http.ResponseWriter, r *http.Request) {
	info, err := self.api.PaymentsApi().ParsePaymentInfo(r)
	if err != nil {
		log.Println(err)
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	ctx := r.Context()
	tx, err := self.api.Db().BeginTx(ctx, nil)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}
	defer tx.Rollback()

	clnt, err := self.api.ClientReg().CurrentClient(r)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	amount, err := info.Purchase.PaymentsTotalTx(tx, ctx)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	devId := clnt.Id()
	t := models.SessionTypeTime.ToUint8()

	result, err := self.api.ConfigApi().WifiRates().ComputeSession(clnt.IpAddr(), amount, t)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	net, err := self.api.NetworkApi().FindByIp(clnt.IpAddr())
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	speed, ok := self.api.ConfigApi().Bandwidth().GetConfig(net.Ifname())
	if !ok {
		err = errors.New("unable to get bandwidth config for " + net.Ifname())
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	minutes := result.TimeMins
	mbytes := result.DataMbytes
	exp := self.api.ConfigApi().Sessions().ComputeExpDays(minutes, mbytes)
	var downMbits uint
	var upMbits uint

	if speed.UseGlobal {
		downMbits = speed.GlobalDownMbits
		upMbits = speed.GlobalUpMbits
	} else {
		downMbits = speed.UserDownMbits
		upMbits = speed.UserUpMbits
	}

	_, err = self.api.Models().Session().CreateTx(tx, ctx, devId, t, minutes, float64(mbytes), &exp, downMbits, upMbits)
	if err != nil {
		log.Println("Error creating session: ", err)
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	err = info.Purchase.Confirm(r.Context())
	if err != nil {
		log.Println(err)
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		self.api.HttpApi().Respond().Error(w, r, err)
		return
	}

	log.Printf("Payment Received: \n%+v", info.Purchase)
	self.api.HttpApi().Respond().SetFlashMsg(w, flash.Success, "Session created successfully.")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func NewPaymentCtrl(api plugin.IPluginApi) *PaymentCtrl {
	return &PaymentCtrl{api}
}
