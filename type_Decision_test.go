package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Decision(t *testing.T) {
	input := `<Decision></Decision>`

    dest := &Decision{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Decision")

	// Insert specific tests here
}
