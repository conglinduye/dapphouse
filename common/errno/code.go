package errno

var (
	OK                  		= &Errno{Code: 0, Message: "OK"}
	InternalServerError 		= &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             		= &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation 				= &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   				= &Errno{Code: 20002, Message: "Database error."}
	ErrToken      				= &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrUserAccountLogin			= &Errno{Code: 20004, Message: "User account or password error."}
	ErrUserAccountExisted		= &Errno{Code: 20005, Message: "User account has been existed."}
	ErrUserAccountPassword		= &Errno{Code: 20006, Message: "User account or password error."}
	ErrUserAccountNotActivate	= &Errno{Code: 20007, Message: "User account not activate."}

	// user errors
	ErrEncrypt           		= &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      		= &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid     		= &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect 		= &Errno{Code: 20104, Message: "The password was incorrect."}
)
