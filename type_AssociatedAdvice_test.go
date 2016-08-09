package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AssociatedAdvice(t *testing.T) {
	input := `<AssociatedAdvice></AssociatedAdvice>`

    dest := &AssociatedAdvice{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element AssociatedAdvice")

	// Insert specific tests here
}