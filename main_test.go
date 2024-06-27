package main

import (
	"testing"
)

func TestCountSmileys(t *testing.T) {
	testCases := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "testCase1: NoSmileys",
			args: []string{"", " ", ": ", "; "},
			want: 0,
		},
		{
			name: "testCase2: [\":)\", \";(\", \";}\", \":-D\"]",
			args: []string{":)", ";(", ";}", ":-D"},
			want: 2,
		},
		{
			name: "testCase3: [\";D\", \":-(\", \":-)\", \";~)\"]",
			args: []string{";D", ":-(", ":-)", ";~)"},
			want: 3,
		},
		{
			name: "testCase4: [\";]\", \":[\", \";*\", \":$\", \";-D\"]",
			args: []string{";]", ":[", ";*", ":$", ";-D"},
			want: 1,
		},
		{
			name: "testCase5: long text",
			args: []string{"Hello :)", "How are you? ;-)", "Good morning :~D", "Good night ;D"},
			want: 0,
		},
		{
			name: "testCase5: empty array",
			args: []string{},
			want: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := countSmileys(tc.args)
			if got != tc.want {
				t.Errorf("countSmileys(%v) = %v; want %v", tc.args, got, tc.want)
			}
		})
	}
}
