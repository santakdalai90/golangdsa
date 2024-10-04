package stack

type StackEmptyError struct {
}

func NewStackEmptyError() *StackEmptyError {
	return &StackEmptyError{}
}

func (e *StackEmptyError) Error() string {
	return "Stack is empty"
}
