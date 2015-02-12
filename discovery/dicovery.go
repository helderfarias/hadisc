package discovery

import (
	"github.com/helderfarias/hadisc/driver"
)

type Discovery interface {
	GenerateConfig(services map[string]driver.Service)
	ReloadProcess() error
}

func Create(tpl, config string) Discovery {
	newDiscovery := new(HAProxy)
	newDiscovery.TplFile = tpl
	newDiscovery.ConfigFile = config
	return newDiscovery
}
