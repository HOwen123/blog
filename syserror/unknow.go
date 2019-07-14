package syserror

type UnknowError struct {
	msg    string
	reason error
}

func (this UnknowError) Code() int {
	return 1000
}

func (this UnknowError) Error() string {
	if len(this.msg) == 0 {
		return "未知错误"
	}
	return this.msg
}

func (this UnknowError) ReasonError() error{
	return this.reason
}
