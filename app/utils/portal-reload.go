package utils

import (
	sdkapi "sdk/api"
	"sync"
)

var (
	subscribers = sync.Map{}
)

// Reload portal when session connected or disconnected
func PortalReload(api sdkapi.IPluginApi, clnt sdkapi.IClientDevice) {
	// _, ok := subscribers.Load(clnt.MacAddr())
	// if !ok {
	// 	subscribers.Store(clnt.MacAddr(), true)
	// 	connectedCh := clnt.Subscribe("session:connected")
	// 	disconnCh := clnt.Subscribe("session:disconnected")

	// 	for {
	// 		select {
	// 		case <-connectedCh:
	// 			api.Http().VueRouter().ReloadPortalItems(clnt)
	// 		case <-disconnCh:
	// 			api.Http().VueRouter().ReloadPortalItems(clnt)
	// 		}
	// 	}
	// }
}
