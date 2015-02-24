package discovery

import (
	"github.com/helderfarias/hadisc/helper"
	"testing"
)

func TestParseFiles(t *testing.T) {

	ha := NewHAProxy("conf_test.tpl", "conf_test.cfg")

	item1 := helper.Service{}
	item1.Domain = "admin"
	item1.Backends = append(item1.Backends, helper.Backend{Container: "container1", Server: "10.0.10.1:3000"})
	item1.Backends = append(item1.Backends, helper.Backend{Container: "container2", Server: "10.0.10.2:3000"})

	item2 := helper.Service{}
	item2.Domain = "cad"
	item2.Backends = append(item2.Backends, helper.Backend{Container: "container1", Server: "10.0.20.1:3000"})
	item2.Backends = append(item2.Backends, helper.Backend{Container: "container2", Server: "10.0.20.2:3000"})

	services := []helper.Service{item1, item2}

	ha.GenerateConfig(services)
}
