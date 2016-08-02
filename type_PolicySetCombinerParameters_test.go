package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicySetCombinerParameters(t *testing.T) {
	input := `<PolicySetCombinerParameters></PolicySetCombinerParameters>`

    dest := &PolicySetCombinerParameters{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicySetCombinerParameters")

	// Insert specific tests here
}
