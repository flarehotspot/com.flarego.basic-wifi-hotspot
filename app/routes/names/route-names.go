package names

import "github.com/flarehotspot/sdk/v1/api/http/router"

const (
	RouteAdminRatesIndex  router.PluginRouteName = "admin:rates:index"
	RouteAdminRatesSave   router.PluginRouteName = "admin:rates:save"
	RouteAdminRatesDelete router.PluginRouteName = "admin:rates:delete"
	RouteInsertCoin       router.PluginRouteName = "insert-coin"
	RoutePaymentReceived  router.PluginRouteName = "payment-received"
	RouteStartSession     router.PluginRouteName = "start-session"
	RouteStopSession      router.PluginRouteName = "stop-session"
)
