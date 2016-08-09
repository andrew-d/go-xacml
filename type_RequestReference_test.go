package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequestReference(t *testing.T) {
	input := `<RequestReference></RequestReference>`

	dest := &RequestReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element RequestReference")

	// Insert specific tests here
}
