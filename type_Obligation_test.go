package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Obligation(t *testing.T) {
	input := `<Obligation></Obligation>`

	dest := &Obligation{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Obligation")

	// Insert specific tests here
}
