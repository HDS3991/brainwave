package berr

type ErrCode int

//go:generate stringer -type=ErrCode -linecomment
const (
	ErrUnknown           ErrCode = iota - 1    // error unknown
	ErrorInvalidArgument ErrCode = iota + 1000 // error argument invalid
	ErrorNotFound                              // error not found
	ErrorPermissionDeny                        // error permission deny
	ErrorWriteOnClose                          // error write on close
)

func NewErr(code ErrCode) ErrI {
	return &Err{
		code: int(code),
		msg:  code.String(),
	}
}
