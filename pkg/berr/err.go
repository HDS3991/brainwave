package berr

type ErrCode int

//go:generate stringer -type=ErrCode
const (
	ErrUnknown           ErrCode = iota - 1
	ErrorInvalidArgument ErrCode = iota + 1000
	ErrorRecordNotFound
	ErrorPermissionDeny
)

func NewErr(code ErrCode) ErrI {
	return &Err{
		code: int(code),
		msg:  code.String(),
	}
}
