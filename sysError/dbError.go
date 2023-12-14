package sysError

type DBError struct {
	Message string
}

func (e *DBError) Error() string {
	return e.Message
}
