package driver

type Driver interface {
	Services() map[string]Service
}

type Service struct {
	Domain   string
	Backends []Backend
}

type Backend struct {
	Container string
	Server    string
}

func Create(host string) Driver {
	newDriver := new(EtcdDriver)
	newDriver.Host = host
	return newDriver
}
