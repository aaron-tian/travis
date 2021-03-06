//nolint
package auth

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/errors"
)

var (
	errInvalidSignature  = fmt.Errorf("Invalid Signature")   //move auth
	errTooManySignatures = fmt.Errorf("Too many signatures") //move auth

	unauthorized = errors.CodeTypeUnauthorized
)

func ErrTooManySignatures() errors.TMError {
	return errors.WithCode(errTooManySignatures, unauthorized)
}
func IsTooManySignaturesErr(err error) bool {
	return errors.IsSameError(errTooManySignatures, err)
}

func ErrInvalidSignature() errors.TMError {
	return errors.WithCode(errInvalidSignature, unauthorized)
}
func IsInvalidSignatureErr(err error) bool {
	return errors.IsSameError(errInvalidSignature, err)
}
