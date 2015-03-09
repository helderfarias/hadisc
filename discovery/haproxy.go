package discovery

import (
	"github.com/helderfarias/hadisc/helper"
	"os/exec"
)

type HAProxy struct {
	TplFile    string
	ConfigFile string
}

func NewHAProxy(tpl, config string) *HAProxy {
	newDiscovery := new(HAProxy)
	newDiscovery.TplFile = tpl
	newDiscovery.ConfigFile = config
	return newDiscovery
}

func (this *HAProxy) GenerateConfig(services []helper.Service) {
	parseAndCreateConfigFile(this.TplFile, this.ConfigFile, services)
}

func (this *HAProxy) ReloadProcess() (err error) {
	cmd := exec.Command("supervisorctl", "restart", "haproxy")

	return cmd.Run()
}
