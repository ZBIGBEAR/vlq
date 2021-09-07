package common

import (
	"testing"
)

type testCase struct {
	input int64
	want  string
}

func TestTen2Two(t *testing.T) {
	testCases := []testCase{
		{input: 0, want: "0"},
		{input: 1, want: "1"},
		{input: 2, want: "10"},
		{input: 3, want: "11"},
		{input: 4, want: "100"},
		{input: 5, want: "101"},
		{input: 6, want: "110"},
		{input: 7, want: "111"},
		{input: 8, want: "1000"},
		{input: 9, want: "1001"},
		{input: 10, want: "1010"},
		{input: 11, want: "1011"},
		{input: 12, want: "1100"},
		{input: 13, want: "1101"},
		{input: 14, want: "1110"},
		{input: 15, want: "1111"},
		{input: 16, want: "10000"},
	}

	for _, testcase := range testCases {
		actual := Ten2Two(testcase.input)
		if testcase.want != actual {
			t.Fatalf("want:%s,actual:%s", testcase.want, actual)
		}

		actual1 := Two2Ten(testcase.want)
		if testcase.input != actual1 {
			t.Fatalf("want:%d,actual:%d", testcase.input, actual1)
		}
	}
}


type checkNumberCase struct {
	input string
	want bool
}

func TestCheckNumber(t *testing.T) {
	testCases := []checkNumberCase{
		{input: "", want: false},
		{input: "1", want: true},
		{input: "9", want: true},
		{input: "a", want: false},
		{input: "A", want: false},
		{input: "+", want: false},
		{input: "-", want: false},
		{input: "-1", want: true},
		{input: "-9", want: true},
		{input: "-a", want: false},
		{input: "-+", want: false},
		{input: "12", want: true},
		{input: "43", want: true},
		{input: "-12", want: true},
		{input: "-43", want: true},
	}

	for _, testcase := range testCases {
		actual := CheckNumber(testcase.input)
		if testcase.want != actual {
			t.Fatalf("want:%v,actual:%v", testcase.want, actual)
		}
	}
}

func TestSupplement(t *testing.T) {
	testCases := []struct{
		input string
		want string
	}{
		{input: "",want: "00000"},
		{input: "1",want: "00001"},
		{input: "11",want: "00011"},
		{input: "101",want: "00101"},
		{input: "0111",want: "00111"},
		{input: "11111",want: "11111"},
	}

	for _, testcase := range testCases {
		actual := Supplement(testcase.input, 5)
		if testcase.want != actual {
			t.Fatalf("want:%v,actual:%v", testcase.want, actual)
		}
	}
}

func TestBinary2Vlq(t *testing.T) {
	testCases := []struct{
		input string
		want string
	}{
		{input: "1",want: "000010"},
		{input: "11",want: "000110"},
		{input: "101",want: "001010"},
		{input: "0111",want: "001110"},
		{input: "11111",want: "111110000001"},
		{input: "111110",want: "111100000011"},
		{input: "1111101",want: "111010000111"},
		{input: "11100011",want: "100110001110"},
		{input: "111011111",want: "111110011101"},
		{input: "1100101111",want: "111110110010000001"},
	}

	for _, testcase := range testCases {
		actual := Binary2Vlq(true,testcase.input)
		if testcase.want != actual {
			t.Fatalf("want:%v,actual:%v", testcase.want, actual)
		}
	}
}

func TestVlq2Base64(t *testing.T) {
	testCases := []struct{
		input string
		want string
	}{
		{input: "000000",want: "A"},
		{input: "000001",want: "B"},
		{input: "000010",want: "C"},
		{input: "000011",want: "D"},
	}

	for _, testcase := range testCases {
		actual := Vlq2Base64(testcase.input)
		if testcase.want != actual {
			t.Fatalf("want:%v,actual:%v", testcase.want, actual)
		}
	}
}