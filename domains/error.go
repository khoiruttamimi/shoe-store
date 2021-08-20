package domains

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrProductIDResource = errors.New("(ProductID) not found or empty")

	ErrProductCodeResource = errors.New("(ProductCode) not found or empty")

	ErrBrandNotFound = errors.New("brand not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrPhonePasswordNotFound = errors.New("(Phone) or (Password) empty")

	ErrInvalidCredential = errors.New("invalid credential")

	ErrForbiddenRoles = errors.New("forbidden roles")

	ErrMissingID = errors.New("missing required id")

	ErrParamID = errors.New("invalid parameter id")

	ErrAdminKey = errors.New("invalid admin key")
)
