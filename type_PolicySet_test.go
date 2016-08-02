package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicySet(t *testing.T) {
	input := `<PolicySet></PolicySet>`

    dest := &PolicySet{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicySet")

	// Insert specific tests here
}
