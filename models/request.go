package models

// Request struct contains info about request.
type Request struct {
	TS     int64  `json:"t"`
	Header string `json:"h"`
	Body   string `json:"b"`
	URI    string `json:"u"`
}
