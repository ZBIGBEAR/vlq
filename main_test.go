package main

import (
	"fmt"
	"testing"
	"vlq/decode"
	"vlq/encode"
)

type TestCase struct {
	input string
	want  string
}

func TestEncodeAndDecode(t *testing.T) {
	testCases := []*TestCase{
		&TestCase{input: "1,333,4222,5555,3322,9888", want: ""},
		&TestCase{input: "43,555,63,34,6,33,-22,33", want: ""},
		&TestCase{input: "-22,-444,222,-333,22,-12", want: ""},
		&TestCase{input: "11,43,-111,-445", want: ""},
		&TestCase{input: "4333,-111,-22222", want: ""},
	}

	for _, testcase := range testCases {
		testcase.want, _ = encode.Encode(testcase.input)
	}

	for _, testcase := range testCases {
		actual, _ := decode.Decode(testcase.want)
		if actual != testcase.input {
			t.Fatalf(fmt.Sprintf("err. actual:%s,input:%s", actual, testcase.input))
		}
	}
}
