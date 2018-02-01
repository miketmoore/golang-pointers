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
