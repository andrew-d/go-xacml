package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StatusCode(t *testing.T) {
	input := `<StatusCode></StatusCode>`

    dest := &StatusCode{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element StatusCode")

	// Insert specific tests here
}
