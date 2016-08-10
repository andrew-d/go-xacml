package xacml

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeAssignmentExpression(t *testing.T) {
	input := `
	<AttributeAssignmentExpression AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1">
		<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">assignment1</AttributeValue>
	</AttributeAssignmentExpression>
	`

	dest := &AttributeAssignmentExpression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeAssignmentExpression")

	assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1", dest.AttributeId)
	if assert.NotNil(t, dest.Expression) {
		switch v := dest.Expression.(type) {
		case *AttributeValue:
			assert.Equal(t, "assignment1", v.Value)
		default:
			panic(fmt.Sprintf("invalid Expression: %#v", v))
		}
	}
}
