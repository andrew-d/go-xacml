package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ObligationExpression(t *testing.T) {
	input := `
	<ObligationExpression FulfillOn="Deny" ObligationId="expression-1">
		<AttributeAssignmentExpression AttributeId="attribute-1">
			<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
		</AttributeAssignmentExpression>
	</ObligationExpression>`

	dest := &ObligationExpression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element ObligationExpression")

	assert.Equal(t, "Deny", dest.FulfillOn)
	assert.Equal(t, "expression-1", dest.ObligationId)

	if assert.Len(t, dest.AttributeAssignmentExpressions, 1) {
		aae := dest.AttributeAssignmentExpressions[0]
		assert.Equal(t, "attribute-1", aae.AttributeId)

		val := aae.Expression.(*AttributeValue)
		assert.Equal(t, "value1", val.Value)
	}
}
