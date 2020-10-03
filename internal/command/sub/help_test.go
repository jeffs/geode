package sub_test

import "testing"

// To test:
//
//  geode help -- returns zero
//  geode help knowncommand -- returns zero
//  geode help unknowncommand -- returns non-zero
//  geode help multiple commands -- returns non-zero
//
// All of the above should print non-empty output.

func TestHelp(t *testing.T) {
	t.Error("TBD")
}
