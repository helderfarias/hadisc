package main

import (
	"flag"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/drive"
	"github.com/helderfarias/hadisc/util"
	"log"
)

var etcd = flag.String("etcd", util.GetOpt("ETCD_HOST", ""), "Etcd Host")
var proxy = flag.String("proxy", util.GetOpt("DISCOVERY_PROXY", "nginx"), "haproxy or nginx")
var tpl = flag.String("template", "/etc/nginx/template/nginx.tpl", "Template config file")
var conf = flag.String("config", "/etc/nginx/sites-enabled/server.conf", "Config file")

func main() {
	flag.Parse()

	log.Println("Initialize...")
	log.Printf("Proxy %s", *proxy)
	log.Printf("Enviroment %s", *etcd)

	handlerDrive := drive.NewEtcdDrive(*etcd)

	var handlerDiscovery discovery.HandlerDiscovery
	if *proxy == "haproxy" {
		handlerDiscovery = discovery.NewHAProxy(*tpl, *conf)
	} else if *proxy == "nginx" {
		handlerDiscovery = discovery.NewNginx(*tpl, *conf)
	} else {
		log.Panic("No such proxy")
	}

	handlerDrive.BootstrapAndWatch(handlerDiscovery)

	log.Println("Shutdown")
}
