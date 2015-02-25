package drive

import (
	"github.com/coreos/go-etcd/etcd"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/helper"
	"log"
	"strings"
)

type EtcdDrive struct {
	Host      string
	EtcClient *etcd.Client
}

func NewEtcdDrive(host string) *EtcdDrive {
	newDrive := new(EtcdDrive)
	newDrive.Host = host
	newDrive.EtcClient = etcd.NewClient([]string{host})
	return newDrive
}

func (this *EtcdDrive) BootstrapAndWatch(handler discovery.HandlerDiscovery) {
	changeChan := make(chan *etcd.Response)
	stopChan := make(chan bool)

	go func() {
		for msg := range changeChan {
			reload := (msg.PrevNode == nil) || (msg.PrevNode.Key != msg.Node.Key) || (msg.Action != "set")
			if reload {
				services := services(this.EtcClient)

				handler.GenerateConfig(services)

				log.Println("Reloading process")
				err := handler.ReloadProcess()
				if err != nil {
					log.Println("Error reload process", err)
				}
			}
		}
	}()

	log.Println("Bootstrap")
	bootstrap(handler, this.EtcClient)

	log.Println("Start watching changes in etcd")
	if _, err := this.EtcClient.Watch("/services", 0, true, changeChan, stopChan); err != nil {
		log.Println("Cannot register watcher for changes in etcd: ", err)
	}
}

func bootstrap(handler discovery.HandlerDiscovery, client *etcd.Client) {
	services := services(client)

	handler.GenerateConfig(services)
}

func services(client *etcd.Client) []helper.Service {
	resp, err := client.Get("services", true, true)
	if err != nil {
		log.Println(err)
		return nil
	}

	services := make([]helper.Service, 0)

	for _, keys := range resp.Node.Nodes {
		service := helper.Service{}

		for _, app := range keys.Nodes {
			appType := app.Key[strings.LastIndex(app.Key, "/")+1:]

			if appType == "domain" {
				for _, item := range app.Nodes {
					service.Domain = makeDomain(app.Key)
					service.Endpoint = item.Value
				}
			}

			if appType == "backend" {
				for _, item := range app.Nodes {
					container := item.Key[strings.LastIndex(item.Key, "/")+1:]

					if item.Nodes != nil && item.Nodes.Len() != 0 {
						server := item.Nodes[0].Value
						service.Backends = append(service.Backends, helper.Backend{Container: container, Server: server})
					}
				}
			}
		}

		services = append(services, service)
	}

	return services
}

func makeDomain(value string) (endpoint string) {
	if !strings.Contains(value, "/") {
		return ""
	}

	service := value[:strings.LastIndex(value, "/")]
	service = service[strings.LastIndex(service, "/")+1:]
	return service
}
