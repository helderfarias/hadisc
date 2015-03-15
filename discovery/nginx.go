package discovery

import (
	"github.com/helderfarias/hadisc/helper"
	"os/exec"
)

type Nginx struct {
	TplFile    string
	ConfigFile string
}

func NewNginx(tpl, config string) *Nginx {
	newDiscovery := new(Nginx)
	newDiscovery.TplFile = tpl
	newDiscovery.ConfigFile = config
	return newDiscovery
}

func (this *Nginx) GenerateConfig(services []helper.Service) {
	parseAndCreateConfigFile(this.TplFile, this.ConfigFile, services)
}

func (this *Nginx) ReloadProcess() (err error) {
	cmd := exec.Command("nginx", "-s", "reload")

	return cmd.Run()
}
