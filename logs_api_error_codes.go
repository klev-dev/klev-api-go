// Code generated by 'gen-errors'; DO NOT EDIT
package klev

const (
	ErrLogPathInvalidCode = "log-path-invalid"
)

func IsErrLogPathInvalid(err error) bool {
	return IsError(err, ErrLogPathInvalidCode)
}
