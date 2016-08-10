package xacml

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VariableDefinition(t *testing.T) {
	input := `
	<VariableDefinition VariableId="the-variable">
		<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">doctor</AttributeValue>
	</VariableDefinition>`

	dest := &VariableDefinition{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element VariableDefinition")

	assert.Equal(t, "the-variable", dest.VariableId)

	switch v := dest.Expression.(type) {
	case *AttributeValue:
		assert.Equal(t, "doctor", v.Value)
	default:
		panic(fmt.Sprintf("invalid Expression: %#v", v))
	}
}
