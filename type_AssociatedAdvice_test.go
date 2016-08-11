package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AssociatedAdvice(t *testing.T) {
	input := `
	<AssociatedAdvice>
		<Advice AdviceId="advice-1">
			<AttributeAssignment
				AttributeId="attribute-1"
				DataType="http://www.w3.org/2001/XMLSchema#string">value-1</AttributeAssignment>
		</Advice>
	</AssociatedAdvice>`

	dest := &AssociatedAdvice{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AssociatedAdvice")

	if assert.Len(t, dest.Advice, 1) {
		assert.Equal(t, "advice-1", dest.Advice[0].AdviceId)
	}
}
