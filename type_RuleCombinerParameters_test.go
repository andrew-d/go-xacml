package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RuleCombinerParameters(t *testing.T) {
	input := `
	<RuleCombinerParameters RuleIdRef="rule1">
		<CombinerParameter ParameterName="input1">
			<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
		</CombinerParameter>
	</RuleCombinerParameters>`

	dest := &RuleCombinerParameters{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element RuleCombinerParameters")

	assert.Equal(t, "rule1", dest.RuleIdRef)
	if assert.Len(t, dest.CombinerParameters, 1) && assert.Len(t, dest.CombinerParameters[0].AttributeValues, 1) {
		assert.Equal(t, "value1", dest.CombinerParameters[0].AttributeValues[0].Value)
	}
}
