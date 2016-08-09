package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CombinerParameter(t *testing.T) {
	input := `<CombinerParameter></CombinerParameter>`

	dest := &CombinerParameter{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element CombinerParameter")

	// Insert specific tests here
}
