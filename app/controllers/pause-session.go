package controllers

import (
	"net/http"

	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func PauseSession(api sdkplugin.PluginApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//code here for pause
	}
}
