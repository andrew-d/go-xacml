package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MissingAttributeDetail(t *testing.T) {
	input := `<MissingAttributeDetail></MissingAttributeDetail>`

    dest := &MissingAttributeDetail{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element MissingAttributeDetail")

	// Insert specific tests here
}
