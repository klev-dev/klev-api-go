// Code generated by 'gen-errors'; DO NOT EDIT
package klev

const (
	ErrLogCompactExpireAgeInvalidCode = "log-compact-expire-age-invalid"
	ErrLogCompactNotAllowedCode       = "log-compact-not-allowed"
	ErrLogCountLimitedCode            = "log-count-limited"
	ErrLogMetadataLimitedCode         = "log-metadata-limited"
	ErrLogNotFoundCode                = "log-not-found"
	ErrLogTrimAgeInvalidMaximumCode   = "log-trim-age-invalid-maximum"
	ErrLogTrimAgeInvalidMinimumCode   = "log-trim-age-invalid-minimum"
)

func IsErrLogCompactExpireAgeInvalid(err error) bool {
	return IsError(err, ErrLogCompactExpireAgeInvalidCode)
}

func IsErrLogCompactNotAllowed(err error) bool {
	return IsError(err, ErrLogCompactNotAllowedCode)
}

func IsErrLogCountLimited(err error) bool {
	return IsError(err, ErrLogCountLimitedCode)
}

func IsErrLogMetadataLimited(err error) bool {
	return IsError(err, ErrLogMetadataLimitedCode)
}

func IsErrLogNotFound(err error) bool {
	return IsError(err, ErrLogNotFoundCode)
}

func IsErrLogTrimAgeInvalidMaximum(err error) bool {
	return IsError(err, ErrLogTrimAgeInvalidMaximumCode)
}

func IsErrLogTrimAgeInvalidMinimum(err error) bool {
	return IsError(err, ErrLogTrimAgeInvalidMinimumCode)
}
