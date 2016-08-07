package xacml

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsBasic(t *testing.T) {
	e := NewErrors("foo")
	e.Add(errors.New("invalid bar"))

	assert.Equal(t, 1, e.NumErrors())
	assert.Equal(t, "- foo: invalid bar", e.Summary())
	assert.Equal(t, "foo: invalid bar", e.String())

	e.Addf("bad asdf%d", 1)
	assert.Equal(t, "- foo: invalid bar\n- foo: bad asdf1", e.Summary())
	assert.Equal(t, "foo: invalid bar, foo: bad asdf1", e.String())
}

func TestErrorsSub(t *testing.T) {
	e := NewErrors("root")

	s1 := e.Sub("field1")
	s2 := e.Sub("field2")
	s3 := e.SubN("array", 1)

	s1.Addf("one")
	s2.Addf("two")
	s3.Addf("three")

	// Note: these should be sorted
	assert.Equal(t, "- root.array[1]: three\n- root.field1: one\n- root.field2: two", e.Summary())
}

type isValidatable struct{}

func (v isValidatable) Validate(errs *Errors) {
	errs.Addf("foobar")
}

func TestValidate(t *testing.T) {
	e := Validate("root", isValidatable{})
	assert.Equal(t, "root: foobar", e.String())
}
