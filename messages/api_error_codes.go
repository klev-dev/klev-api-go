// Code generated by 'make gen-errors'; DO NOT EDIT
package messages

import "github.com/klev-dev/klev-api-go/errors"

const (
	ErrMessagesPathInvalid      = "ERR_KLEV_MESSAGES_API_0001"
	ErrLenParameterInvalid      = "ERR_KLEV_MESSAGES_API_0002"
	ErrEncodingInvalid          = "ERR_KLEV_MESSAGES_API_0003"
	ErrEncodingParameterInvalid = "ERR_KLEV_MESSAGES_API_0004"
	ErrEncodingFieldInvalid     = "ERR_KLEV_MESSAGES_API_0005"
	ErrPollParameterInvalid     = "ERR_KLEV_MESSAGES_API_0006"
	ErrOffsetIDParameterInvalid = "ERR_KLEV_MESSAGES_API_0007"
	ErrOffsetParameterInvalid   = "ERR_KLEV_MESSAGES_API_0008"
	ErrOffsetLogMismatch        = "ERR_KLEV_MESSAGES_API_0009"
)

func IsErrMessagesPathInvalid(err error) bool {
	return errors.IsError(err, ErrMessagesPathInvalid)
}

func IsErrLenParameterInvalid(err error) bool {
	return errors.IsError(err, ErrLenParameterInvalid)
}

func IsErrEncodingInvalid(err error) bool {
	return errors.IsError(err, ErrEncodingInvalid)
}

func IsErrEncodingParameterInvalid(err error) bool {
	return errors.IsError(err, ErrEncodingParameterInvalid)
}

func IsErrEncodingFieldInvalid(err error) bool {
	return errors.IsError(err, ErrEncodingFieldInvalid)
}

func IsErrPollParameterInvalid(err error) bool {
	return errors.IsError(err, ErrPollParameterInvalid)
}

func IsErrOffsetIDParameterInvalid(err error) bool {
	return errors.IsError(err, ErrOffsetIDParameterInvalid)
}

func IsErrOffsetParameterInvalid(err error) bool {
	return errors.IsError(err, ErrOffsetParameterInvalid)
}

func IsErrOffsetLogMismatch(err error) bool {
	return errors.IsError(err, ErrOffsetLogMismatch)
}