package navs

import "github.com/flarehotspot/core/sdk/api/http/navigation"

type AdminNav struct {
	category navigation.INavCategory
	text     string
	href     string
}

func (self *AdminNav) Category() navigation.INavCategory {
	return self.category
}

func (self *AdminNav) Text() string {
	return self.text
}

func (self *AdminNav) Href() string {
	return self.href
}
