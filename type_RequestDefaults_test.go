package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequestDefaults(t *testing.T) {
	input := `<RequestDefaults></RequestDefaults>`

    dest := &RequestDefaults{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element RequestDefaults")

	// Insert specific tests here
}
