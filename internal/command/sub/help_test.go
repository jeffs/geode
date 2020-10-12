package sub_test

import (
	"testing"

	"github.com/jeffs/geode/internal/command/sub"
)

// TestHelpGood checks that sub.Help returns nil on valid input.
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
			if err := sub.Help(tc.args); err != nil {
				t.Fatalf("got %q, want nil", err)
			}
		})
	}
}

// TestHelpGood checks that sub.Help returns an error on invalid input.
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
			if err := sub.Help(tc.args); err == nil {
				t.Fatal("got nil, want error")
			}
		})
	}
}
