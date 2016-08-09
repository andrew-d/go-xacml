package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Target(t *testing.T) {
	input := `
	<Target>
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
		</AnyOf>
		<AnyOf>
			<AllOf>
				<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
					<AttributeValue
						DataType="http://www.w3.org/2001/XMLSchema#string">doctor</AttributeValue>
					<AttributeDesignator
						AttributeId="role"
						Category="urn:oasis:names:tc:xacml:1.0:subject-category:access-subject"
						DataType="http://www.w3.org/2001/XMLSchema#string"
						MustBePresent="false"/>
				</Match>
				<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
					<AttributeValue
						DataType="http://www.w3.org/2001/XMLSchema#string">medical record</AttributeValue>
					<AttributeDesignator
						AttributeId="resource-type"
						Category="urn:oasis:names:tc:xacml:3.0:attribute-category:resource"
						DataType="http://www.w3.org/2001/XMLSchema#string"
						MustBePresent="false"/>
				</Match>
			</AllOf>
		</AnyOf>
	</Target>
	`

	dest := &Target{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Target")

	assert.Len(t, dest.AnyOf, 2)
}
