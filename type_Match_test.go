package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Match(t *testing.T) {
	input := `<Match></Match>`

    dest := &Match{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Match")

	// Insert specific tests here
}
