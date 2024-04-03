package types

type Word struct {
	Id    int64
	Front string `json:"front"`
	Back  string `json:"back"`
	_     struct{}
}
