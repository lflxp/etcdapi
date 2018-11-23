package models

type Key struct {
	Key       string `json:"key"`
	Recursive bool   `json:"recursive"`
}

type Value struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Ttl   int64  `json:"ttl"`
}
