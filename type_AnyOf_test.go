package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AnyOf(t *testing.T) {
	input := `
	<AnyOf>
		<AllOf>
			<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
				<AttributeValue
					DataType="http://www.w3.org/2001/XMLSchema#string">read</AttributeValue>
				<AttributeDesignator
					AttributeId="urn:oasis:names:tc:xacml:1.0:action:action-id"
					Category="urn:oasis:names:tc:xacml:3.0:attribute-category:action"
					DataType="http://www.w3.org/2001/XMLSchema#string"
					MustBePresent="false"/>
			</Match>
		</AllOf>
		<AllOf>
			<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
				<AttributeValue
					DataType="http://www.w3.org/2001/XMLSchema#string">write</AttributeValue>
				<AttributeDesignator
					AttributeId="urn:oasis:names:tc:xacml:1.0:action:action-id"
					Category="urn:oasis:names:tc:xacml:3.0:attribute-category:action"
					DataType="http://www.w3.org/2001/XMLSchema#string"
					MustBePresent="false"/>
			</Match>
		</AllOf>
	</AnyOf>`

	dest := &AnyOf{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AnyOf")

	if assert.Len(t, dest.AllOf, 2) {
		assert.Equal(t, "read", dest.AllOf[0].Matches[0].AttributeValue.Value)

		assert.Equal(t, "write", dest.AllOf[1].Matches[0].AttributeValue.Value)
	}
}
