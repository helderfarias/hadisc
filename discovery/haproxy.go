package discovery

import (
	"github.com/helderfarias/hadisc/driver"
	"log"
)

func GenerateConfig(services map[string]driver.Service) {

	for _, v := range services {
		log.Println("ok --> ", v.Container, v.Server)
	}

	// template = env.get_template('haproxy.cfg.tmpl')
	//    with open("/etc/haproxy.cfg", "w") as f:
	//        f.write(template.render(services=services))

}
