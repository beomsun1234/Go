package embedded

type Nike struct {
	Shose
}

func NewNike() *Shose {

	return NewShose().BuildLogo("nike")
}
