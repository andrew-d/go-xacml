package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Apply(t *testing.T) {
	input := `<Apply></Apply>`

    dest := &Apply{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Apply")

	// Insert specific tests here
}
