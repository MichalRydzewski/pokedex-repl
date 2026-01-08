package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []string
	}{
		"blanks around":          {input: "  hello  world  ", want: []string{"hello", "world"}},
		"single word":            {input: "word", want: []string{"word"}},
		"single blank, no words": {input: " ", want: []string{}},
		"empty string":           {input: "", want: []string{}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%v: expected: %v, got: %v", name, got, tc.want)
			}
		})
	}
}
