package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Result(t *testing.T) {
	input := `<Result></Result>`

	dest := &Result{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Result")

	// Insert specific tests here
}
