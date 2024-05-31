package navs

import (
	sdkacct "sdk/api/accounts"
	sdkhttp "sdk/api/http"
	sdkplugin "sdk/api/plugin"
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
