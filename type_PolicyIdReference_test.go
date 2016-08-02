package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyIdReference(t *testing.T) {
	input := `<PolicyIdReference></PolicyIdReference>`

    dest := &PolicyIdReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyIdReference")

	// Insert specific tests here
}
