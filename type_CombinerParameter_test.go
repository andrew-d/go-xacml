package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CombinerParameter(t *testing.T) {
	input := `
	<CombinerParameter ParameterName="input1">
		<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
		<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value2</AttributeValue>
	</CombinerParameter>`

	dest := &CombinerParameter{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element CombinerParameter")

	assert.Equal(t, "input1", dest.ParameterName)

	if assert.Len(t, dest.AttributeValues, 2) {
		assert.Equal(t, "value1", dest.AttributeValues[0].Value)
		assert.Equal(t, "value2", dest.AttributeValues[1].Value)
	}
}
