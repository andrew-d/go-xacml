package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Description(t *testing.T) {
	input := `<Description></Description>`

    dest := &Description{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Description")

	// Insert specific tests here
}
