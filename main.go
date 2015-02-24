package main

import (
	"flag"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/drive"
	"log"
)

var etcd = flag.String("etcd", "http://localhost:4001", "Etcd Host")
var tpl = flag.String("template", "/etc/haproxy/haproxy.tpl", "Template config file")
var conf = flag.String("config", "/etc/haproxy/haproxy.conf", "Config file")

func main() {
	flag.Parse()

	log.Println("Initialize...")

	handlerDrive := drive.NewEtcdDrive(*etcd)

	handlerDiscovery := discovery.NewHAProxy(*tpl, *conf)

	handlerDrive.BootstrapAndWatch(handlerDiscovery)

	log.Println("Shutdown")
}
