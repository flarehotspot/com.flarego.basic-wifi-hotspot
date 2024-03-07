package controllers

import (
	"errors"
	"net/http"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func StartSession(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if api.SessionsMgr().IsConnected(clnt) {
			msg := "You are already connected to the internet."
			res.SendFlashMsg(w, "error", msg, 500)
			return
		}

		err = api.SessionsMgr().Connect(clnt)
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}
		msg := "You are now connected to internet."
		res.SendFlashMsg(w, "success", msg, 200)
		res.RedirectToPortal(w)
	}
}

func PauseSession(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pause
		res := api.Http().VueResponse()
		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			res.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if !api.SessionsMgr().IsConnected(clnt) {
			msg := "You are not connected to internet."
			res.SendFlashMsg(w, "error", msg, 500)
			res.RedirectToPortal(w)
		}

		msg := "You are now disconnected to internet."
		err = api.SessionsMgr().Disconnect(clnt, errors.New(msg))
		if err != nil {
			res.SendFlashMsg(w, "error", "Cannot disconnect to internet", 500)
			return
		}

		res.RedirectToPortal(w)
	}
}
