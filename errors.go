package httpbot

import "errors"

var (
	ErrRequestTimeOut = errors.New("request timeout")
	ErrRequestSend    = errors.New("err send request")
)
