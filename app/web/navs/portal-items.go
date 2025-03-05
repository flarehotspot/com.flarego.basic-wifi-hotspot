package navs

import (
	"net/http"
	sdkplugin "sdk/api"
)

func SetPortalItems(api sdkplugin.IPluginApi) {
	api.Http().Navs().PortalNavsFactory(func(r *http.Request) []sdkplugin.PortalNavItemOpt {
		navs := []sdkplugin.PortalNavItemOpt{
			{
				Label:     "Insert Coin",
				IconFile:  "icons/wifi-logo.png",
				RouteName: "purchase.wifi",
			},
		}

		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			api.Logger().Error(err.Error())
			return navs
		}

		_, hasRunningSession := api.SessionsMgr().CurrSession(clnt)
		if hasRunningSession {
			navs = append(navs, sdkplugin.PortalNavItemOpt{
				Label:     "Pause Session",
				RouteName: "portal.sessions.stop",
			})
		} else {
			_, err = api.SessionsMgr().GetSession(r.Context(), clnt)
			hasSession := err == nil

			if hasSession {
				navs = append(navs, sdkplugin.PortalNavItemOpt{
					Label:     "Start Session",
					RouteName: "portal.sessions.start",
				})
			}
		}

		return navs
	})
}
