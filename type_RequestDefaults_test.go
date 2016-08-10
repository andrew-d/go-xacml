package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequestDefaults(t *testing.T) {
	input := `
	<RequestDefaults>
		<XPathVersion>http://www.w3.org/TR/1999/Rec-xpath-19991116</XPathVersion>
	</RequestDefaults>`

	dest := &RequestDefaults{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element RequestDefaults")

	assert.Equal(t, "http://www.w3.org/TR/1999/Rec-xpath-19991116", dest.XPathVersion.URI)
}
