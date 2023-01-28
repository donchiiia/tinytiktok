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

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	//s := ServiceErr
	s := NewErrNo(999, "testErr")
	s.ErrMsg = err.Error()
	return s
}
