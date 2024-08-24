package jwt

import (
	"errors"
	"fmt"
)

var (
	ErrJwt                   = errors.New("JWT")
	ErrJwtWrongSigningMethod = fmt.Errorf("%w: wrong signing method", ErrJwt)
	ErrJwtParse              = fmt.Errorf("%w: cannot parse", ErrJwt)
	ErrJwtInvalid            = fmt.Errorf("%w: invalid token", ErrJwt)
	ErrJwtInvalidClaims      = fmt.Errorf("%w: invalid claims", ErrJwt)
)
