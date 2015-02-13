package discovery

import (
	"github.com/helderfarias/hadisc/helper"
	"log"
	"os"
	"os/exec"
	"text/template"
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
	tpl, parserError := template.ParseFiles(this.TplFile)
	if parserError != nil {
		log.Fatal("Cannot parser template file", parserError)
	}

	fileCreated, _ := os.Create(this.ConfigFile)
	defer fileCreated.Close()

	backends := make([]helper.Backend, 0)
	for _, s := range services {
		for _, b := range s.Backends {
			if len(b.Container) > 0 && len(b.Server) > 0 {
				backends = append(backends, b)
			}
		}
	}

	err := tpl.Execute(fileCreated, backends)
	if err != nil {
		log.Println("Cannot create file config", err)
	}
}

func (this *HAProxy) ReloadProcess() (err error) {
	cmd := exec.Command("haproxy", "-D", "-f", this.ConfigFile, "-p", "/var/run/haproxy.pid")

	return cmd.Run()
}
