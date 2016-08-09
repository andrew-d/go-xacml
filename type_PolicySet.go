package xacml

import "encoding/xml"

// 5.1: The <PolicySet> element is a top-level element in the XACML policy
// schema.  <PolicySet> is an aggregation of other policy sets and policies.
// Policy sets MAY be included in an enclosing <PolicySet> element either
// directly using the <PolicySet> element or indirectly using the
// <PolicySetIdReference> element.  Policies MAY be included in an enclosing
// <PolicySet> element either directly using the <Policy> element or indirectly
// using the <PolicyIdReference> element.
//
// A <PolicySet> element may be evaluated, in which case the evaluation
// procedure defined in Section 7.13 SHALL be used.
//
// If a <PolicySet> element contains references to other policy sets or
// policies in the form of URLs, then these references MAY be resolvable.
//
// Policy sets and policies included in a <PolicySet> element MUST be combined
// using the algorithm identified by the PolicyCombiningAlgId attribute.
// <PolicySet> is treated exactly like a <Policy> in all policy-combining
// algorithms.
//
// A <PolicySet> element MAY contain a <PolicyIssuer> element. The
// interpretation of the <PolicyIssuer> element is explained in the separate
// administrative policy profile [XACMLAdmin].
//
// The <Target> element defines the applicability of the <PolicySet> element to
// a set of decision requests.  If the <Target> element within the <PolicySet>
// element matches the request context, then the <PolicySet> element MAY be
// used by the PDP in making its authorization decision.  See Section 7.13.
//
// The <ObligationExpressions> element contains a set of obligation expressions
// that MUST be evaluated into obligations by the PDP and the resulting
// obligations MUST be fulfilled by the PEP in conjunction with the
// authorization decision.  If the PEP does not understand or cannot fulfill
// any of the obligations, then it MUST act according to the PEP bias.  See
// Section 7.2 and 7.18.
//
// The <AdviceExpressions> element contains a set of advice expressions that
// MUST be evaluated into advice by the PDP. The resulting advice MAY be safely
// ignored by the PEP in conjunction with the authorization decision.  See
// Section 7.18.
type PolicySet struct {
	XMLName xml.Name `xml:"PolicySet"`

	// Policy set identifier.  It is the responsibility of the PAP to
	// ensure that no two policies visible to the PDP have the same
	// identifier.  This MAY be achieved by following a predefined URN or
	// URI scheme.  If the policy set identifier is in the form of a URL,
	// then it MAY be resolvable.
	PolicySetId string `xml:",attr"`

	// The version number of the PolicySet.
	Version string `xml:",attr"`

	// The identifier of the policy-combining algorithm by which the
	// <PolicySet>, <CombinerParameters>, <PolicyCombinerParameters> and
	// <PolicySetCombinerParameters> components MUST be combined.  Standard
	// policy-combining algorithms are listed in Appendix Appendix C.
	// Standard policy-combining algorithm identifiers are listed in
	// Section B.9.
	PolicyCombiningAlgId string `xml:",attr"`

	// If present, limits the depth of delegation which is authorized by
	// this policy set. See the delegation profile [XACMLAdmin].
	MaxDelegationDepth int `xml:",attr"`

	// A free-form description of the policy set.
	Description string `xml:"Description"`

	// Attributes of the issuer of the policy set.
	PolicyIssuer *PolicyIssuer `xml:"PolicyIssuer"`

	// A set of default values applicable to the policy set.  The scope of
	// the <PolicySetDefaults> element SHALL be the enclosing policy set.
	PolicySetDefaults *PolicySetDefaults `xml:"PolicySetDefaults"`

	// The <Target> element defines the applicability of a policy set to a
	// set of decision requests.
	//
	// The <Target> element MAY be declared by the creator of the
	// <PolicySet> or it MAY be computed from the <Target> elements of the
	// referenced <Policy> elements, either as an intersection or as a
	// union.
	Target *Target `xml:"Target"`

	// A policy set that is included in this policy set.
	PolicySets []PolicySet `xml:"PolicySet"`

	// A policy that is included in this policy set.
	Policies []Policy `xml:"Policy"`

	// A reference to a policy set that MUST be included in this policy
	// set.  If <PolicySetIdReference> is a URL, then it MAY be resolvable.
	PolicySetIdReferences []PolicySetIdReference `xml:"PolicySetIdReference"`

	// A reference to a policy that MUST be included in this policy set.
	// If the <PolicyIdReference> is a URL, then it MAY be resolvable.
	PolicyIdReferences []PolicyIdReference `xml:"PolicyIdReference"`

	// Contains the set of <ObligationExpression> elements.  See Section
	// 7.18 for a description of how the set of obligations to be returned
	// by the PDP shall be determined.
	ObligationExpressions *ObligationExpressions `xml:"ObligationExpressions"`

	// Contains the set of <AdviceExpression> elements.  See Section 7.18
	// for a description of how the set of advice to be returned by the PDP
	// shall be determined.
	AdviceExpressions *AdviceExpressions `xml:"AdviceExpressions"`

	// Contains a sequence of <CombinerParameter> elements. The parameters
	// apply to the combining algorithm as such and it is up to the
	// specific combining algorithm to interpret them and adjust its
	// behavior accordingly.
	CombinerParameters *CombinerParameters `xml:"CombinerParameters"`

	// Contains a sequence of <CombinerParameter> elements that are
	// associated with a particular <Policy> or <PolicyIdReference> element
	// within the <PolicySet>. It is up to the specific combining algorithm
	// to interpret them and adjust its behavior accordingly.
	PolicyCombinerParameters *PolicyCombinerParameters `xml:"PolicyCombinerParameters"`

	// Contains a sequence of <CombinerParameter> elements that are
	// associated with a particular <PolicySet> or <PolicySetIdReference>
	// element within the <PolicySet>. It is up to the specific combining
	// algorithm to interpret them and adjust its behavior accordingly.
	PolicySetCombinerParameters *PolicySetCombinerParameters `xml:"PolicySetCombinerParameters"`
}

