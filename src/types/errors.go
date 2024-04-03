package types

type ValidationErr struct {
	Msg string
}

func (err ValidationErr) Error() string {
	return err.Msg
}

type NotFoundErr struct {
	Msg string
}

func (err NotFoundErr) Error() string {
	return err.Msg
}
