package errno

const (
	// success
	ErrSuccessCode int = iota + 0
	// unknown error
	ErrUnknownCode
)

// Service Error Code
const (
	// service run error
	ErrServiceCode int = iota + 10001
	// parameter parsing error
	ErrParamParseCode
	// input parameter error
	ErrParamCode
	// database service error
	ErrDBCode
)

// Authentication And Authorization Error Code
const (
	ErrTokenExpiredCode int = iota + 10101
	ErrSignatureInvalidCode
	ErrTokenBadFormCode
	ErrTokenInvalidCode
	ErrTokenValidationCode
)

// User Service Error Code
const (
	// login failed return this code
	ErrLoginCode int = iota + 10201
	ErrUserAlreadyExistCode
	ErrUserNotExistCode
)

// Feed Service Error Code
const (
	// video processing error code
	ErrVideoProcCode int = iota + 10301
	ErrVideoNotExistCode
)

// Comment Service Error Code
const (
	// comment text length error code
	ErrCommentTextCode int = iota + 10401
)

// Action Operation Type Error Code
const (
	// action type is restricted to either 1 or 2
	ErrActionTypeCode int = iota + 10501
)
