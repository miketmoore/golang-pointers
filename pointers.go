package pointers

// Container is a struct that encapsulates a slice of strings
// and exposes various functions for manipulating the slice
type Container struct {
	things []string
}

// New creates an instance of the Container struct and adds a slice of strings
func New(things []string) Container {
	return Container{things: things}
}

// Things returns a pointer to the slice of strings
func (o *Container) Things() []string {
	return o.things
}

// Pop removes the last value from the internal slice of strings and returns this value
func (o *Container) Pop() string {
	things := o.things
	l := len(things)
	thing := things[l-1:][0]
	o.things = things[l:]
	return thing
}

// Push adds a string to the end of the internal slice of strings
func (o *Container) Push(s ...string) {
	o.things = append(o.things, s...)
}

// SetThings updates the internal slice of strings
func (o *Container) SetThings(things []string) {
	o.things = things
}

// Exchange "exchanges" or "swaps" the value located at the specified index, between the two slices
func Exchange(a Container, b Container, indexToExchange int) {
	aThings := a.Things()
	bThings := b.Things()
	if indexToExchange >= 0 && indexToExchange < len(aThings) && len(aThings) == len(bThings) {

		av := aThings[indexToExchange : indexToExchange+1][0]
		bv := bThings[indexToExchange : indexToExchange+1][0]

		aThingsNew := aThings[0:indexToExchange]
		aThingsNew = append(aThingsNew, bv)
		aThingsNew = append(aThingsNew, aThings[indexToExchange+1:]...)

		bThingsNew := bThings[0:indexToExchange]
		bThingsNew = append(bThingsNew, av)
		bThingsNew = append(bThingsNew, bThings[indexToExchange+1:]...)

		a.SetThings(aThings)
		b.SetThings(bThings)
	}
}
