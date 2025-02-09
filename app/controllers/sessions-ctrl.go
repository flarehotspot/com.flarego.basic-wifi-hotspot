package controllers

import (
	"errors"
	"net/http"

	sdkapi "sdk/api"
)

func StartSession(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start
		res := api.Http().HttpResponse()

		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, r, err, http.StatusUnprocessableEntity)
			return
		}

		if api.SessionsMgr().IsConnected(clnt) {
			msg := "You are already connected to the internet."
			err := errors.New(msg)
			api.Http().HttpResponse().Error(w, r, err, http.StatusInternalServerError)
			return
		}

		err = api.SessionsMgr().Connect(r.Context(), clnt, "You are now connected to internet.")
		if err != nil {
			api.Http().HttpResponse().Error(w, r, err, http.StatusInternalServerError)
			return
		}

		api.Http().HttpResponse().FlashMsg(w, r, "Session started successfully.", sdkapi.FlashMsgSuccess)
		http.Redirect(w, r, "http://google.com", http.StatusSeeOther)
	}
}

func PauseSession(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pause
		res := api.Http().HttpResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, r, err, http.StatusUnprocessableEntity)
			return
		}

		if !api.SessionsMgr().IsConnected(clnt) {
			msg := "You are not connected to internet."
			res.FlashMsg(w, r, msg, sdkapi.FlashMsgError)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		msg := "You are now disconnected from internet."
		err = api.SessionsMgr().Disconnect(r.Context(), clnt, msg)
		if err != nil {
			res.FlashMsg(w, r, err.Error(), sdkapi.FlashMsgError)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		api.Http().HttpResponse().FlashMsg(w, r, msg, sdkapi.FlashMsgWarning)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
