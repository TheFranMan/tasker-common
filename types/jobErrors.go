package types

type Retry struct {
	err string
}

func (r Retry) Error() string {
	return r.err
}

type Failure struct {
	err string
}

func (f Failure) Error() string {
	return f.err
}
