package pointers_test

import (
	"testing"

	pointers "github.com/miketmoore/golang-pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewObj(t *testing.T) {
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
