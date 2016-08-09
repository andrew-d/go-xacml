package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Policy(t *testing.T) {
	input := `
	<Policy PolicyId="test-policy"
			RuleCombiningAlgId="urn:oasis:names:tc:xacml:1.0:rule-combining-algorithm:first-applicable"
			Version="1.0">
		<Description>test policy</Description>
		<Target></Target>
		<Rule Effect="Permit" RuleId="test-rule">
			<Description>test rule</Description>
			<Target>
				<AnyOf>
					<AllOf>
						<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-regexp-match">
							<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">
								http://localhost:8280/services/echo/
							</AttributeValue>
							<AttributeDesignator
								AttributeId="urn:oasis:names:tc:xacml:1.0:resource:resource-id"
								Category="urn:oasis:names:tc:xacml:3.0:attribute-category:resource"
								DataType="http://www.w3.org/2001/XMLSchema#string"
								MustBePresent="true"/>
						</Match>
						<Match MatchId="urn:oasis:names:tc:xacml:1.0:function:string-equal">
							<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">
								read
							</AttributeValue>
							<AttributeDesignator
								AttributeId="urn:oasis:names:tc:xacml:1.0:action:action-id"
								Category="urn:oasis:names:tc:xacml:3.0:attribute-category:action"
								DataType="http://www.w3.org/2001/XMLSchema#string"
								MustBePresent="true"/>
						</Match>
					</AllOf>
				</AnyOf>
			</Target>
			<Condition>
				<Apply FunctionId="urn:oasis:names:tc:xacml:1.0:function:string-subset">
					<Apply FunctionId="urn:oasis:names:tc:xacml:1.0:function:string-bag">
						<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">
							admin
						</AttributeValue>
					</Apply>
					<AttributeDesignator
						AttributeId="group"
						Category="urn:oasis:names:tc:xacml:3.0:group"
						DataType="http://www.w3.org/2001/XMLSchema#string"
						MustBePresent="true"/>
				</Apply>
			</Condition>
		</Rule>
		<Rule Effect="Deny" RuleId="deny-rule"></Rule>
	</Policy>
	`

	dest := &Policy{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Policy")

	assert.Equal(t, "test-policy", dest.PolicyId)
	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:rule-combining-algorithm:first-applicable", dest.RuleCombiningAlgId)
	assert.Equal(t, "test policy", dest.Description)

	if assert.Len(t, dest.Rules, 2) {
		assert.Equal(t, "test-rule", dest.Rules[0].RuleId)
		assert.Equal(t, "test rule", dest.Rules[0].Description)

		assert.Equal(t, "deny-rule", dest.Rules[1].RuleId)
	}
}
