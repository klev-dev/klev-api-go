// Code generated by 'gen-errors'; DO NOT EDIT
package klev

const (
	ErrOffsetLogIDFieldInvalidCode = "offset-log-id-field-invalid"
	ErrOffsetPathInvalidCode       = "offset-path-invalid"
)

func IsErrOffsetLogIDFieldInvalid(err error) bool {
	return IsError(err, ErrOffsetLogIDFieldInvalidCode)
}

func IsErrOffsetPathInvalid(err error) bool {
	return IsError(err, ErrOffsetPathInvalidCode)
}
