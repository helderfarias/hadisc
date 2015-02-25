package main

import (
	"flag"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/drive"
	"log"
	"os"
)

var etcd = flag.String("etcd", "", "Etcd Host")
var tpl = flag.String("template", "/etc/haproxy/haproxy.tpl", "Template config file")
var conf = flag.String("config", "/etc/haproxy/haproxy.conf", "Config file")

func main() {
	flag.Parse()

	log.Println("Initialize...")

	if *etcd == "" {
		log.Println("No such -etcd' in the command line, try lookup by enviroment")
		*etcd = os.Getenv("ETCD_HOST_DISCOVERY")
		log.Printf("Enviroment %s", *etcd)
	}

	handlerDrive := drive.NewEtcdDrive(*etcd)

	handlerDiscovery := discovery.NewHAProxy(*tpl, *conf)

	handlerDrive.BootstrapAndWatch(handlerDiscovery)

	log.Println("Shutdown")
}
