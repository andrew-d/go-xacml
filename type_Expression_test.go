package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Expression(t *testing.T) {
	input := `<Expression></Expression>`

    dest := &Expression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Expression")

	// Insert specific tests here
}
