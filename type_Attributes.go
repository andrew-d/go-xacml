package xacml

import "encoding/xml"

// 5.44: The <Attributes> element specifies attributes of a subject, resource,
// action, environment or another category by listing a sequence of <Attribute>
// elements associated with the category.
type Attributes struct {
	XMLName xml.Name `xml:"Attributes"`

	// This attribute indicates which attribute category the contained
	// attributes belong to. The Category attribute is used to
	// differentiate between attributes of subject, resource, action,
	// environment or other categories.
	Category string `xml:",attr"`

	// This attribute provides a unique identifier for this <Attributes>
	// element. See [XMLid] It is primarily intended to be referenced in
	// multiple requests. See [Multi].
	Id string `xml:"http://www.w3.org/XML/1998/namespace id,attr,omitempty"`

	// Specifies additional sources of attributes in free form XML document
	// format which can be referenced using <AttributeSelector> elements.
	Content *Content `xml:"Content"`

	// A sequence of attributes that apply to the category of the request.
	Attributes []Attribute `xml:"Attribute"`
}

func (a Attributes) Validate(errs *Errors) {
	if a.Category == "" {
		errs.Addf("Category not given")
	}

	if a.Content != nil {
		a.Content.Validate(errs.Sub("Content"))
	}

	for i, el := range a.Attributes {
		el.Validate(errs.SubN("Attributes", i))
	}
}
