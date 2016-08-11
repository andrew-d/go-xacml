package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdviceExpressions(t *testing.T) {
	input := `
	<AdviceExpressions>
		<AdviceExpression AppliesTo="Deny" AdviceId="advice1">
			<AttributeAssignmentExpression AttributeId="attribute1">
				<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
			</AttributeAssignmentExpression>
		</AdviceExpression>
	</AdviceExpressions>`

	dest := &AdviceExpressions{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AdviceExpressions")

	if assert.Len(t, dest.AdviceExpressions, 1) {
		assert.Equal(t, "advice1", dest.AdviceExpressions[0].AdviceId)
	}
}
