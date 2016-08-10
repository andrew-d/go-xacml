package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Status(t *testing.T) {
	input := `
	<Status>
		<StatusCode
			Value="urn:oasis:names:tc:xacml:1.0:status:ok"/>
		<StatusMessage>foobar</StatusMessage>
		<StatusDetail>details</StatusDetail>
	</Status>`

	dest := &Status{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Status")

	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:status:ok", dest.StatusCode.Value)
	assert.Equal(t, "foobar", dest.StatusMessage)
	assert.Equal(t, "details", dest.StatusDetail.XML)
}
