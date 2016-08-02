package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Function(t *testing.T) {
	input := `<Function></Function>`

    dest := &Function{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Function")

	// Insert specific tests here
}
