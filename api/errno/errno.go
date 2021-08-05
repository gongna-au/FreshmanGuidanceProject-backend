package errno

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrQuery            = &Errno{Code: 10003, Message: "Error occurred while getting url queries."}
	ErrPathParam        = &Errno{Code: 10004, Message: "Error occurred while getting path param."}
	ErrUserInput        = &Errno{Code: 10005, Message: "Error occurred while user input "}
	ErrServerJudge      = &Errno{Code: 10006, Message: "Error occurred while server judge studentId."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrDecoding   = &Errno{Code: 20003, Message: "Base64 decoding error."}

	// auth errors
	ErrTokenInvalid     = &Errno{Code: 30001, Message: "The token was invalid."}
	ErrPermissionDenied = &Errno{Code: 30002, Message: "Permission denied."}

	// user errors
	ErrUserNotFound      = &Errno{Code: 40001, Message: "The user was not found."}
	ErrPasswordIncorrect = &Errno{Code: 40002, Message: "The password was incorrect."}

	// search errors
	ErrSearch = &Errno{Code: 50001, Message: "Search not found"}
)
