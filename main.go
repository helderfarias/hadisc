package main

import (
	"flag"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/drive"
	"log"
)

var etcd = flag.String("etcd", "http://192.168.59.103:4001", "Etcd Host")
var tpl = flag.String("template", "/etc/template.tpl", "Template config file")
var conf = flag.String("config", "/etc/config.conf", "Config file")

func main() {
	flag.Parse()

	log.Println("Initialize...")

	handlerDrive := drive.NewEtcdDrive(*etcd)

	handlerDiscovery := discovery.NewHAProxy(*tpl, *conf)

	handlerDrive.Watch(handlerDiscovery)

	log.Println("Shutdown")
}
