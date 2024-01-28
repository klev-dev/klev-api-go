// Code generated by 'make gen-errors'; DO NOT EDIT
package logs

import "github.com/klev-dev/klev-api-go/errors"

const (
	ErrNotFound         = "ERR_KLEV_LOGS_0001"
	ErrMaxMetadata      = "ERR_KLEV_LOGS_0002"
	ErrAgeCompactExpire = "ERR_KLEV_LOGS_0003"
	ErrMaxCount         = "ERR_KLEV_LOGS_0004"
)

func IsErrNotFound(err error) bool {
	return errors.IsError(err, ErrNotFound)
}

func IsErrMaxMetadata(err error) bool {
	return errors.IsError(err, ErrMaxMetadata)
}

func IsErrAgeCompactExpire(err error) bool {
	return errors.IsError(err, ErrAgeCompactExpire)
}

func IsErrMaxCount(err error) bool {
	return errors.IsError(err, ErrMaxCount)
}