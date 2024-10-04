package types

type Retry struct {
	Err error
}

func (r Retry) Error() string {
	return r.Err.Error()
}

type Failure struct {
	Err error
}

func (f Failure) Error() string {
	return f.Err.Error()
}
