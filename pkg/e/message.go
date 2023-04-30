package e

var ErrorMessages = map[int]string{
	SUCCESS: "ok",
	ERROR:   "Something went wrong",

	ERROR_AUTH_INVALID_TOKEN:       "Invalid authentication token",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Authentication token is expired",
	ERROR_AUTH_TOKEN_MISSING:       "Authentication token is missing",
	ERROR_AUTH:                     "Authentication failed",
}

func GetErrorMessage(code int) string {
	msg, ok := ErrorMessages[code]
	if ok {
		return msg
	}

	return ErrorMessages[ERROR]
}
