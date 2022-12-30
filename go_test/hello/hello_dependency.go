package hello

/*
hello inject dependency
*/
type HelloDependencyInterface interface {
	GetHelloDependency() string
}

type HelloDependency struct {
	s *Something
}

func NewHelloDependency(di_s *Something) HelloDependencyInterface {
	return &HelloDependency{
		s: di_s,
	}
}

func (hd *HelloDependency) GetHelloDependency() string {
	return hd.s.GetHelloSomething()
}

type Something struct {
}

func (ss *Something) GetHelloSomething() string {
	return "hello"
}
