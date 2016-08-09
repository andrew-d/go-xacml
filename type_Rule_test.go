package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Rule(t *testing.T) {
	input := `<Rule></Rule>`

	dest := &Rule{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Rule")

	// Insert specific tests here
}
