package navs

import (
	"net/http"

	"github.com/flarehotspot/sdk/api/http/navigation"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
)

type SessionBtnNav struct {
	api plugin.IPluginApi
	r   *http.Request
}

func NewSessionBtnNav(api plugin.IPluginApi, r *http.Request) *SessionBtnNav {
	return &SessionBtnNav{api, r}
}

func (nav *SessionBtnNav) Handler(r *http.Request) navigation.IPortalItemData {
	return NewSessionBtnNav(nav.api, r)
}

func (nav *SessionBtnNav) IconPath() string {
	return ""
}

func (nav *SessionBtnNav) Text() string {
	return "Start Session"
}

func (nav *SessionBtnNav) Description() string {
	return ""
}

func (nav *SessionBtnNav) Href() string {
	return nav.api.Utils().UrlForRoute(names.RouteStartSession)
}
