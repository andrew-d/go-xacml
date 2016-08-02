package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeAssignment(t *testing.T) {
	input := `<AttributeAssignment></AttributeAssignment>`

    dest := &AttributeAssignment{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeAssignment")

	// Insert specific tests here
}
