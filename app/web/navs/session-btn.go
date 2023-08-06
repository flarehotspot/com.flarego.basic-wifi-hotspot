package navs

import (
	"context"
	"errors"
	"net/http"

	"github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	"github.com/flarehotspot/sdk/v1.0.0/api/connmgr"
	"github.com/flarehotspot/sdk/v1.0.0/api/plugin"
	"github.com/flarehotspot/sdk/v1.0.0/utils/contexts"
)

type SessionBtnNav struct {
	api plugin.IPluginApi
	r   *http.Request
}

func NewSessionBtnNav(api plugin.IPluginApi, r *http.Request) *SessionBtnNav {
	return &SessionBtnNav{api, r}
}

func (self *SessionBtnNav) IconPath() string {
	return ""
}

func (self *SessionBtnNav) Text() string {
	clnt, err := self.client()
	if err != nil {
		return err.Error()
	}

	if self.api.ClientMgr().IsConnected(clnt) {
		return "Pause"
	}

	if self.canConnect(self.r.Context()) {
		return "Connect"
	}

	return "No Session"
}

func (self *SessionBtnNav) Description() string {
	return ""
}

func (self *SessionBtnNav) Href() string {
	clnt, err := self.client()
	if err != nil {
		return err.Error()
	}

	if !self.api.ClientMgr().IsConnected(clnt) {
		if self.canConnect(self.r.Context()) {
			return self.api.HttpApi().Router().UrlForRoute(names.RouteStartSession)
		}
		return "/"
	} else {
		return self.api.HttpApi().Router().UrlForRoute(names.RouteStopSession)
	}
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

func (self *SessionBtnNav) canConnect(ctx context.Context) bool {
	clnt, err := self.client()
	if err != nil {
		return false
	}

	return clnt.HasSession(ctx)
}
