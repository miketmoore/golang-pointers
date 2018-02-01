package pointers

type Obj struct {
	things []string
}

func New(things []string) Obj {
	return Obj{things: things}
}

func (o *Obj) Things() *[]string {
	return &o.things
}

func (o *Obj) Pop() string {
	things := o.things
	l := len(things)
	thing := things[l-1:][0]
	o.things = things[l:]
	return thing
}

func (o *Obj) Append(s ...string) {
	o.things = append(o.things, s...)
}
