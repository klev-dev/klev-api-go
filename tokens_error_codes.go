// Code generated by 'gen-errors'; DO NOT EDIT
package klev

const (
	ErrTokenAuthInvalid       = "token-auth-invalid"
	ErrTokenAuthInvalidFormat = "token-auth-invalid-format"
	ErrTokenAuthInvalidID     = "token-auth-invalid-id"
	ErrTokenCountLimited      = "token-count-limited"
	ErrTokenMetadataLimited   = "token-metadata-limited"
	ErrTokenNotFound          = "token-not-found"
)

func IsErrTokenAuthInvalid(err error) bool {
	return IsError(err, ErrTokenAuthInvalid)
}

func IsErrTokenAuthInvalidFormat(err error) bool {
	return IsError(err, ErrTokenAuthInvalidFormat)
}

func IsErrTokenAuthInvalidID(err error) bool {
	return IsError(err, ErrTokenAuthInvalidID)
}

func IsErrTokenCountLimited(err error) bool {
	return IsError(err, ErrTokenCountLimited)
}

func IsErrTokenMetadataLimited(err error) bool {
	return IsError(err, ErrTokenMetadataLimited)
}

func IsErrTokenNotFound(err error) bool {
	return IsError(err, ErrTokenNotFound)
}
