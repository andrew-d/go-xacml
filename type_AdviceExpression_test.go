package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdviceExpression(t *testing.T) {
	input := `
	<AdviceExpression AppliesTo="Deny" AdviceId="advice1">
		<AttributeAssignmentExpression AttributeId="attribute1">
			<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
		</AttributeAssignmentExpression>
	</AdviceExpression>
	`

	dest := &AdviceExpression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AdviceExpression")

	assert.Equal(t, "advice1", dest.AdviceId)
	assert.Equal(t, "Deny", dest.AppliesTo)

	if assert.Len(t, dest.AttributeAssignmentExpressions, 1) {
		aae := dest.AttributeAssignmentExpressions[0]
		assert.Equal(t, "attribute1", aae.AttributeId)

		val := aae.Expression.(*AttributeValue)
		assert.Equal(t, "value1", val.Value)
	}
}
