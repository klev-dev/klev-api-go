// Code generated by 'gen-errors'; DO NOT EDIT
package klev

import (
	"fmt"
)

func ErrIngressWebhookTypeInvalid(s string, valid string) error {
	return &APIError{
		Code:    "ingress-webhook-type-invalid",
		Message: fmt.Sprintf("'%s' is not a valid ingress webhook type. Must be one of '%s'", s, valid),
	}
}
