package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyIssuer(t *testing.T) {
	input := `<PolicyIssuer></PolicyIssuer>`

    dest := &PolicyIssuer{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyIssuer")

	// Insert specific tests here
}
