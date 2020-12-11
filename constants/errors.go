package constants

import "errors"

var (
	ErrUnsupportedRequest   = errors.New("Unsupported request")
	ErrReplyShouldBeNotNull = errors.New("reply must not be null")
)
