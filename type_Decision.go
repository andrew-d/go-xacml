package xacml

import "encoding/xml"

var kValidDecisions = []string{
	"Permit",
	"Deny",
	"Indeterminate",
	"NotApplicable",
}

// 5.53: The <Decision> element contains the result of policy evaluation.
type Decision struct {
	XMLName xml.Name `xml:"Decision"`

	// The decision
	Value string `xml:",chardata"`
}

func (d Decision) Validate(errs *Errors) {
	if d.Value == "" {
		errs.Addf("Decision not given")
		return
	}

	valid := false
	for _, opt := range kValidDecisions {
		if opt == d.Value {
			valid = true
			break
		}
	}

	if !valid {
		errs.Addf("Invalid value for Decision: %s", d.Value)
	}
}
