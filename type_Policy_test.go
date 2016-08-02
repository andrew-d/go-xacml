package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Policy(t *testing.T) {
	input := `<Policy></Policy>`

    dest := &Policy{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Policy")

	// Insert specific tests here
}
