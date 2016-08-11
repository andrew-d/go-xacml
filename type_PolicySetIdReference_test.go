package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicySetIdReference(t *testing.T) {
	input := `
	<PolicySetIdReference
		Version="1.0"
		EarliestVersion="0.1"
		LatestVersion="2.0"
		>urn:oasis:names:tc:xacml:2.0:conformance-test:IIIG300:policyset</PolicySetIdReference>`

	dest := &PolicySetIdReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicySetIdReference")

	assert.Equal(t, "1.0", dest.Version)
	assert.Equal(t, "0.1", dest.EarliestVersion)
	assert.Equal(t, "2.0", dest.LatestVersion)

	assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IIIG300:policyset", dest.Value)
}
