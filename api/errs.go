package api

// Error messages
const (
	ErrInvalidRequestBody  = "invalid request body"
	ErrInvalidBookID       = "invalid book ID"
	ErrBookNotFound        = "book not found"
	ErrDuplicateBook       = "book with this title and author already exists"
	ErrInternalServerError = "internal server error"
	ErrFetchBooks          = "failed to fetch book(s)"
)
