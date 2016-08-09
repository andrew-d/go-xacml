package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicySetDefaults(t *testing.T) {
	input := `
	<PolicySetDefaults>
		<XPathVersion>http://www.w3.org/TR/1999/Rec-xpath-19991116</XPathVersion>
	</PolicySetDefaults>`

	dest := &PolicySetDefaults{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicySetDefaults")

	assert.Equal(t, "http://www.w3.org/TR/1999/Rec-xpath-19991116", dest.XPathVersion.URI)
}
