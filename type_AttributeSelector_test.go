package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeSelector(t *testing.T) {
	input := `<AttributeSelector></AttributeSelector>`

	dest := &AttributeSelector{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeSelector")

	// Insert specific tests here
}
