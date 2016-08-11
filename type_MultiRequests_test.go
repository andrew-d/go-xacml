package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MultiRequests(t *testing.T) {
	input := `
	<MultiRequests>
		<RequestReference>
			<AttributesReference ReferenceId="subject1"/>
		</RequestReference>
	</MultiRequests>`

	dest := &MultiRequests{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element MultiRequests")

	if refs := dest.RequestReferences; assert.Len(t, refs, 1) {
		if arefs := refs[0].AttributesReferences; assert.Len(t, arefs, 1) {
			assert.Equal(t, "subject1", arefs[0].ReferenceId)
		}
	}
}
