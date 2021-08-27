package errno

import "fmt"

type err struct {
	code    int
	message string
}

func (e err) Error() string {
	return fmt.Sprintf("code:%d, message:%s", e.code, e.message)
}

var (
	ok                  = &err{0, "ok"}
	internalServerError = &err{10001, "internal server error"}
	errBind             = &err{10002, "binding the request body to the struct failed"}

	errUserNotFound = &err{20102, "user was not found "}
)

func IsOK(e error) bool {
	t, b := e.(err)
	return b && t.code == ok.code
}

func IsInterServerError(e error) bool {
	t, b := e.(err)
	return b && t.code == internalServerError.code
}
func IsBindingError(e error) bool {
	t, b := e.(err)
	return b && t.code == errBind.code
}
func IsUserNotFoundError(e error) bool {
	t, b := e.(err)
	return b && t.code == errUserNotFound.code
}
