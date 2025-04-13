package structs

// FileCopierError is a custom error type that wraps errors with an operation name.
type FileCopierError struct {
	Op  string
	Err error
}

// NewFileCopierError is a Factory function (constructor)
func NewFileCopierError(op string, err error) *FileCopierError {
	return &FileCopierError{
		Op:  op,
		Err: err,
	}
}

// Error implements the error interface for FileCopierError.
func (e FileCopierError) Error() string {
	return e.Err.Error()
}

// Unwrap allows you to retrieve the original error, and use As/Is.
func (e FileCopierError) Unwrap() error {
	return e.Err
}
