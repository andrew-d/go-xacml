package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ObligationExpressions(t *testing.T) {
	input := `
	<ObligationExpressions>
		<ObligationExpression FulfillOn="Deny" ObligationId="expression-1">
			<AttributeAssignmentExpression AttributeId="attribute-1">
				<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
			</AttributeAssignmentExpression>
		</ObligationExpression>
	</ObligationExpressions>`

	dest := &ObligationExpressions{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element ObligationExpressions")

	if assert.Len(t, dest.ObligationExpressions, 1) {
		assert.Equal(t, "expression-1", dest.ObligationExpressions[0].ObligationId)
	}
}
