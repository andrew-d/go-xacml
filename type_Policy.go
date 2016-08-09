package xacml

import "encoding/xml"

// 5.14: The <Policy> element is the smallest entity that SHALL be presented to
// the PDP for evaluation.
//
// A <Policy> element may be evaluated, in which case the evaluation procedure
// defined in Section 7.12 SHALL be used.
//
// The main components of this element are the <Target>, <Rule>,
// <CombinerParameters>, <RuleCombinerParameters>, <ObligationExpressions> and
// <AdviceExpressions> elements and the RuleCombiningAlgId attribute.
//
// A <Policy> element MAY contain a <PolicyIssuer> element. The interpretation
// of the <PolicyIssuer> element is explained in the separate administrative
// policy profile [XACMLAdmin].
//
// The <Target> element defines the applicability of the <Policy> element to a
// set of decision requests.  If the <Target> element within the <Policy>
// element matches the request context, then the <Policy> element MAY be used
// by the PDP in making its authorization decision.  See Section 7.12.
//
// The <Policy> element includes a sequence of choices between
// <VariableDefinition> and <Rule> elements.
//
// Rules included in the <Policy> element MUST be combined by the algorithm
// specified by the RuleCombiningAlgId attribute.
//
// The <ObligationExpressions> element contains a set of obligation expressions
// that MUST be evaluated into obligations by the PDP and the resulting
// obligations MUST be fulfilled by the PEP in conjunction with the
// authorization decision.  If the PEP does not understand, or cannot fulfill,
// any of the obligations, then it MUST act according to the PEP bias.  See
// Section 7.2 and 7.18.
//
// The <AdviceExpressions> element contains a set of advice expressions that
// MUST be evaluated into advice by the PDP. The resulting advice MAY be safely
// ignored by the PEP in conjunction with the authorization decision.  See
// Section 7.18.
type Policy struct {
	XMLName xml.Name `xml:"Policy"`

	// Policy identifier.  It is the responsibility of the PAP to ensure that
	// no two policies visible to the PDP have the same identifier.  This MAY
	// be achieved by following a predefined URN or URI scheme.  If the policy
	// identifier is in the form of a URL, then it MAY be resolvable.
	PolicyId string `xml:",attr"`

	// The version number of the Policy.
	Version string `xml:",attr"`

	// The identifier of the rule-combining algorithm by which the <Policy>,
	// <CombinerParameters> and <RuleCombinerParameters> components MUST be
	// combined.  Standard rule-combining algorithms are listed in Appendix
	// Appendix C.  Standard rule-combining algorithm identifiers are listed in
	// Section B.9.
	RuleCombiningAlgId string `xml:",attr"`

	// If present, limits the depth of delegation which is authorized by this
	// policy. See the delegation profile [XACMLAdmin].
	MaxDelegationDepth int `xml:",attr"`

	// A free-form description of the policy.  See Section 5.2.
	Description *Description `xml:"Description"`

	// Attributes of the issuer of the policy.
	PolicyIssuer *PolicyIssuer `xml:"PolicyIssuer"`

	// Defines a set of default values applicable to the policy.  The scope of
	// the <PolicyDefaults> element SHALL be the enclosing policy.
	PolicyDefaults *PolicyDefaults `xml:"PolicyDefaults"`

	// The <Target> element defines the applicability of a <Policy> to a set of
	// decision requests.
	//
	// The <Target> element MAY be declared by the creator of the <Policy>
	// element, or it MAY be computed from the <Target> elements of the
	// referenced <Rule> elements either as an intersection or as a union.
	Target *Target `xml:"Target"`

	// A conjunctive sequence of obligation expressions which MUST be evaluated
	// into obligations byt the PDP. The corresponsding obligations MUST be
	// fulfilled by the PEP in conjunction with the authorization decision.
	// See Section 7.18 for a description of how the set of obligations to be
	// returned by the PDP SHALL be determined. See section 7.2 about
	// enforcement of obligations.
	ObligationExpressions *ObligationExpressions `xml:"ObligationExpressions"`

	// A conjunctive sequence of advice expressions which MUST evaluated into
	// advice by the PDP. The corresponding advice provide supplementary
	// information to the PEP in conjunction with the authorization decision.
	// See Section 7.18 for a description of how the set of advice to be
	// returned by the PDP SHALL be determined.
	AdviceExpressions *AdviceExpressions `xml:"AdviceExpressions"`

	// A sequence of parameters to be used by the rule-combining algorithm. The
	// parameters apply to the combining algorithm as such and it is up to the
	// specific combining algorithm to interpret them and adjust its behavior
	// accordingly.
	CombinerParameters []CombinerParameters `xml:"CombinerParameters"`

	// A sequence of <RuleCombinerParameter> elements that are associated with
	// a particular <Rule> element within the <Policy>.. It is up to the
	// specific combining algorithm to interpret them and adjust its behavior
	// accordingly.
	RuleCombinerParameters []RuleCombinerParameters `xml:"RuleCombinerParameters"`

	// Common function definitions that can be referenced from anywhere in a
	// rule where an expression can be found.
	VariableDefinition []VariableDefinition `xml:"VariableDefinition"`

	// A sequence of rules that MUST be combined according to the
	// RuleCombiningAlgId attribute.  Rules whose <Target> elements and
	// conditions match the decision request MUST be considered.  Rules whose
	// <Target> elements or conditions do not match the decision request SHALL
	// be ignored.
	Rules []Rule `xml:"Rule"`
}

func (p Policy) Validate(errs *Errors) {
	// Attributes
	if p.PolicyId == "" {
		errs.Addf("PolicyId not given")
	}
	if p.Version == "" {
		errs.Addf("Version not given")
	}
	if p.RuleCombiningAlgId == "" {
		errs.Addf("RuleCombiningAlgId not given")
	}
	if p.MaxDelegationDepth < 0 {
		errs.Addf("MaxDelegationDepth cannot be less than 0 (%d)", p.MaxDelegationDepth)
	}

	// Required children
	if p.Target == nil {
		errs.Addf("Target not given")
	} else {
		p.Target.Validate(errs.Sub("Target"))
	}

	// Optional children
	if p.Description != nil {
		p.Description.Validate(errs.Sub("Description"))
	}

	if p.PolicyIssuer != nil {
		p.PolicyIssuer.Validate(errs.Sub("PolicyIssuer"))
	}

	if p.PolicyDefaults != nil {
		p.PolicyDefaults.Validate(errs.Sub("PolicyDefaults"))
	}

	if p.ObligationExpressions != nil {
		p.ObligationExpressions.Validate(errs.Sub("ObligationExpressions"))
	}

	if p.AdviceExpressions != nil {
		p.AdviceExpressions.Validate(errs.Sub("AdviceExpressions"))
	}

	// Sequences
	for i, el := range p.CombinerParameters {
		el.Validate(errs.SubN("CombinerParameters", i))
	}
	for i, el := range p.RuleCombinerParameters {
		el.Validate(errs.SubN("RuleCombinerParameters", i))
	}
	for i, el := range p.VariableDefinition {
		el.Validate(errs.SubN("VariableDefinition", i))
	}
	for i, el := range p.Rules {
		el.Validate(errs.SubN("Rules", i))
	}
}
