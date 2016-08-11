package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Obligations(t *testing.T) {
	input := `
	<Obligations>
		<Obligation ObligationId="obligation-id" >
			<AttributeAssignment
				AttributeId="attribute-1"
				DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeAssignment>
		</Obligation>
	</Obligations>`

	dest := &Obligations{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Obligations")

	if assert.Len(t, dest.Obligations, 1) {
		assert.Equal(t, "obligation-id", dest.Obligations[0].ObligationId)
		if assert.Len(t, dest.Obligations[0].AttributeAssignments, 1) {
			assert.Equal(t, "attribute-1", dest.Obligations[0].AttributeAssignments[0].AttributeId)
			assert.Equal(t, "value1", dest.Obligations[0].AttributeAssignments[0].Value)
		}
	}
}
