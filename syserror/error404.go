package syserror

type Error404 struct {
	UnknowError
}

func (this Error404) Code() int{
	return 1002
}

func (this Error404) Error() string{
	return "非法访问"
}