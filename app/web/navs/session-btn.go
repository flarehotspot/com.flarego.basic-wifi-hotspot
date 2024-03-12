package navs

import (
	"context"
	"errors"
	"net/http"

	// "github.com/flarehotspot/com.flarego.basic-wifi-hotspot/app/routes/names"
	connmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	plugin "github.com/flarehotspot/sdk/api/plugin"
)

type SessionBtnNav struct {
	api plugin.PluginApi
	r   *http.Request
}

func NewSessionBtnNav(api plugin.PluginApi, r *http.Request) *SessionBtnNav {
	return &SessionBtnNav{api, r}
}

func (self *SessionBtnNav) IconPath() string {
	return ""
}

func (self *SessionBtnNav) Text() string {
	clnt, err := self.Client()
	if err != nil {
		return err.Error()
	}

	if self.api.SessionsMgr().IsConnected(clnt) {
		return "Pause"
	}

	if self.CanConnect(self.r.Context()) {
		return "Connect"
	}

	return "No Session"
}

func (self *SessionBtnNav) Description() string {
	return ""
}

func (self *SessionBtnNav) Href() string {
	clnt, err := self.Client()
	if err != nil {
		return err.Error()
	}

	if !self.api.SessionsMgr().IsConnected(clnt) {
		if self.CanConnect(self.r.Context()) {
			return self.api.Http().HttpRouter().UrlForRoute("session.start")
		}
		return "/"
	} else {
		return self.api.Http().HttpRouter().UrlForRoute("session.start")
	}
}

func (self *SessionBtnNav) Client() (connmgr.ClientDevice, error) {
	if self.r == nil {
		return nil, errors.New("Session http request is not initialized.")
	}

	clntSym := self.r.Context().Value(sdkhttp.ClientCtxKey)
	clnt, ok := clntSym.(connmgr.ClientDevice)
	if !ok {
		return nil, errors.New("Could not determine client device.")
	}
	return clnt, nil
}

func (self *SessionBtnNav) CanConnect(ctx context.Context) bool {
	clnt, err := self.Client()
	if err != nil {
		return false
	}

	return self.api.SessionsMgr().HasSession(ctx, clnt.Id())
}
