package controllers

import (
	"net/http"

	sdkplugin "sdk/api/plugin"
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
			res.SetFlashMsg("error", msg)
			return
		}

		err = api.SessionsMgr().Connect(r.Context(), clnt, "You are now connected to internet.")
		if err != nil {
			res.Error(w, err.Error(), 500)
			return
		}

		data := map[string]string{
			"redirect_url": "http://google.com",
		}

		res.Json(w, data, http.StatusOK)
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
			res.SetFlashMsg("error", err.Error())
			return
		}
		res.RedirectToPortal(w)
	}
}
