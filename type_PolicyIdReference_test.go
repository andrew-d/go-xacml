package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyIdReference(t *testing.T) {
	input := `
	<PolicyIdReference
		Version="1.0"
		EarliestVersion="0.1"
		LatestVersion="2.0"
		>urn:oasis:names:tc:xacml:2.0:conformance-test:IIE001:policy1</PolicyIdReference>`

	dest := &PolicyIdReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyIdReference")

	assert.Equal(t, "1.0", dest.Version)
	assert.Equal(t, "0.1", dest.EarliestVersion)
	assert.Equal(t, "2.0", dest.LatestVersion)

	assert.Equal(t, "urn:oasis:names:tc:xacml:2.0:conformance-test:IIE001:policy1", dest.Value)
}
