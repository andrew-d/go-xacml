package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CombinerParameters(t *testing.T) {
	input := `<CombinerParameters></CombinerParameters>`

    dest := &CombinerParameters{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element CombinerParameters")

	// Insert specific tests here
}
