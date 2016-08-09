package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllOf(t *testing.T) {
	input := `
	<AllOf>
		<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
			<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">John Doe</AttributeValue>
			<AttributeDesignator
				AttributeId="urn:oasis:names:tc:xacml:1.0:subject:subject-id"
				Category="urn:oasis:names:tc:xacml:1.0:subject-category:access-subject"
				DataType="http://www.w3.org/2001/XMLSchema#string"
				MustBePresent="false"/>
		</Match>
		<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
			<AttributeValue
				DataType="http://www.w3.org/2001/XMLSchema#anyURI">FakeLocation</AttributeValue>
			<AttributeDesignator
				AttributeId="urn:oasis:names:tc:xacml:1.0:resource:resource-id"
				Category="urn:oasis:names:tc:xacml:3.0:attribute-category:resource"
				DataType="http://www.w3.org/2001/XMLSchema#anyURI"
				MustBePresent="false"/>

		</Match>
	</AllOf>`

	dest := &AllOf{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AllOf")

	if assert.Len(t, dest.Matches, 2) {
		assert.Equal(t, "John Doe", dest.Matches[0].AttributeValue.Value)
		assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:subject:subject-id", dest.Matches[0].AttributeDesignator.AttributeId)

		assert.Equal(t, "FakeLocation", dest.Matches[1].AttributeValue.Value)
		assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:resource:resource-id", dest.Matches[1].AttributeDesignator.AttributeId)
	}
}
