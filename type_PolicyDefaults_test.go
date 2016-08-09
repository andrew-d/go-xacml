package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyDefaults(t *testing.T) {
	input := `
	<PolicyDefaults>
		<XPathVersion>http://www.w3.org/TR/1999/Rec-xpath-19991116</XPathVersion>
	</PolicyDefaults>`

	dest := &PolicyDefaults{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyDefaults")

	assert.Equal(t, "http://www.w3.org/TR/1999/Rec-xpath-19991116", dest.XPathVersion.URI)
}
