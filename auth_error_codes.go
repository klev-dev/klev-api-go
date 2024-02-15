// Code generated by 'gen-errors'; DO NOT EDIT
package klev

const (
	ErrAuthenticationFailedCode       = "authentication-failed"
	ErrAuthorizationFailedCode        = "authorization-failed"
	ErrAuthorizationHeaderInvalidCode = "authorization-header-invalid"
	ErrAuthorizationHeaderMissingCode = "authorization-header-missing"
	ErrContentTypeInvalidCode         = "content-type-invalid"
	ErrContentTypeMissingCode         = "content-type-missing"
	ErrSubscriptionMissingCode        = "subscription-missing"
)

func IsErrAuthenticationFailed(err error) bool {
	return IsError(err, ErrAuthenticationFailedCode)
}

func IsErrAuthorizationFailed(err error) bool {
	return IsError(err, ErrAuthorizationFailedCode)
}

func IsErrAuthorizationHeaderInvalid(err error) bool {
	return IsError(err, ErrAuthorizationHeaderInvalidCode)
}

func IsErrAuthorizationHeaderMissing(err error) bool {
	return IsError(err, ErrAuthorizationHeaderMissingCode)
}

func IsErrContentTypeInvalid(err error) bool {
	return IsError(err, ErrContentTypeInvalidCode)
}

func IsErrContentTypeMissing(err error) bool {
	return IsError(err, ErrContentTypeMissingCode)
}

func IsErrSubscriptionMissing(err error) bool {
	return IsError(err, ErrSubscriptionMissingCode)
}
