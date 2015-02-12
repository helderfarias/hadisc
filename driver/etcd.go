package driver

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
	"strings"
)

type EtcdDriver struct {
	Host string
}

func (this *EtcdDriver) Services() map[string]Service {
	client := etcd.NewClient([]string{this.Host})

	resp, err := client.Get("services", true, true)
	if err != nil {
		log.Fatal(err)
	}

	services := make(map[string]Service)

	for _, keys := range resp.Node.Nodes {
		service := Service{}

		for _, app := range keys.Nodes {
			appType := app.Key[strings.LastIndex(app.Key, "/")+1:]

			if appType == "domain" {
				for _, item := range app.Nodes {
					service.Domain = item.Value
				}
			}

			if appType == "backend" {
				for _, item := range app.Nodes {
					container := item.Key[strings.LastIndex(item.Key, "/")+1:]
					server := item.Nodes[0].Value
					service.Backends = append(service.Backends, Backend{Container: container, Server: server})
				}
			}
		}

		services[service.Domain] = service
	}

	return services
}
