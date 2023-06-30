package berr

import (
	"brainwave/pkg/i18n"
	"fmt"
)

const (
	detailKey = "detail"
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
	detail := ""
	if e.detail != nil {
		detail = fmt.Sprintf("%s - %v", detail, e.detail)
	}
	if e.err != nil {
		detail = fmt.Sprintf("%s - %v", detail, e.err.Error())
	}
	if e.msg == "" {
		return detail
	}
	return i18n.GetErrMsg(e.msg, map[string]interface{}{detailKey: detail})
}

func (e Err) Msg() string {
	return i18n.GetErrMsg(e.msg, nil)
}

func (e Err) Code() int {
	return e.code
}

func (e Err) Wrap(err error) Err {
	if e.err == nil {
		e.err = err
	} else {
		e.err = fmt.Errorf("%v - %v", e.err, err.Error())
	}
	return e
}

func (e *Err) AddDetail(detail any) *Err {
	if e.detail == nil {
		e.detail = detail
		return e
	}
	e.detail = fmt.Sprintf("%v - %v", e.detail, detail)
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
		msg:  i18n.GetErrMsg(ErrUnknown.String(), map[string]interface{}{detailKey: err.Error()}),
		err:  err,
	}
}
