package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributeDesignator(t *testing.T) {
	input := `<AttributeDesignator></AttributeDesignator>`

    dest := &AttributeDesignator{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributeDesignator")

	// Insert specific tests here
}
