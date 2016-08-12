package xacml

import "encoding/xml"

// 5.28: The <Function> element SHALL be used to name a function as an argument
// to the function defined by the parent <Apply> element.
type Function struct {
	XMLName xml.Name `xml:"Function"`

	// The identifier of the function.
	FunctionId string `xml:",attr"`
}

func (f Function) Validate(errs *Errors) {
	if f.FunctionId == "" {
		errs.Addf("FunctionId not given")
	}
}
