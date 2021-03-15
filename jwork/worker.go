package jwork

import "io"

// Worker wraps json worker
type Worker interface {
	Work()
	Resp() io.Reader
	Err() error
}
