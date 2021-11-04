package controller

// ResponseController set the default type of a simple object response.
// All responses objects should implement this interface.
//
// Example:
//	func (e getUserExams) Execute(pl interface{}) (result controller.ResponseController) {
//	}
type ResponseController interface {
	SetError(errs error)
	GetError() error

	SetResponse(response interface{})
	GetResponse() interface{}
}

type baseResponseController struct {
	err      error
	response interface{}
}

func (b baseResponseController) GetError() error {
	return b.err
}

func (b baseResponseController) GetResponse() interface{} {
	return b.response
}

func (b *baseResponseController) SetError(err error) {
	b.err = err
}

func (b *baseResponseController) SetResponse(response interface{}) {
	b.response = response
}

func NewResponseController() ResponseController {
	return &baseResponseController{}
}
