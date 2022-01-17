package domain

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}


type AlreadyExistError struct {
	Error string
}

func NewAlreadyExistError(error string) AlreadyExistError {
	return AlreadyExistError{Error: error}
}

type BadRequestError struct {
	Error string
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{Error: error}
}
