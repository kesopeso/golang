package unexported

type HasUnexported struct {
	A string
	b bool
}

func (hu HasUnexported) IsActivated() bool {
	return hu.b
}
