package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequestReference(t *testing.T) {
	input := `
	<RequestReference>
		<AttributesReference ReferenceId="subject1"/>
		<AttributesReference ReferenceId="resource"/>
		<AttributesReference ReferenceId="action"/>
		<AttributesReference ReferenceId="environment"/>
	</RequestReference>`

	dest := &RequestReference{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element RequestReference")

	// Insert specific tests here
	if assert.Len(t, dest.AttributesReferences, 4) {
		assert.Equal(t, "subject1", dest.AttributesReferences[0].ReferenceId)
		assert.Equal(t, "resource", dest.AttributesReferences[1].ReferenceId)
		assert.Equal(t, "action", dest.AttributesReferences[2].ReferenceId)
		assert.Equal(t, "environment", dest.AttributesReferences[3].ReferenceId)
	}
}
