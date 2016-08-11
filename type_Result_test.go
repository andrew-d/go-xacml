package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Result(t *testing.T) {
	input := `
	<Result>
		<Decision>Deny</Decision>
		<Status>
			<StatusCode Value="urn:oasis:names:tc:xacml:1.0:status:ok"/>
		</Status>

		<Obligations>
			<Obligation ObligationId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:obligation-1">
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1"
					DataType="http://www.w3.org/2001/XMLSchema#string">assignment1</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicSingleValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">J. Hibbert</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicMultiValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">C. Everet Koop</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicMultiValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">Victor Frankenstein</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicMultiValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">John Jeckel</AttributeAssignment>
			</Obligation>
		</Obligations>
		<AssociatedAdvice>
			<Advice AdviceId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:Advice-1">
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:assignment1"
					DataType="http://www.w3.org/2001/XMLSchema#string">assignment1</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicSingleValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">J. Hibbert</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicMultiValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">C. Everet Koop</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicMultiValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">Victor Frankenstein</AttributeAssignment>
				<AttributeAssignment
					AttributeId="urn:oasis:names:tc:xacml:2.0:conformance-test:IID302:dynamicMultiValue"
					DataType="http://www.w3.org/2001/XMLSchema#string">John Jeckel</AttributeAssignment>
			</Advice>
		</AssociatedAdvice>
		<PolicyIdentifierList>
			<PolicyIdReference Version="1.0">policy1</PolicyIdReference>
			<PolicySetIdReference Version="1.0">policyset1</PolicySetIdReference>
		</PolicyIdentifierList>
		<Attributes Category="attributes-category">
			<Attribute IncludeInResult="true" AttributeId="attribute1">
				<AttributeValue DataType="http://www.w3.org/2001/XMLSchema#string">value1</AttributeValue>
			</Attribute>
		</Attributes>
	</Result>`

	dest := &Result{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element Result")

	errs := NewErrors("Result")
	dest.Validate(errs)
	assert.Equal(t, 0, errs.NumErrors(), "Validate() returned errors:\n%s", errs.Summary())

	assert.Equal(t, "Deny", dest.Decision.Value)
	assert.Equal(t, "urn:oasis:names:tc:xacml:1.0:status:ok", dest.Status.StatusCode.Value)
	assert.Len(t, dest.Obligations.Obligations, 1)
	assert.Len(t, dest.AssociatedAdvice.Advice, 1)
	assert.Len(t, dest.PolicyIdentifierList.PolicyIdReferences, 1)
	assert.Len(t, dest.PolicyIdentifierList.PolicySetIdReferences, 1)
	if assert.Len(t, dest.Attributes, 1) {
		assert.Len(t, dest.Attributes[0].Attributes, 1)
	}
}
