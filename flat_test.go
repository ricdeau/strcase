package strcase

import (
	"testing"
)

func TestToFlat(t *testing.T) {
	cases := [][]string{
		{"testCase", "testcase"},
		{"TestCase", "testcase"},
		{"Test Case", "testcase"},
		{" Test Case", "testcase"},
		{"Test Case ", "testcase"},
		{" Test Case ", "testcase"},
		{"test", "test"},
		{"test_case", "testcase"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "manymanywords"},
		{"manyManyWords", "manymanywords"},
		{"AnyKind of_string", "anykindofstring"},
		{"numbers2and55with000", "numbers2and55with000"},
		{"JSONData", "jsondata"},
		{"userID", "userid"},
		{"AAAbbb", "aaabbb"},
		{"1A2", "1a2"},
		{"A1B", "a1b"},
		{"A1A2A3", "a1a2a3"},
		{"A1 A2 A3", "a1a2a3"},
		{"AB1AB2AB3", "ab1ab2ab3"},
		{"AB1 AB2 AB3", "ab1ab2ab3"},
		{"some string", "somestring"},
		{" some string", "somestring"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToFlat(in)
		if result != out {
			t.Errorf("%q (%q != %q)", in, result, out)
		}
	}

}
