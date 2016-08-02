package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyDefaults(t *testing.T) {
	input := `<PolicyDefaults></PolicyDefaults>`

    dest := &PolicyDefaults{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyDefaults")

	// Insert specific tests here
}
