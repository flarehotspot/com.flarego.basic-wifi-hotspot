package navs

import (
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/v1/api"
	"github.com/flarehotspot/sdk/v1/api/http/navigation"
)

func AdminWifiRates(API api.IPluginApi) {
	wifiRatesNav := &AdminNav{
		category: navigation.CategoryPayments,
		text:     "Wifi Rates",
		href:     API.HttpApi().Router().UrlForRoute(names.RouteAdminRatesIndex),
	}

	API.NavApi().AdminNavsFn(func(r *http.Request) []navigation.IAdminNavItem {
		navs := []navigation.IAdminNavItem{}
		navs = append(navs, wifiRatesNav)
		return navs
	})
}
