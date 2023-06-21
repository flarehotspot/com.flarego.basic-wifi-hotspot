package controllers

import (
	"log"
	"net/http"

	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/constants"
	"github.com/flarehotspot/sdk/utils/http/req"
)

type PaymentCtrl struct {
	api plugin.IPluginApi
}

func (self *PaymentCtrl) PaymentRecevied(w http.ResponseWriter, r *http.Request) {
	info, err := self.api.PaymentsApi().ParsePaymentInfo(r)
	if err != nil {
		log.Println(err)
		self.api.HttpApi().Respond().Error(w, err)
		return
	}

	ctx := r.Context()
	tx, err := self.api.Db().BeginTx(ctx, nil)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, err)
		return
	}
	defer tx.Rollback()

	clnt, err := req.ClientDevice(r)
	if err != nil {
		self.api.HttpApi().Respond().Error(w, err)
		return
	}

	devId := clnt.Id()
	stype := constants.SessionTypeTime.ToUint8()
	_, err = self.api.Models().Session().CreateTx(tx, ctx, devId, stype, 100, 0, 0, 0, nil, 111, 222)
	if err != nil {
		log.Println("Error creating session: ", err)
		self.api.HttpApi().Respond().Error(w, err)
		return
	}

	err = info.Purchase.Confirm(r.Context())
	if err != nil {
		log.Println(err)
		self.api.HttpApi().Respond().Error(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		self.api.HttpApi().Respond().Error(w, err)
		return
	}

	log.Printf("Payment Received: \n%+v", info.Purchase)
	self.api.HttpApi().Respond().SetFlashMsg(w, constants.FlashTypeSuccess, "Session created successfully.")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func NewPaymentCtrl(api plugin.IPluginApi) *PaymentCtrl {
	return &PaymentCtrl{api}
}
