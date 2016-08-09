package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicySetIdReference(t *testing.T) {
	input := `<PolicySetIdReference></PolicySetIdReference>`

	dest := &PolicySetIdReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicySetIdReference")

	// Insert specific tests here
}
