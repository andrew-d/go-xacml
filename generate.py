#!/usr/bin/env python3

from os import path

elements = [
    'Advice',
    'AdviceExpression',
    'AdviceExpressions',
    'AllOf',
    'AnyOf',
    'Apply',
    'AssociatedAdvice',
    'Attribute',
    'AttributeAssignment',
    'AttributeAssignmentExpression',
    'AttributeDesignator',
    'AttributeSelector',
    'AttributeValue',
    'Attributes',
    'AttributesReference',
    'CombinerParameter',
    'CombinerParameters',
    'Condition',
    'Content',
    'Decision',
    'Description',
    'Expression',
    'Function',
    'Match',
    'MissingAttributeDetail',
    'MultiRequests',
    'Obligation',
    'ObligationExpression',
    'ObligationExpressions',
    'Obligations',
    'Policy',
    'PolicyCombinerParameters',
    'PolicyDefaults',
    'PolicyIdReference',
    'PolicyIdentifierList',
    'PolicyIssuer',
    'PolicySet',
    'PolicySetCombinerParameters',
    'PolicySetDefaults',
    'PolicySetIdReference',
    'Request',
    'RequestDefaults',
    'RequestReference',
    'Response',
    'Result',
    'Rule',
    'RuleCombinerParameters',
    'Status',
    'StatusCode',
    'StatusDetail',
    'StatusMessage',
    'Target',
    'VariableDefinition',
    'VariableReference',
    'XPathVersion',
]


struct_template = """
package xacml

import "encoding/xml"

// 5.XX: TODO DESCRIPTION HERE
type ELEMENTNAME struct {
	XMLName xml.Name `xml:"ELEMENTNAME"`

	// Insert fields here
}
"""


test_template = """
package xacml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ELEMENTNAME(t *testing.T) {
	input := `<ELEMENTNAME></ELEMENTNAME>`

    dest := &ELEMENTNAME{}
	err := xml.Unmarshal([]byte(input), dest)
	assert.NoError(t, err, "Error unmarshalling input for element ELEMENTNAME")

	// Insert specific tests here
}
"""


def generate_struct(el):
    s = struct_template.replace('ELEMENTNAME', el).strip() + '\n'
    fname = 'type_{0}.go'.format(el)

    if path.exists(fname) and path.getsize(fname) > len(s):
        print('[-] Skipping existing file: {0}'.format(fname))
        return

    print('[*] Writing definition to file: {0}'.format(fname))
    with open(fname, 'w') as f:
        f.write(s)


def generate_test(el):
    s = test_template.replace('ELEMENTNAME', el).strip() + '\n'
    fname = 'type_{0}_test.go'.format(el)

    if path.exists(fname) and path.getsize(fname) > len(s):
        print('[-] Skipping existing file: {0}'.format(fname))
        return

    print('[*] Writing tests to file: {0}'.format(fname))
    with open(fname, 'w') as f:
        f.write(s)


def main():
    for el in elements:
        generate_struct(el)
        generate_test(el)


if __name__ == "__main__":
    main()
