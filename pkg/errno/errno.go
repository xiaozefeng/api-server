package errno

import "fmt"

type Err struct {
	code    int
	message string
}

func (e Err) Error() string {
	return fmt.Sprintf("code:%d, message:%s", e.code, e.message)
}

var (
	ok                  = &Err{0, "ok"}
	internalServerError = &Err{10001, "internal server error"}
	errBind             = &Err{10002, "binding the request body to the struct failed"}

	errUserNotFound = &Err{20102, "user was not found "}
)

func IsOK(err error) bool {
	t, b := err.(Err)
	return b && t.code == ok.code
}

func IsInterServerError(err error) bool {
	t, b := err.(Err)
	return b && t.code == internalServerError.code
}
func IsBindingError(err error) bool {
	t, b := err.(Err)
	return b && t.code == errBind.code
}
func IsUserNotFoundError(err error) bool {
	t, b := err.(Err)
	return b && t.code == errUserNotFound.code
}
