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

		w.Write([]byte("must start session"))

		// data := map[string]string{
		// 	"redirect_url": "http://google.com",
		// }

		// res.Json(w, data, http.StatusOK)
	}
}

func PauseSession(api sdkapi.IPluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pause
		// res := api.Http().VueResponse()
		// clnt, err := api.Http().GetClientDevice(r)
		// if err != nil {
		// 	res.Error(w, err.Error(), http.StatusUnprocessableEntity)
		// 	return
		// }

		// if !api.SessionsMgr().IsConnected(clnt) {
		// 	msg := "You are not connected to internet."
		// 	res.SetFlashMsg("error", msg)
		// 	res.RedirectToPortal(w)
		// 	return
		// }

		// msg := "You are now disconnected to internet."
		// err = api.SessionsMgr().Disconnect(r.Context(), clnt, msg)
		// if err != nil {
		// 	res.SetFlashMsg("error", err.Error())
		// 	return
		// }
		// res.RedirectToPortal(w)
	}
}
