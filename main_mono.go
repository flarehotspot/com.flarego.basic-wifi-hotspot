//go:build mono

package wifihotspot

import (
	"log"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes"
	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/web/navs"
	"github.com/flarehotspot/sdk"
	"github.com/flarehotspot/sdk/v1.0.0/api"
)

func Init(_sdk sdk.SDK) {
	sym, err := _sdk.GetVersion(api.VERSION)
	if err != nil {
		log.Println("Unable to get plugin api: ", err)
	}

	apiv1 := sym.(api.IPluginApi)

	routes.PortalRoutes(apiv1)
	routes.AdminRoutes(apiv1)

	navs.SetPortalItems(apiv1)
	navs.AdminWifiRates(apiv1)
	log.Printf("Success loading plugin: %s", apiv1.Name())
}
