package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdviceExpressions(t *testing.T) {
	input := `<AdviceExpressions></AdviceExpressions>`

	dest := &AdviceExpressions{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AdviceExpressions")

	// Insert specific tests here
}
