// Code generated by 'gen-errors'; DO NOT EDIT
package klev

import (
	"fmt"
)

func ErrACLActionInvalid(action string, valid string) error {
	return &APIError{
		Code:    "acl-action-invalid",
		Message: fmt.Sprintf("'%s' is not a valid action. Must be one of '%s'", action, valid),
	}
}

func ErrACLObjectMissing() error {
	return &APIError{
		Code:    "acl-object-missing",
		Message: "ACL object is missing",
	}
}

func ErrACLSubjectInvalid(subject string, valid string) error {
	return &APIError{
		Code:    "acl-subject-invalid",
		Message: fmt.Sprintf("'%s' is not a valid ACL subject. Must be one of '%s'", subject, valid),
	}
}
