package controllers

import (
	"errors"
	"net/http"

	sdkapi "sdk/api"
)

func StartSession(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start
		res := api.Http().Response()

		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, r, err, http.StatusUnprocessableEntity)
			return
		}

		if api.SessionsMgr().IsConnected(clnt) {
			msg := "You are already connected to the internet."
			err := errors.New(msg)
			api.Http().Response().Error(w, r, err, http.StatusInternalServerError)
			return
		}

		msg := "You are now connected to internet."
		err = api.SessionsMgr().Connect(r.Context(), clnt, msg)
		if err != nil {
			api.Http().Response().Error(w, r, err, http.StatusInternalServerError)
			return
		}

		res.FlashMsg(w, r, msg, sdkapi.FlashMsgSuccess)
		res.RedirectToPortal(w, r)
	}
}

func PauseSession(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pause
		res := api.Http().Response()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, r, err, http.StatusUnprocessableEntity)
			return
		}

		if !api.SessionsMgr().IsConnected(clnt) {
			msg := "You are not connected to internet."
			res.FlashMsg(w, r, msg, sdkapi.FlashMsgError)
			res.RedirectToPortal(w, r)
			return
		}

		msg := "You are now disconnected from internet."
		err = api.SessionsMgr().Disconnect(r.Context(), clnt, msg)
		if err != nil {
			res.FlashMsg(w, r, err.Error(), sdkapi.FlashMsgError)
			res.RedirectToPortal(w, r)
			return
		}

		res.FlashMsg(w, r, msg, sdkapi.FlashMsgWarning)
		res.RedirectToPortal(w, r)
	}
}
