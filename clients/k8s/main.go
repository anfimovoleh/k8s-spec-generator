package k8s

type Client interface {
	Namespaces() ([]string, error)
}

type clientImpl struct{}

func New(addr string) Client {
	return &clientImpl{}
}

func (clientImpl) Namespaces() ([]string, error) {
	//TODO implement me
	panic("implement me")
}
