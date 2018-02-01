package pointers

type Container struct {
	things []string
}

func New(things []string) Container {
	return Container{things: things}
}

func (o *Container) Things() *[]string {
	return &o.things
}

func (o *Container) Pop() string {
	things := o.things
	l := len(things)
	thing := things[l-1:][0]
	o.things = things[l:]
	return thing
}

func (o *Container) Append(s ...string) {
	o.things = append(o.things, s...)
}

func (o *Container) SetThings(things *[]string) {
	o.things = *things
}

func Exchange(a *Container, b *Container, indexToExchange int) {
	aThings := *a.Things()
	bThings := *b.Things()
	if indexToExchange >= 0 && indexToExchange < len(aThings) && len(aThings) == len(bThings) {

		av := aThings[indexToExchange : indexToExchange+1][0]
		bv := bThings[indexToExchange : indexToExchange+1][0]

		aThingsNew := aThings[0:indexToExchange]
		aThingsNew = append(aThingsNew, bv)
		aThingsNew = append(aThingsNew, aThings[indexToExchange+1:]...)

		bThingsNew := bThings[0:indexToExchange]
		bThingsNew = append(bThingsNew, av)
		bThingsNew = append(bThingsNew, bThings[indexToExchange+1:]...)

		a.SetThings(&aThings)
		b.SetThings(&bThings)
	}
}
