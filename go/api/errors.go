package api

import "errors"

var (
	errInvalidRequest = errors.New("invalid_request")
	errCreatingTest   = errors.New("error_creating_test")
	errMissingID      = errors.New("missing_id")
	errTestNotFound   = errors.New("test_not_found")
)
