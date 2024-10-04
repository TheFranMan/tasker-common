package types

type Retry struct {
	err error
}

func (r Retry) Error() string {
	return r.err.Error()
}

type Failure struct {
	err error
}

func (f Failure) Error() string {
	return f.err.Error()
}
