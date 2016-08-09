package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdviceExpression(t *testing.T) {
	input := `<AdviceExpression></AdviceExpression>`

	dest := &AdviceExpression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AdviceExpression")

	// Insert specific tests here
}
