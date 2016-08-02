package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Attributes(t *testing.T) {
	input := `<Attributes></Attributes>`

    dest := &Attributes{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Attributes")

	// Insert specific tests here
}
