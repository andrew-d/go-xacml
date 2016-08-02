package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeValue(t *testing.T) {
	input := `<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">doctor</AttributeValue>`

	dest := &AttributeValue{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeValue")

	// Insert specific tests here
	assert.Equal(t, "doctor", dest.Value)
	assert.Equal(t, "http://www.w3.org/2001/XMLSchema#string", dest.DataType)
}
