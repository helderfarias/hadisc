package helper

type Service struct {
	Domain   string
	Endpoint string
	Backends []Backend
}

type Backend struct {
	Container string
	Server    string
}
