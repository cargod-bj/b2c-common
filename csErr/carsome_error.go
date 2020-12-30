package csErr

type CSError interface {
	error
	// 错误code
	Code() string
	// 错误信息
	Message() string
}

type csError struct {
	CSError
	// 错误code
	code string
	// 错误信息
	message string
	// 用于打印的message
	printM string
	// 父级错误
	parent error
}

// 新建一个carsomeError
func New(code, message string) CSError {
	e := csError{code: code, message: message, printM: message}
	return &e
}

// 新建一个carsomeError
func CreateError(parent error, code, message, pmsg string) CSError {
	e := csError{code: code, message: message, parent: parent, printM: pmsg}
	return &e
}

// 获取response code
func (e *csError) Code() string {
	return e.code
}

// 获取response message
func (e *csError) Message() string {
	return e.message
}

// 返回详细的错误信息，如果有原始错误信息，也会输出原始错误信息
func (e *csError) Error() string {
	om := ""
	if e.parent != nil {
		om = " parent: " + e.parent.Error()
	}
	return "(" + e.code + ")" + e.printM + om
}

// 返回主要错误信息，不返回原始错误信息
func (e *csError) String() string {
	return e.message + " (" + e.code + ")"
}
