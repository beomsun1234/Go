package embedded

type Shose struct {
	logo string
	size int
}


func NewShose() *Shose {
	return &Shose{}
}

func (s *Shose) GetSize() int {
	return s.size
}

func (s *Shose) GetLogo() string {
	return s.logo
}

func (s *Shose) BuildLogo(logo string) *Shose {
	s.logo = logo
	return s
}

func (s *Shose) BuildSize(size int) *Shose {
	s.size = size
	return s
}
