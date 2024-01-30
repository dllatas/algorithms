package stack

import (
	"testing"
)

func TestMultipleBalance(t *testing.T) {

	var balancedTests = []struct {
		desc  string
		input string
		specs map[string]string
		want  bool
	}{
		{
			desc:  "simple case with a single parenthesis type",
			input: "()",
			want:  true,
			specs: map[string]string{
				"(": ")",
				"{": "}",
				"[": "]",
			},
		},
		{
			desc:  "using 3 diff parenthesis embedded",
			input: "({[]})",
			want:  true,
			specs: map[string]string{
				"(": ")",
				"{": "}",
				"[": "]",
			},
		},
		{
			desc:  "using 3 diff parenthesis in sequence",
			input: "(){}[]",
			want:  true,
			specs: map[string]string{
				"(": ")",
				"{": "}",
				"[": "]",
			},
		},
		{
			desc:  "starting with closing parenthesis",
			input: "){}[]",
			want:  false,
			specs: map[string]string{
				"(": ")",
				"{": "}",
				"[": "]",
			},
		},
		{
			desc:  "closing parenthesis between two balanced sets",
			input: "{})[]",
			want:  false,
			specs: map[string]string{
				"(": ")",
				"{": "}",
				"[": "]",
			},
		},
		{
			desc:  "using 3 diff parenthesis in sequence but spec does not support one",
			input: "(){}[]",
			want:  true,
			specs: map[string]string{
				"(": ")",
				"{": "}",
			},
		},
	}

	for _, test := range balancedTests {
		var got = mutipleBalance(test.input, test.specs)

		if got != test.want {
			t.Errorf("(%s) got: %v, want: %v", test.desc, got, test.want)
		}
	}
}
