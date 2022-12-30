package hello

type Hello struct {
	hd HelloDependencyInterface
}

func NewHello(hello_di HelloDependencyInterface) *Hello {
	return &Hello{hd: hello_di}
}

func (h *Hello) GetHello() string {
	return h.hd.GetHelloDependency()
}
