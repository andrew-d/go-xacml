package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyCombinerParameters(t *testing.T) {
	input := `<PolicyCombinerParameters></PolicyCombinerParameters>`

	dest := &PolicyCombinerParameters{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyCombinerParameters")

	// Insert specific tests here
}
