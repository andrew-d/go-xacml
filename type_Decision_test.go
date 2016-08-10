package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Decision(t *testing.T) {
	input := `<Decision>Permit</Decision>`

	dest := &Decision{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Decision")

	assert.Equal(t, "Permit", dest.Value)

	errs := NewErrors("Decision")
	dest.Validate(errs)
	assert.Equal(t, 0, errs.NumErrors(), "Validate() returned errors:\n%s", errs.Summary())
}
