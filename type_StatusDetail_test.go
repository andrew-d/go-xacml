package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StatusDetail(t *testing.T) {
	input := `<StatusDetail></StatusDetail>`

    dest := &StatusDetail{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element StatusDetail")

	// Insert specific tests here
}
