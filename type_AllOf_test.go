package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllOf(t *testing.T) {
	input := `<AllOf></AllOf>`

    dest := &AllOf{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AllOf")

	// Insert specific tests here
}
