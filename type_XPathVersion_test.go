package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_XPathVersion(t *testing.T) {
	input := `<XPathVersion>http://www.w3.org/TR/1999/Rec-xpath-19991116</XPathVersion>`

	dest := &XPathVersion{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element XPathVersion")

	assert.Equal(t, "http://www.w3.org/TR/1999/Rec-xpath-19991116", dest.URI)

	errs := NewErrors("XPathVersion")
	dest.Validate(errs)

	assert.Equal(t, 0, errs.NumErrors(), "Validate() returned errors:\n%s", errs.Summary())
}
