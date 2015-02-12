package main

import (
	"flag"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/driver"
	"log"
	"time"
)

const (
	POLL_TIMEOUT = 5
)

var etcd = flag.String("etcd", "http://192.168.59.103:4001", "Etcd Host")
var tpl = flag.String("template", "/etc/template.tpl", "Template config file")
var conf = flag.String("config", "/etc/config.conf", "Config file")

func main() {
	flag.Parse()
	driver := driver.Create(*etcd)
	discovery := discovery.Create(*tpl, *conf)

	for {
		servs := driver.Services()

		if len(servs) == 0 {
			time.Sleep(POLL_TIMEOUT)
			continue
		}

		log.Println("Configuration changed")
		discovery.GenerateConfig(servs)

		log.Println("Reloading...")
		err := discovery.ReloadProcess()
		if err != nil {
			log.Println(err)
			time.Sleep(POLL_TIMEOUT)
			continue
		}

		time.Sleep(POLL_TIMEOUT)
	}

}
