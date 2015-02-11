package driver

const (
	ETCD = 1
)

type Service struct {
	Container string
	Server    string
}

type Driver interface {
	Services(host string) map[string]Service
}

func Create(driverType int) Driver {
	if driverType == ETCD {
		return new(EtcdDriver)
	}

	return nil
}
