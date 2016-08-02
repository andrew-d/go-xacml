package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyIdentifierList(t *testing.T) {
	input := `<PolicyIdentifierList></PolicyIdentifierList>`

    dest := &PolicyIdentifierList{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyIdentifierList")

	// Insert specific tests here
}
