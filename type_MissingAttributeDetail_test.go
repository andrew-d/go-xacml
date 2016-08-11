package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MissingAttributeDetail(t *testing.T) {
	input := `
	<MissingAttributeDetail
		Category="category1"
		AttributeId="attribute1"
		DataType="http://www.w3.org/2001/XMLSchema#string"
		Issuer="issuer1">

		<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
	</MissingAttributeDetail>`

	dest := &MissingAttributeDetail{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element MissingAttributeDetail")

	assert.Equal(t, "category1", dest.Category)
	assert.Equal(t, "attribute1", dest.AttributeId)
	assert.Equal(t, "http://www.w3.org/2001/XMLSchema#string", dest.DataType)
	assert.Equal(t, "issuer1", *dest.Issuer)

	assert.Equal(t, "value1", dest.AttributeValue.Value)
}
