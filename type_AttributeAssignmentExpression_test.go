package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeAssignmentExpression(t *testing.T) {
	input := `<AttributeAssignmentExpression></AttributeAssignmentExpression>`

	dest := &AttributeAssignmentExpression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeAssignmentExpression")

	// Insert specific tests here
}
