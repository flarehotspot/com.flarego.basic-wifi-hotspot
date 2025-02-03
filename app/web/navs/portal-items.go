package navs

import (
	"net/http"
	sdkplugin "sdk/api"
)

func SetPortalItems(api sdkplugin.IPluginApi) {
	api.Http().Navs().PortalNavsFactory(func(r *http.Request) []sdkplugin.PortalNavItemOpt {
		return []sdkplugin.PortalNavItemOpt{
			{
				Label:     "Insert Coin",
				RouteName: "purchase:wifi",
			},
		}
	})
}
