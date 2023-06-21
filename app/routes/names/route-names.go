package names

import "github.com/flarehotspot/sdk/api/http/router"

const (
	RouteAdminRatesIndex router.PluginRouteName = "admin:rates:index"
	RouteInsertCoin      router.PluginRouteName = "insert-coin"
	RoutePaymentReceived router.PluginRouteName = "payment-received"
	RouteStartSession    router.PluginRouteName = "start-session"
	RouteStopSession     router.PluginRouteName = "stop-session"
)
