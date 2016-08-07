package xacml

import (
	"fmt"
	"sort"
	"strings"
)

// Interface for structs that know how to validate their contents.
type Validatable interface {
	// Validate the contents of this structure, recursing into sub-structures
	Validate(errs *Errors)
}

// Validates an object that conforms to the Validatable interface.
func Validate(name string, obj Validatable) *Errors {
	errs := NewErrors(name)
	obj.Validate(errs)
	return errs
}

// Container for multiple errors.
type Errors struct {
	m    map[string][]error
	path string
}

func NewErrors(rootName string) *Errors {
	ret := &Errors{
		m:    make(map[string][]error),
		path: rootName,
	}
	return ret
}

// Adds an error for this field
func (e *Errors) Add(err error) {
	e.m[e.path] = append(e.m[e.path], err)
}

// Adds a formatted error for this field.
func (e *Errors) Addf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	e.Add(err)
}

// Creates a new sub-Errors object by appending the given field name to the
// current path.
func (e *Errors) Sub(fieldName string) *Errors {
	ret := &Errors{
		m:    e.m,
		path: fmt.Sprintf("%s.%s", e.path, fieldName),
	}
	return ret
}

// Creates a new sub-Errors object by appending the given field name to the
// current path, along with an indicator
func (e *Errors) SubN(fieldName string, index int) *Errors {
	ret := &Errors{
		m:    e.m,
		path: fmt.Sprintf("%s.%s[%d]", e.path, fieldName, index),
	}
	return ret
}

// Returns the number of errors contained.
func (e *Errors) NumErrors() int {
	count := 0
	for _, errs := range e.m {
		count += len(errs)
	}

	return count
}

// Returns a list of all field names in sorted order.
func (e *Errors) fields() []string {
	ret := []string{}
	for field, _ := range e.m {
		ret = append(ret, field)
	}

	sort.Strings(ret)
	return ret
}

// Returns a human-readable summary of this errors object, or the empty string
// if no errors are contained.
func (e *Errors) Summary() string {
	lines := []string{}
	for _, field := range e.fields() {
		for _, err := range e.m[field] {
			lines = append(lines, fmt.Sprintf("- %s: %s", field, err))
		}
	}

	return strings.Join(lines, "\n")
}

func (e *Errors) String() string {
	arr := []string{}
	for _, field := range e.fields() {
		for _, err := range e.m[field] {
			arr = append(arr, fmt.Sprintf("%s: %s", field, err))
		}
	}

	return strings.Join(arr, ", ")
}
