package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Advice(t *testing.T) {
	input := `
	<Advice
		AdviceId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:Advice-1">
		<AttributeAssignment
			AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1"
			DataType="http://www.w3.org/2001/XMLSchema#string">assignment1</AttributeAssignment>
		<AttributeAssignment
			AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicSingleValue"
			DataType="http://www.w3.org/2001/XMLSchema#string">J. Hibbert</AttributeAssignment>
	</Advice>`

	dest := &Advice{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Advice")

	// Insert specific tests here
	assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:Advice-1", dest.AdviceId)

	if assert.Len(t, dest.Assignments, 2) {
		a1 := dest.Assignments[0]
		assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1", a1.AttributeId)
		assert.Equal(t, "http://www.w3.org/2001/XMLSchema#string", a1.DataType)
		assert.Equal(t, "assignment1", a1.Value)

		a2 := dest.Assignments[1]
		assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicSingleValue", a2.AttributeId)
		assert.Equal(t, "http://www.w3.org/2001/XMLSchema#string", a2.DataType)
		assert.Equal(t, "J. Hibbert", a2.Value)
	}
}
