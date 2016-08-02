package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AnyOf(t *testing.T) {
	input := `<AnyOf></AnyOf>`

    dest := &AnyOf{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AnyOf")

	// Insert specific tests here
}