func (p PolicySet) Validate(errs *Errors) {
	// Attributes
	if p.PolicySetId == "" {
		errs.Addf("PolicySetId not given")
	}
	if p.Version == "" {
		errs.Addf("Version not given")
	}
	if p.PolicyCombiningAlgId == "" {
		errs.Addf("PolicyCombiningAlgId not given")
	}
	if p.MaxDelegationDepth < 0 {
		errs.Addf("MaxDelegationDepth cannot be less than 0 (%d)", p.MaxDelegationDepth)
	}

	// Sequences
	for i, el := range p.PolicySets {
		el.Validate(errs.SubN("PolicySets", i))
	}
	for i, el := range p.Policies {
		el.Validate(errs.SubN("Policies", i))
	}
	for i, el := range p.PolicySetIdReferences {
		el.Validate(errs.SubN("PolicySetIdReferences", i))
	}
	for i, el := range p.PolicyIdReferences {
		el.Validate(errs.SubN("PolicyIdReferences", i))
	}

	// Child elements
	if p.Target != nil {
		p.Target.Validate(errs.Sub("Target"))
	}
	if p.ObligationExpressions != nil {
		p.ObligationExpressions.Validate(errs.Sub("ObligationExpressions"))
	}
	if p.AdviceExpressions != nil {
		p.AdviceExpressions.Validate(errs.Sub("AdviceExpressions"))
	}
	if p.CombinerParameters != nil {
		p.CombinerParameters.Validate(errs.Sub("CombinerParameters"))
	}
	if p.PolicyCombinerParameters != nil {
		p.PolicyCombinerParameters.Validate(errs.Sub("PolicyCombinerParameters"))
	}
	if p.PolicySetCombinerParameters != nil {
		p.PolicySetCombinerParameters.Validate(errs.Sub("PolicySetCombinerParameters"))
	}
}
