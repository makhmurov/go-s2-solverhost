package verify

import "encoding/json"

type Request struct {
	UserName string  `json:"user_name"`
	Task     string  `json:"task"`
	Results  Results `json:"results"`
}

/*
	Payload - slice of argument lists
	Results - slice of results (may be lists)
*/
type Results struct {
	Payload []json.RawMessage `json:"payload"`
	Results []interface{}     `json:"results"`
}

type Response struct {
	Percent int    `json:"percent"`
	Fails   []Fail `json:"fails"`
}

type Fail struct {
	OriginalResult string
	ExternalResult string
	DataSet        int
}
