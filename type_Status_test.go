package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Status(t *testing.T) {
	input := `<Status></Status>`

	dest := &Status{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Status")

	// Insert specific tests here
}
