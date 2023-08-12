package types

type Message struct {
	Payload string `json:"payload"`
}

type Health struct {
	Status string `json:"status"`
}
