package ferrors

import "fmt"

type (
	errorString struct {
		msg      string // error message from the origin
		layer    string // layer where the error occurred (e.g. "controller", "service", "db", "validator")
		userMsg  string // message to be shown to the user
		kind     string // type of the error (e.g. "db execution")
		internal string // internal error message
	}

	FError interface {
		Message() string             // error message from the origin
		Layer() string               // layer where the error occurred (e.g. "controller", "service", "db", "validator")
		UserMsg() string             // message to be shown to the user
		Kind() string                // type of the error (e.g. "db execution")
		Internal() string            // internal error message
		SetLayer(layer string)       // set the layer where the error occurred (e.g. "controller", "service", "db", "validator")
		SetUserMsg(userMsg string)   // set the message to be shown to the user
		SetKind(kind string)         // set the type of the error (e.g. "db execution")
		SetInternal(internal string) // set the internal error message
		Error() string               // string representation of the error
	}
)

// Error returns a string representation of the error
func (e *errorString) Error() string {
	compinedString := fmt.Sprintf("{Error: '%s', Layer: '%s', UserMsg: '%s', Kind: '%s', Internal: '%s'}", e.msg, e.layer, e.userMsg, e.kind, e.internal)
	return compinedString
}

// Message returns the error message from the origin
func (e *errorString) Message() string {
	return e.msg
}

// Layer returns the layer where the error occurred (e.g. "controller", "service", "db", "validator")
func (e *errorString) Layer() string {
	return e.layer
}

// UserMsg returns the message to be shown to the user
func (e *errorString) UserMsg() string {
	return e.userMsg
}

// Kind returns the type of the error (e.g. "db execution")
func (e *errorString) Kind() string {
	return e.kind
}

// Internal returns the internal error message
func (e *errorString) Internal() string {
	return e.internal
}

// SetLayer sets the layer where the error occurred (e.g. "controller", "service", "db", "validator")
func (e *errorString) SetLayer(layer string) {
	e.layer = layer
}

// SetUserMsg sets the message to be shown to the user
func (e *errorString) SetUserMsg(userMsg string) {
	e.userMsg = userMsg
}

// SetKind sets the type of the error (e.g. "db execution")
func (e *errorString) SetKind(kind string) {
	e.kind = kind
}

// SetInternal sets the internal error message
func (e *errorString) SetInternal(internal string) {
	e.internal = internal
}

// New returns a new FError
func New(text string) FError {
	return &errorString{msg: text, layer: "db"}
}

// FromError returns a FError from a normal error
func FromError(err error) FError {
	if err == nil {
		return nil
	}
	if e, ok := err.(FError); ok {
		return e
	}
	return New(err.Error())
}
