package navs

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/http/contexts"
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

func (self *SessionBtnNav) Handler(r *http.Request) navigation.IPortalItemData {
	return NewSessionBtnNav(self.api, r)
}

func (self *SessionBtnNav) IconPath() string {
	return ""
}

func (self *SessionBtnNav) Text() string {
	clnt, err := self.client()
	if err != nil {
		return err.Error()
	}

	if clnt.IsConnected() {
		return "Pause"
	}

	return "Connect"
}

func (self *SessionBtnNav) Description() string {
	return ""
}

func (self *SessionBtnNav) Href() string {
	return self.api.Utils().UrlForRoute(names.RouteStartSession)
}

func (self *SessionBtnNav) client() (connmgr.IClientDevice, error) {
	if self.r == nil {
		return nil, errors.New("Session http request is not initialized.")
	}

	clntSym := self.r.Context().Value(contexts.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.IClientDevice)
	if !ok {
		return nil, errors.New("Could not determine client device.")
	}
	return clnt, nil
}
