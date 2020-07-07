package error

type Result struct {
	success      bool
	data         interface{}
	errorCode    string
	errorMessage string
}
