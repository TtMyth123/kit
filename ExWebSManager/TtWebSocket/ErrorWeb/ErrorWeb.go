package ErrorWeb

const (
	EC_ConnNil = 1
	EC_WriteMessage = 1000
	EC_Marshal  = 1001
)

type ErrorWebS struct {
	code int
	strErr string

	error
}

func (this ErrorWebS)Error() string {
	return this.strErr
}

func NewErrorWeb(code int,strErr string) *ErrorWebS  {
	aErrorWebS := &ErrorWebS{code:code,strErr:strErr}
	return aErrorWebS
}