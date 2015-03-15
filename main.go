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
var tpl = flag.String("template", util.GetOpt("DISCOVERY_TPL", "/etc/nginx/template/nginx.tpl"), "Template config file")
var conf = flag.String("config", util.GetOpt("DISCOVERY_CFG", "/etc/nginx/sites-enabled/server.conf"), "Config file")

func main() {
	flag.Parse()

	log.Println("Initialize...")
	log.Printf("Enviroment %s", *etcd)

	handlerDrive := drive.NewEtcdDrive(*etcd)

	var handlerDiscovery discovery.HandlerDiscovery
	if *proxy == "haproxy" {
		log.Printf("Proxy haproxy")
		handlerDiscovery = discovery.NewHAProxy(*tpl, *conf)
	} else if *proxy == "nginx" {
		log.Printf("Proxy nginx")
		handlerDiscovery = discovery.NewNginx(*tpl, *conf)
	} else {
		log.Panic("No such proxy")
	}

	handlerDrive.BootstrapAndWatch(handlerDiscovery)

	log.Println("Shutdown")
}
