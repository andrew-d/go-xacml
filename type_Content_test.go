package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Content(t *testing.T) {
	input := `<Content></Content>`

	dest := &Content{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Content")

	// Insert specific tests here
}
