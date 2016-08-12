package main

import (
	"encoding/xml"
	"log"
	"os"

	"github.com/andrew-d/go-xacml"
)

func main() {
	log.SetOutput(os.Stderr)

	el := os.Args[1]

	var obj xacml.Validatable
	switch el {
	case "Request":
		obj = &xacml.Request{}
	case "Response":
		obj = &xacml.Response{}
	case "Policy":
		obj = &xacml.Policy{}
	case "PolicySet":
		obj = &xacml.PolicySet{}
	default:
		log.Fatalf("invalid object type %q", el)
	}

	if err := xml.NewDecoder(os.Stdin).Decode(obj); err != nil {
		log.Fatalf("error decoding input: %s", err)
	}

	errs := xacml.NewErrors(el)
	obj.Validate(errs)

	if errs.NumErrors() > 0 {
		log.Printf("error validating input element")
		log.Println(errs.Summary())
		os.Exit(1)
	}

	if err := xml.NewEncoder(os.Stdout).Encode(obj); err != nil {
		log.Fatalf("error encoding element: %s", err)
	}
}
