package discovery

import (
	"github.com/helderfarias/hadisc/driver"
	"log"
	"os"
	"os/exec"
	"text/template"
)

type HAProxy struct {
	TplFile    string
	ConfigFile string
}

func (this *HAProxy) GenerateConfig(services map[string]driver.Service) {
	tpl, parserError := template.ParseFiles(this.TplFile)
	if parserError != nil {
		log.Fatal("Cannot parser template file", parserError)
	}

	fileCreated, _ := os.Create(this.ConfigFile)
	defer fileCreated.Close()

	err := tpl.Execute(fileCreated, services)
	if err != nil {
		log.Println("Cannot create file config", err)
	}
}

func (this *HAProxy) ReloadProcess() error {
	cmd := exec.Command("haproxy", "-D", "-f", this.ConfigFile, "-p", "/var/run/haproxy.pid")

	return cmd.Run()
}
