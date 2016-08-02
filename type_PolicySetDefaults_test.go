package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicySetDefaults(t *testing.T) {
	input := `<PolicySetDefaults></PolicySetDefaults>`

    dest := &PolicySetDefaults{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicySetDefaults")

	// Insert specific tests here
}
