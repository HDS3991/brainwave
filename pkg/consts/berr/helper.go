package berr

import (
	"errors"
	"fmt"
)

type Err struct {
	code   int
	msg    string
	detail any
	err    error
}

type ErrI interface {
	error
	Msg() string
	Code() int
	Wrap(err error) Err
}

func (e Err) Error() string {
	content := e.msg
	if e.err != nil {
		content = fmt.Sprintf("%s - [%v]", content, e.err.Error())
	}
	if e.detail != nil {
		content = fmt.Sprintf("%s - [%v]", content, e.detail)
	}
	if content == "" {
		if e.err != nil {
			return e.err.Error()
		}
		return errors.New(e.Msg()).Error()
	}
	return content
}

func (e Err) Msg() string {
	return e.msg
}

func (e Err) Code() int {
	return e.code
}

func (e Err) Wrap(err error) Err {
	if e.err == nil {
		e.err = err
	} else {
		e.err = fmt.Errorf("%v: [%v]", e.err, err.Error())
	}
	return e
}

func DecodeErr(err error) ErrI {
	if err == nil {
		return nil
	}

	switch typed := err.(type) {
	case *Err:
		return *typed
	case Err:
		return typed
	}

	return Err{
		code: int(ErrUnknown),
		msg:  ErrUnknown.String(),
		err:  err,
	}
}
