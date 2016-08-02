package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Target(t *testing.T) {
	input := `<Target></Target>`

    dest := &Target{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Target")

	// Insert specific tests here
}
