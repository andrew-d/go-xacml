package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ObligationExpression(t *testing.T) {
	input := `<ObligationExpression></ObligationExpression>`

	dest := &ObligationExpression{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element ObligationExpression")

	// Insert specific tests here
}
