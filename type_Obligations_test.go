package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Obligations(t *testing.T) {
	input := `<Obligations></Obligations>`

	dest := &Obligations{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Obligations")

	// Insert specific tests here
}
