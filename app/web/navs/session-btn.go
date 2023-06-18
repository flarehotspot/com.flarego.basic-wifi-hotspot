package navs

import (
	"errors"
	"net/http"

	"github.com/flarehotspot/sdk/api/connmgr"
	"github.com/flarehotspot/sdk/api/plugin"
	"github.com/flarehotspot/sdk/utils/constants"
	"github.com/flarehotspot/wifi-hotspot/app/routes/names"
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

	if self.canConnect() {
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
		if self.canConnect() {
			return self.api.Utils().UrlForRoute(names.RouteStartSession)
		}
		return "/"
	} else {
		return self.api.Utils().UrlForRoute(names.RouteStopSession)
	}
}

func (self *SessionBtnNav) client() (connmgr.IClientDevice, error) {
	if self.r == nil {
		return nil, errors.New("Session http request is not initialized.")
	}

	clntSym := self.r.Context().Value(constants.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.IClientDevice)
	if !ok {
		return nil, errors.New("Could not determine client device.")
	}
	return clnt, nil
}

func (self *SessionBtnNav) canConnect() bool {
	clnt, err := self.client()
	if err != nil {
		return false
	}

	return clnt.HasSession()
}
