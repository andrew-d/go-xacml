package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RuleCombinerParameters(t *testing.T) {
	input := `<RuleCombinerParameters></RuleCombinerParameters>`

    dest := &RuleCombinerParameters{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element RuleCombinerParameters")

	// Insert specific tests here
}
