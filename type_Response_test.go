package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Response(t *testing.T) {
	input := `
	<Response>
		<Result>
			<Decision>Permit</Decision>
			<Status>
				<StatusCode Value="urn:oasis:names:tc:xacml:1.0:status:ok"/>
			</Status>
		</Result>
		<Result>
			<Decision>Deny</Decision>
			<Status>
				<StatusCode Value="urn:oasis:names:tc:xacml:1.0:status:ok"/>
			</Status>
		</Result>
	</Response>`

	dest := &Response{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Response")

	if assert.Len(t, dest.Results, 2) {
		assert.Equal(t, "Permit", dest.Results[0].Decision.Value)
		assert.Equal(t, "Deny", dest.Results[1].Decision.Value)
	}
}
