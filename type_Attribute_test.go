package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Attribute(t *testing.T) {
	input := `<Attribute></Attribute>`

    dest := &Attribute{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Attribute")

	// Insert specific tests here
}
