package utils

import (
	"sync/atomic"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

var (
	isSubscribed = atomic.Bool{}
)

// Reload portal when session connected or disconnected
func PortalReload(api sdkplugin.PluginApi, clnt sdkconnmgr.ClientDevice) {
	if !isSubscribed.Load() {
		isSubscribed.Store(true)

		connectedCh := clnt.Subscribe("session:connected")
		disconnCh := clnt.Subscribe("session:disconnected")

		for {
			select {
			case <-connectedCh:
				api.Http().VueRouter().ReloadPortalItems(clnt)
			case <-disconnCh:
				api.Http().VueRouter().ReloadPortalItems(clnt)
			}
		}
	}
}
