package berr

type ErrCode int

//go:generate stringer -type=ErrCode -linecomment
const (
	ErrorInvalidArgument ErrCode = iota + 1000 // error argument invalid
	ErrorNotFound                              // error not found
	ErrorPermissionDeny                        // error permission deny
	ErrorWriteOnClose                          // error write on close
)
