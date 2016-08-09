package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MultiRequests(t *testing.T) {
	input := `<MultiRequests></MultiRequests>`

	dest := &MultiRequests{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element MultiRequests")

	// Insert specific tests here
}
