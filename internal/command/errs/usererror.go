// Package errs defines a type representing user error.  It is a popular
// convention for Unix commands to return 2 (from main) for usage errors, and
// other non-zero codes for other errors.
package errs

// User represents a user error, such as an unrecognized subcommand.
type User struct {
	What string
}

// Error implements the standard go error interface.
func (e User) Error() string {
	return e.What
}
