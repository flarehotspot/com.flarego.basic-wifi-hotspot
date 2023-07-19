package navs

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/http/navigation"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
)

func AdminWifiRates(api plugin.IPluginApi) {
	wifiRatesNav := &AdminNav{
		category: navigation.CategoryPayments,
		text:     "Wifi Rates",
		href:     api.HttpApi().Router().UrlForRoute(names.RouteAdminRatesIndex),
	}

	api.NavApi().AdminNavsFn(func(r *http.Request) []navigation.IAdminNavItem {
		navs := []navigation.IAdminNavItem{}
		navs = append(navs, wifiRatesNav)
		return navs
	})
}
