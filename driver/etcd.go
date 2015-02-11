package driver

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
	"strings"
)

type EtcdDriver struct {
}

func (e *EtcdDriver) Services(host string) map[string]Service {
	client := etcd.NewClient([]string{host})

	resp, err := client.Get("backends", true, true)
	if err != nil {
		log.Fatal(err)
	}

	services := make(map[string]Service)

	for _, node := range resp.Node.Nodes {
		if node.Dir {
			for _, i := range node.Nodes {
				if strings.Count(i.Key[1:], "/") != 2 {
					continue
				}

				name := strings.Split(i.Key[1:], "/")[1]
				server := i.Value
				services[name] = Service{Container: name, Server: server}
			}
		}
	}

	return services
}
