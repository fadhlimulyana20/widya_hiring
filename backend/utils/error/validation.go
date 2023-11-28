package error

type ValueValidationError struct {
	Err error
}

func (v *ValueValidationError) Error() string {
	return v.Err.Error()
}
