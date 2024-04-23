package utils

import (
	"sync"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

var (
	subscribers = sync.Map{}
)

// Reload portal when session connected or disconnected
func PortalReload(api sdkplugin.PluginApi, clnt sdkconnmgr.ClientDevice) {
	_, ok := subscribers.Load(clnt.MacAddr())
	if !ok {
		subscribers.Store(clnt.MacAddr(), true)
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
