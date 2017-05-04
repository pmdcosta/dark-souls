package saves

const (
	ErrDirNotFound = Error("directory not found")
	ErrDirFailed   = Error("failed to create directory")
)

// Error represents an error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
