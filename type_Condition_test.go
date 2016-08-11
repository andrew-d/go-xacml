package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Condition(t *testing.T) {
	input := `
	<Condition>
		<Apply FunctionId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
			<Apply FunctionId="urn:oasis:names:tc:xacml:1.0:function:string-one-and-only">
				<AttributeDesignator
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:test"
					Category="urn:oasis:names:tc:xacml:1.0:subject-category:access-subject"
					DataType="http://www.w3.org/2001/XMLSchema#string"
					MustBePresent="false"/>
			</Apply>
			<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">John Doe</AttributeValue>
		</Apply>
	</Condition>`

	dest := &Condition{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Condition")

	if assert.NotNil(t, dest.Expression) {
		cond := dest.Expression.(*Apply)
		assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:function:string-equal", cond.FunctionId)
	}
}
