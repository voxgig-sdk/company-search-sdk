package core

type CompanySearchError struct {
	IsCompanySearchError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewCompanySearchError(code string, msg string, ctx *Context) *CompanySearchError {
	return &CompanySearchError{
		IsCompanySearchError: true,
		Sdk:              "CompanySearch",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *CompanySearchError) Error() string {
	return e.Msg
}
