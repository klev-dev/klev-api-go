// Code generated by 'make gen-errors'; DO NOT EDIT
package klev

const (
	ErrContentTypeMissing         = "ERR_KLEV_API_AUTH_0001"
	ErrContentTypeInvalid         = "ERR_KLEV_API_AUTH_0002"
	ErrAuthorizationHeaderMissing = "ERR_KLEV_API_AUTH_0003"
	ErrAuthorizationHeaderInvalid = "ERR_KLEV_API_AUTH_0004"
	ErrAuthenticationFailed       = "ERR_KLEV_API_AUTH_0005"
	ErrSubscriptionMissing        = "ERR_KLEV_API_AUTH_0006"
)

func IsErrContentTypeMissing(err error) bool {
	return IsError(err, ErrContentTypeMissing)
}

func IsErrContentTypeInvalid(err error) bool {
	return IsError(err, ErrContentTypeInvalid)
}

func IsErrAuthorizationHeaderMissing(err error) bool {
	return IsError(err, ErrAuthorizationHeaderMissing)
}

func IsErrAuthorizationHeaderInvalid(err error) bool {
	return IsError(err, ErrAuthorizationHeaderInvalid)
}

func IsErrAuthenticationFailed(err error) bool {
	return IsError(err, ErrAuthenticationFailed)
}

func IsErrSubscriptionMissing(err error) bool {
	return IsError(err, ErrSubscriptionMissing)
}