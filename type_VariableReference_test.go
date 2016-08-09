package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VariableReference(t *testing.T) {
	input := `<VariableReference VariableId="isFooBar"/>`

	dest := &VariableReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element VariableReference")

	assert.Equal(t, "isFooBar", dest.VariableId)
}
