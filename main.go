package main

import (
	"container/list"
	"github.com/helderfarias/hadisc/discovery"
	"github.com/helderfarias/hadisc/driver"
	"log"
	"time"
)

const (
	POLL_TIMEOUT = 5
)

func contains(src, dest *list.List) bool {
	return false
}

func main() {
	driver := driver.Create(driver.ETCD)
	if driver == nil {
		log.Fatal("No such driver")
	}

	for {
		servs := driver.Services("http://10.89.4.165:4001")

		if len(servs) == 0 {
			time.Sleep(POLL_TIMEOUT)
			continue
		}

		log.Println("Configuration changed, reloading HAProxy...")
		discovery.GenerateConfig(servs)

		// ret = call(["./reload-haproxy.sh"])
		//         if ret != 0:
		//             print "reloading haproxy returned: ", ret
		//             time.sleep(POLL_TIMEOUT)
		//             continue
		//         current_services = services

		time.Sleep(POLL_TIMEOUT)
		break
	}
}
