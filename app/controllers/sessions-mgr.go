package controllers

import (
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

		err = api.SessionsMgr().Connect(r.Context(), clnt, "Gets client device")
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}
		msg := "You are now connected to internet."
		res.SetFlashMsg("success", msg)

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
			res.SetFlashMsg("error", msg)
			res.RedirectToPortal(w)
			return
		}

		msg := "You are now disconnected to internet."
		err = api.SessionsMgr().Disconnect(r.Context(), clnt, msg)
		if err != nil {
			res.SetFlashMsg("error", "Cannot disconnect to internet")
			return
		}
		res.RedirectToPortal(w)
	}
}
