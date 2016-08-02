package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Request(t *testing.T) {
	input := `<Request></Request>`

    dest := &Request{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Request")

	// Insert specific tests here
}
