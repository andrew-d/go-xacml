package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Obligation(t *testing.T) {
	input := `
	<Obligation ObligationId="obligation-id" >
		<AttributeAssignment
			AttributeId="attribute-1"
			DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeAssignment>
	</Obligation>`

	dest := &Obligation{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Obligation")

	assert.Equal(t, "obligation-id", dest.ObligationId)
	if assert.Len(t, dest.AttributeAssignments, 1) {
		assert.Equal(t, "attribute-1", dest.AttributeAssignments[0].AttributeId)
		assert.Equal(t, "value1", dest.AttributeAssignments[0].Value)
	}
}
