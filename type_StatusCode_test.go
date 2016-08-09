package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StatusCode(t *testing.T) {
	input := `
	<StatusCode Value="parent">
		<StatusCode
			Value="urn:oasis:names:tc:xacml:1.0:status:ok"/>
	</StatusCode>`

	dest := &StatusCode{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element StatusCode")

	assert.Equal(t, "parent", dest.Value)
	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:status:ok", dest.StatusCode.Value)
}
