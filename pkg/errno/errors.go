// 参照github.com\pkg\errors@v0.9.1中的WithStack函数，实现了WithCode函数，用于将Errno转换为error结构

package errno

import (
	"fmt"
	"runtime"
)

// 手动写入stack.go包下的私有函数
// stack represents a stack of program counters.
type stack []uintptr

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

type withCode struct {
	code  int
	msg   string
	cause error
	*stack
}

func WithCode(errno ErrNo) error {
	return &withCode{
		code:  errno.ErrCode,
		msg:   errno.ErrMsg,
		stack: callers(),
	}
}

func WrapC(err error, errno ErrNo) error {
	if err == nil {
		return nil
	}

	return &withCode{
		code:  errno.ErrCode,
		msg:   errno.ErrMsg,
		cause: err,
		stack: callers(),
	}
}

// Error return the externally-safe error message.
func (w *withCode) Error() string { return fmt.Sprintf("%v", w) }

// Cause return the cause of the withCode error.
func (w *withCode) Cause() error { return w.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withCode) Unwrap() error { return w.cause }
