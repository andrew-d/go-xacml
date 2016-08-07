package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Match(t *testing.T) {
	input := `
	<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
		<AttributeValue
			DataType="http://www.w3.org/2001/XMLSchema#string">read</AttributeValue>
		<AttributeDesignator
			AttributeId="urn:oasis:names:tc:xacml:1.0:action:action-id"
			Category="urn:oasis:names:tc:xacml:3.0:attribute-category:action"
			DataType="http://www.w3.org/2001/XMLSchema#string"
			MustBePresent="false"/>
	</Match>`

	dest := &Match{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Match")

	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:function:string-equal", dest.MatchId)
	assert.Equal(t, "read", dest.AttributeValue.Value)

	assert.Nil(t, dest.AttributeSelector)
	if assert.NotNil(t, dest.AttributeDesignator) {
		assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:action:action-id", dest.AttributeDesignator.AttributeId)
	}
}
