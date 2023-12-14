package sysError

type TransactionError struct {
	Message string
}

func (e *TransactionError) Error() string {
	return e.Message
}
