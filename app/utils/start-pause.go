package utils

import (
	"net/http"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func TextBtn(api sdkplugin.PluginApi, r *http.Request) string {
	clnt, err := api.Http().GetClientDevice(r)
	if err != nil {
		return err.Error()
	}

	if api.SessionsMgr().IsConnected(clnt) {
		return "Pause Session"
	} else if api.SessionsMgr().HasSession(r.Context(), clnt.Id()) {
		return "Start Session"
	} else {
		return ""
	}
}

func SessionRoute(api sdkplugin.PluginApi, r *http.Request) string {
	clnt, err := api.Http().GetClientDevice(r)
	if err != nil {
		return err.Error()
	}

	if api.SessionsMgr().HasSession(r.Context(), clnt.Id()) {
		if api.SessionsMgr().IsConnected(clnt) {
			return "portal.pause-session"
		} else {
			return "portal.start-session"
		}
	} else {
		return "/"
	}
}
