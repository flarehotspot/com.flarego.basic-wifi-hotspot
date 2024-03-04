package names

import sdkhttp "github.com/flarehotspot/core/sdk/api/http"

const (
	RouteAdminRatesIndex  sdkhttp.PluginRouteName = "admin:rates:index"
	RouteAdminRatesSave   sdkhttp.PluginRouteName = "admin:rates:save"
	RouteAdminRatesDelete sdkhttp.PluginRouteName = "admin:rates:delete"
	RouteInsertCoin       sdkhttp.PluginRouteName = "insert-coin"
	RoutePaymentReceived  sdkhttp.PluginRouteName = "payment-received"
	RouteStartSession     sdkhttp.PluginRouteName = "start-session"
	RouteStopSession      sdkhttp.PluginRouteName = "stop-session"
	RoutePayments         sdkhttp.PluginRouteName = "save.settings"
)
