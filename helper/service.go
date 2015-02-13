package helper

type Service struct {
	Domain   string
	Backends []Backend
}

type Backend struct {
	Container string
	Server    string
}

