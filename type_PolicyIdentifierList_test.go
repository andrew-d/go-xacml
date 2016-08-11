package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolicyIdentifierList(t *testing.T) {
	input := `
	<PolicyIdentifierList>
		<PolicyIdReference>policy1</PolicyIdReference>
		<PolicyIdReference>policy2</PolicyIdReference>
		<PolicySetIdReference>policyset1</PolicySetIdReference>
		<PolicySetIdReference>policyset2</PolicySetIdReference>
	</PolicyIdentifierList>`

	dest := &PolicyIdentifierList{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element PolicyIdentifierList")

	assert.Equal(t, "policy1", dest.PolicyIdReferences[0].Value)
	assert.Equal(t, "policy2", dest.PolicyIdReferences[1].Value)
	assert.Equal(t, "policyset1", dest.PolicySetIdReferences[0].Value)
	assert.Equal(t, "policyset2", dest.PolicySetIdReferences[1].Value)
}
