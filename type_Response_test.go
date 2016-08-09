package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Response(t *testing.T) {
	input := `<Response></Response>`

	dest := &Response{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Response")

	// Insert specific tests here
}
