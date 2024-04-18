package types

type Word struct {
	Id         int64
	Front      string `json:"front"`
	Back       string `json:"back"`
	Reversable bool   `json:"reversable"`
	Info       string `json:"info"`
	_          struct{}
}

type Deck struct {
	Id    int64
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Words []Word `json:"words"`
	_     struct{}
}
