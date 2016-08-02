package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ObligationExpressions(t *testing.T) {
	input := `<ObligationExpressions></ObligationExpressions>`

    dest := &ObligationExpressions{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element ObligationExpressions")

	// Insert specific tests here
}
