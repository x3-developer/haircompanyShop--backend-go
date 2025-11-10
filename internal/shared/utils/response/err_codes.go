package response

type ErrorCode string

const (
	BadRequest       ErrorCode = "BAD_REQUEST"
	ServerError      ErrorCode = "SERVER_ERROR"
	NotUnique        ErrorCode = "NOT_UNIQUE"
	MethodNotAllowed ErrorCode = "METHOD_NOT_ALLOWED"
	NotFound         ErrorCode = "NOT_FOUND"
	NotBlank         ErrorCode = "NOT_BLANK"
	MinLength        ErrorCode = "MIN_LENGTH"
	MaxLength        ErrorCode = "MAX_LENGTH"
	Forbidden        ErrorCode = "FORBIDDEN"
)

func GetErrorCodeByTag(tag string) ErrorCode {
	switch tag {
	case "required":
		return NotBlank
	case "min":
		return MinLength
	case "max":
		return MaxLength
	default:
		return ServerError
	}
}
