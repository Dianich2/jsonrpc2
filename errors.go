package jsonrpc2

// создадим для удобства константы для кодов ошибок
const (
	CodeParseError     = -32700
	CodeInvalidRequest = -32600
	CodeMethodNotFound = -32601
	CodeInvalidParams  = -32602
	CodeInternalError  = -32603
)

// напишем функции-хелперы для ошибок
// (будут возвращать ссылку на нашу структурку Error)

func ErrParseError(message string) *Error {
	return &Error{
		Code:    CodeParseError,
		Message: message,
	}
}

func ErrInvalidRequest(message string) *Error {
	return &Error{
		Code:    CodeInvalidRequest,
		Message: message,
	}
}

func ErrMethodNotFound(methodName string) *Error {
	return &Error{
		Code:    CodeInvalidRequest,
		Message: "Method not found",
		Data:    methodName,
	}
}

func ErrInvalidParams(message string) *Error {
	return &Error{
		Code:    CodeInvalidParams,
		Message: message,
	}
}

func ErrInternalError(message string) *Error {
	return &Error{
		Code:    CodeInternalError,
		Message: message,
	}
}
