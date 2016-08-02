package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VariableReference(t *testing.T) {
	input := `<VariableReference></VariableReference>`

    dest := &VariableReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element VariableReference")

	// Insert specific tests here
}
