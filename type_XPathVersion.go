package xacml

import (
	"encoding/xml"
	"strings"
)

// 5.5: The <XPathVersion> element SHALL specify the version of the XPath
// specification to be used by <AttributeSelector> elements and XPath-based
// functions in the policy set or policy.
type XPathVersion struct {
	XMLName xml.Name `xml:"XPathVersion"`
	URI     string   `xml:",chardata"`
}

func (v XPathVersion) Validate(errs *Errors) {
	if v.URI == "" {
		errs.Addf("XPathVersion not given")
		return
	}

	validURIs := []string{
		"http://www.w3.org/TR/1999/REC-xpath-19991116",
		"http://www.w3.org/TR/2007/REC-xpath20-20070123",
	}

	valid := false
	for _, uri := range validURIs {
		if strings.ToLower(uri) == strings.ToLower(v.URI) {
			valid = true
			break
		}
	}

	if !valid {
		errs.Addf("Unknown URI for XPathVersion: %s", v.URI)
	}
}
