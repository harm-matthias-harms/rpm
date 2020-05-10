package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type t1 struct {
	A string
	B string
	C string
}

type t2 struct {
	A string
}

func TestConvert(t *testing.T) {
	a := t1{A: "a", B: "b", C: "c"}
	b := new(t2)
	Convert(a, b)
	assert.Equal(t, a.A, b.A)
}
