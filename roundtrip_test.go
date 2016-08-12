package xacml

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func roundTrip(ty, input string) (string, error) {
	var obj Validatable
	switch ty {
	case "Request":
		obj = &Request{}
	case "Response":
		obj = &Response{}
	case "Policy":
		obj = &Policy{}
	case "PolicySet":
		obj = &PolicySet{}
	default:
		return "", fmt.Errorf("invalid object type %q", ty)
	}

	if err := xml.Unmarshal([]byte(input), obj); err != nil {
		return "", fmt.Errorf("error decoding input: %s", err)
	}

	errs := NewErrors(ty)
	obj.Validate(errs)

	if errs.NumErrors() > 0 {
		fmt.Println("error validating input element")
		fmt.Println(errs.Summary())
		return "", fmt.Errorf("got %d errors while validating", errs.NumErrors())
	}

	ret, err := xml.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("error encoding element: %s", err)
	}

	return string(ret), nil
}

func compareXML(one, two string) (bool, error) {
	f1, err := ioutil.TempFile("", "xacml")
	if err != nil {
		return false, err
	}
	defer f1.Close()
	defer os.Remove(f1.Name())

	_, err = io.WriteString(f1, one)
	if err != nil {
		return false, err
	}

	f2, err := ioutil.TempFile("", "xacml")
	if err != nil {
		return false, err
	}
	defer f1.Close()
	defer os.Remove(f2.Name())

	_, err = io.WriteString(f2, two)
	if err != nil {
		return false, err
	}

	f1.Sync()
	f2.Sync()

	out, err := exec.Command(
		"python",
		"./scripts/xml_equal.py",
		f1.Name(),
		f2.Name(),
	).CombinedOutput()

	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(out))
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func TestRoundTrip(t *testing.T) {
	typeRe := regexp.MustCompile(`\A[A-Z]+[0-9]+(Policy|Request|Response).xml\z`)
	files, err := filepath.Glob(filepath.Join("parse-tests", "*.xml"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		// If the type doesn't match, then exit.
		matches := typeRe.FindStringSubmatch(filepath.Base(file))
		if matches == nil || matches[1] == "" {
			fmt.Printf("skipping file: %s\n", file)
			continue
		}
		ty := matches[1]

		fmt.Printf("round-tripping file: %s\n", file)
		binput, err := ioutil.ReadFile(file)
		if !assert.NoError(t, err, "error reading input file %s: %s", file, err) {
			continue
		}
		input := string(binput)

		// Special-case "<PolicySet>"
		if ty == "Policy" && strings.Contains(input, "<PolicySet") {
			ty = "PolicySet"
		}

		output, err := roundTrip(ty, input)
		if !assert.NoError(t, err, "error round-tripping input file %s: %s", file, err) {
			continue
		}

		same, err := compareXML(input, output)
		if !assert.NoError(t, err, "error comparing file %s: %s", file, err) {
			continue
		}

		assert.True(t, same, "input and output mismatch for file: %s", file)
	}
}
