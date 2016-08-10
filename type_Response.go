package xacml

import "encoding/xml"

// 5.47: The <Response> element is an abstraction layer used by the policy
// language.  Any proprietary system using the XACML specification MUST
// transform an XACML context <Response> element into the form of its
// authorization decision.
//
// The <Response> element encapsulates the authorization decision produced by
// the PDP.  It includes a sequence of one or more results, with one <Result>
// element per requested resource.  Multiple results MAY be returned by some
// implementations, in particular those that support the XACML Profile for
// Requests for Multiple Resources [Multi].  Support for multiple results is
// OPTIONAL.
type Response struct {
	XMLName xml.Name `xml:"Response"`

	// An authorization decision result.  See Section 5.48.
	Results []Result `xml:"Result"`
}

func (r Response) Validate(errs *Errors) {
	if len(r.Results) < 1 {
		errs.Addf("Results not given")
	} else {
		for i, el := range r.Results {
			el.Validate(errs.SubN("Results", i))
		}
	}
}
