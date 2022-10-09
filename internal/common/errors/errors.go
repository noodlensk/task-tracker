package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeNotFound       = ErrorType{"not-found"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
)

type SlugError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewSlugError(err, slug string) SlugError {
	return SlugError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeUnknown,
	}
}

func NewAuthorizationError(err, slug string) SlugError {
	return SlugError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeAuthorization,
	}
}

func NewIncorrectInputError(err, slug string) SlugError {
	return SlugError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeIncorrectInput,
	}
}

func NewNotFoundError(err, slug string) SlugError {
	return SlugError{
		error:     err,
		slug:      slug,
		errorType: ErrorTypeNotFound,
	}
}
