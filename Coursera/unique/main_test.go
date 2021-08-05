package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
2
3
4`

var testOkResult = `1
2
3
4
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("Test for OK Failed - error")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("Test for OK Failed - results not match\n %v %v", result, testOkResult)
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("Test for OK Failed - error")
	}
}
