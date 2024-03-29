package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AttributesReference(t *testing.T) {
	input := `<AttributesReference ReferenceId="subject1"/>`

	dest := &AttributesReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AttributesReference")

	assert.Equal(t, "subject1", dest.ReferenceId)
}
