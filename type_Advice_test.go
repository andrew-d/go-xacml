package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Advice(t *testing.T) {
	input := `<Advice></Advice>`

    dest := &Advice{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Advice")

	// Insert specific tests here
}
