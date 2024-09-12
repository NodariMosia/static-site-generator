package htmlnodes

import "errors"

var (
	ErrUnimplementedMethodCall      = errors.New("unimplemented method call")
	ErrEmptyTag                     = errors.New("empty tag is not allowed")
	ErrEmptyValue                   = errors.New("empty value is not allowed")
	ErrEmptyChildren                = errors.New("empty children is not allowed")
	ErrToHTMLCalledOnNilNodePointer = errors.New("ToHTML called on nil node pointer")
)
