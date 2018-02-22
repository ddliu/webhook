package receiver

type ReceiverNotExistError struct {
	msg string
}

func (e *ReceiverNotExistError) Error() string {
	return e.msg
}
