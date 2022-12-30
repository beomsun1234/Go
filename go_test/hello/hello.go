package hello

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) GetHello() string {
	return "hello"
}
