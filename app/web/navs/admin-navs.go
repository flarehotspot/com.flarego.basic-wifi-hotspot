package navs

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func SetAdminNavs(api sdkplugin.PluginApi) {
	api.Http().VueRouter().AdminNavsFunc(func(r *http.Request) []sdkhttp.VueAdminNav {
		return []sdkhttp.VueAdminNav{
			{
				Category:  sdkhttp.NavCategoryPayments,
				Label:     "WiFi Hotspot Payment Settings",
				RouteName: "admin.payment-settings",
			},
		}
	})
}
