package errors

import "errors"

var (
	ErrorAuthorIDIsRequired    = errors.New("AuthorID is required")
	ErrorTitleIsRequired       = errors.New("title is required")
	ErrorDescriptionIsRequired = errors.New("description is required")
)
