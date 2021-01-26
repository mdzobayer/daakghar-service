package jwork

import "io"

type errWork struct {
	err error
}

func (e errWork) Work() {

}

func (e errWork) Err() error {
	return e.err
}

func (e errWork) Resp() io.Reader {
	return nil
}

// NewErr returns err Worker
func NewErr(err error) Worker {
	return &errWork{
		err: err,
	}
}
