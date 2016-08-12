package xacml

import "encoding/xml"

// 5.30: The <AttributeSelector> element produces a bag of unnamed and
// uncategorized attribute values.  The values shall be constructed from the
// node(s) selected by applying the XPath expression given by the element's
// Path attribute to the XML content indicated by the element's Category
// attribute.  Support for the <AttributeSelector> element is OPTIONAL.
//
// See section 7.3.7 for details of <AttributeSelector> evaluation.
type AttributeSelector struct {
	XMLName xml.Name `xml:"AttributeSelector"`

	// This attribute SHALL specify the attributes category of the <Content>
	// element containing the XML from which nodes will be selected. It also
	// indicates the attributes category containing the applicable
	// ContextSelectorId attribute, if the element includes a ContextSelectorId
	// xml attribute.
	Category string `xml:",attr"`

	// This attribute refers to the attribute (by its AttributeId) in the
	// request context in the category given by the Category attribute. The
	// referenced attribute MUST have data type
	// urn:oasis:names:tc:xacml:3.0:data-type:xpathExpression, and must select
	// a single node in the <Content> element.  The XPathCategory attribute of
	// the referenced attribute MUST be equal to the Category attribute of the
	// attribute selector.
	ContextSelectorId string `xml:",attr,omitempty"`

	// This attribute SHALL contain an XPath expression to be evaluated against
	// the specified XML content.  See Section 7.3.7 for details of the XPath
	// evaluation during <AttributeSelector> processing.
	Path string `xml:",attr"`

	// The attribute specifies the datatype of the values returned from the
	// evaluation of this <AttributeSelector> element.
	DataType string `xml:",attr"`

	// This attribute governs whether the element returns “Indeterminate” or an
	// empty bag in the event the XPath expression selects no node.  See
	// Section 7.3.5.  Also see Sections 7.19.2 and 7.19.3.
	MustBePresent *bool `xml:",attr"`
}

func (a *AttributeSelector) Validate(errs *Errors) {
	if a.Category == "" {
		errs.Addf("Category not given")
	}
	if a.Path == "" {
		errs.Addf("Path not given")
	}
	if a.DataType == "" {
		errs.Addf("DataType not given")
	}
	if a.MustBePresent == nil {
		errs.Addf("MustBePresent not given")
	}
}
