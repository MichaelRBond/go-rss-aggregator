package apiResponses

// OkMsg returns a a 200 status code and string message as data
func OkMsg(message string) *APIResponseDefault {
	return &APIResponseDefault{
		Meta: buildAPIResponseMeta(200),
		Data: APIResponseMessage{message},
	}
}

// ErrorInternalServer returns a 500 status code and Error information
func ErrorInternalServer() *APIResponseDefaultError {
	return &APIResponseDefaultError{
		Meta: buildAPIResponseMetaWithError(500, "InternalServerError", "Internal Server Error"),
		Data: APIResponseMessage{"Internal Server Error"},
	}
}