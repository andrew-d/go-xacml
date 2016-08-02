package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VariableDefinition(t *testing.T) {
	input := `<VariableDefinition></VariableDefinition>`

    dest := &VariableDefinition{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element VariableDefinition")

	// Insert specific tests here
}
