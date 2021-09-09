package TtErrors


type TtError interface {
	Error() string
	Code() int
}
type errorString struct {
	s string
	c int
}
func New(strError string) TtError {
	return &errorString{s: strError, c: 0}
}

func NewError( strError string,code int) TtError {
	return &errorString{s: strError, c: code}
}

func (e *errorString) Error() string {
	return e.s
}
func (e *errorString) Code() int {
	return e.c
}