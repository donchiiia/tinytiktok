package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}
func NewErrNo(code int, msg string) ErrNo {
	return ErrNo{ErrCode: code, ErrMsg: msg}
}

// 由 string 构建ErrNo
func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ConvertErr 将 error 转换为 Errno格式
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	// 如果已经是Errno格式，直接返回
	if errors.As(err, &Err) {
		return Err
	}
	s := UnknownErr
	s.ErrMsg = err.Error()
	return s
}

var (
	Success                 = NewErrNo(ErrSuccessCode, "Success")
	ServiceErr              = NewErrNo(ErrServiceCode, "Service is unable to start successfully")
	UnknownErr              = NewErrNo(ErrUnknownCode, "Unknown Error")
	ParamParseErr           = NewErrNo(ErrParamParseCode, "Could not parse the param")
	ParamErr                = NewErrNo(ErrParamCode, "Wrong Parameter has been given")
	DBErr                   = NewErrNo(ErrDBCode, "Database runtime error")
	LoginErr                = NewErrNo(ErrLoginCode, "Wrong username or password")
	UserNotExistErr         = NewErrNo(ErrUserNotExistCode, "User does not exists")
	UserAlreadyExistErr     = NewErrNo(ErrUserAlreadyExistCode, "User already exists")
	TokenExpiredErr         = NewErrNo(ErrTokenExpiredCode, "Token has been expired")
	SignatureInvalidErr     = NewErrNo(ErrSignatureInvalidCode, "Signature is invalid")
	TokenBadFormErr         = NewErrNo(ErrTokenBadFormCode, "The Token is in the wrong format")
	TokenValidationErr      = NewErrNo(ErrTokenValidationCode, "Token validation error")
	TokenInvalidErr         = NewErrNo(ErrTokenInvalidCode, "Token Invalid")
	VideoProcErr            = NewErrNo(ErrVideoProcCode, "Could not process video")
	VideoNotExistErr        = NewErrNo(ErrVideoNotExistCode, "User does not exists")
	TextLenLimitExceededErr = NewErrNo(ErrCommentTextCode, "Comment text length exceeds limit")
	ActionTypeErr           = NewErrNo(ErrActionTypeCode, "Action type is invalid")
)
