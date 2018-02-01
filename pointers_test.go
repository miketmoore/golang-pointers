package pointers_test

import (
	"testing"

	pointers "github.com/miketmoore/golang-pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	o := pointers.New([]string{"a", "b"})
	assert.Equal(t, []string{"a", "b"}, *o.Things())
}

func TestPop(t *testing.T) {
	o := pointers.New([]string{"a"})

	thing := o.Pop()
	assert.Equal(t, "a", thing)

	remainingThings := *o.Things()
	assert.Equal(t, 0, len(remainingThings))
}

func TestAppend(t *testing.T) {
	o := pointers.New([]string{})
	o.Append("a")
	assert.Equal(t, []string{"a"}, *o.Things())
}

func TestPass(t *testing.T) {
	a := pointers.New([]string{"ping"})
	b := pointers.New([]string{"pong"})

	var x, y string

	x = a.Pop()
	y = b.Pop()
	b.Append(x)
	a.Append(y)

	assert.Equal(t, "pong", (*a.Things())[0])
	assert.Equal(t, "ping", (*b.Things())[0])
}

func TestExchange(t *testing.T) {
	a := pointers.New([]string{"a", "b", "c"})
	b := pointers.New([]string{"d", "e", "f"})

	pointers.Exchange(&a, &b, 0)

	assert.Equal(t, []string{"d", "b", "c"}, *a.Things())
	assert.Equal(t, []string{"a", "e", "f"}, *b.Things())
}
