package discovery

import (
	"github.com/helderfarias/hadisc/helper"
	"log"
	"os"
	"text/template"
)

type HandlerDiscovery interface {
	GenerateConfig(services []helper.Service)
	ReloadProcess() (err error)
}

func parseAndCreateConfigFile(tplFile, configFile string, services []helper.Service) {
	tpl, parserError := template.ParseFiles(tplFile)
	if parserError != nil {
		log.Fatal("Cannot parser template file", parserError)
	}

	fileCreated, _ := os.Create(configFile)
	defer fileCreated.Close()

	err := tpl.Execute(fileCreated, services)
	if err != nil {
		log.Println("Cannot create file config", err)
	}
}
