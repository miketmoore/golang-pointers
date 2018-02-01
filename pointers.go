package pointers

type Obj struct {
	stuff []string
}

func New(stuff []string) Obj {
	return Obj{stuff: stuff}
}

func (o *Obj) Stuff() *[]string {
	return &o.stuff
}
