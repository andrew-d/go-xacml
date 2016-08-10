package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Function(t *testing.T) {
	input := `<Function FunctionId="urn:oasis:names:tc:xacml:1.0:function:string-equal"/>`

	dest := &Function{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Function")

	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:function:string-equal", dest.FunctionId)
}
