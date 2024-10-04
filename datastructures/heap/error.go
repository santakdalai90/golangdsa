package heap

type HeapEmptyError struct {
}

func NewHeapEmptyError() *HeapEmptyError {
	return &HeapEmptyError{}
}

func (e *HeapEmptyError) Error() string {
	return "Stack is empty"
}
