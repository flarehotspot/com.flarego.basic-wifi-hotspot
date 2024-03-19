package navs

import (
	sdkacct "github.com/flarehotspot/sdk/api/accounts"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func SetAdminNavs(api sdkplugin.PluginApi) {
	api.Http().VueRouter().AdminNavsFunc(func(acct sdkacct.Account) []sdkhttp.VueAdminNav {
		return []sdkhttp.VueAdminNav{
			{
				Category:  sdkhttp.NavCategoryPayments,
				Label:     "WiFi Hotspot Payment Settings",
				RouteName: "admin.payment-settings",
			},
		}
	})
}
