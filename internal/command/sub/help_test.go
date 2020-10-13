package sub_test

import (
	"strings"
	"testing"

	"github.com/jeffs/geode/internal/command/sub"
)

// TestHelpGood checks that, on valid input, sub.Help prints output and returns
// nil.
func TestHelpGood(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"no args", []string{}},
		{"known topic", []string{"help"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sout strings.Builder
			if err := sub.Help(tc.args, &sout); err != nil {
				t.Fatalf("got %q, want nil", err)
			}

			if s := sout.String(); strings.TrimSpace(s) == "" {
				t.Fatal("got blank output, wanted help")
			}
		})
	}
}

// TestHelpBad checks that, on invalid input, sub.Help prints nothing and
// returns an error.
func TestHelpBad(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"unknown topic", []string{"nonesuch"}},
		{"too many topics", []string{"help help"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sout strings.Builder
			if err := sub.Help(tc.args, &sout); err == nil {
				t.Fatal("got nil, want error")
			}

			if s := sout.String(); s != "" {
				t.Fatalf("got %q, wanted empty", s)
			}
		})
	}
}
