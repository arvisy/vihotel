package helper

type Response struct {
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}

// for swagger
type Message struct {
	Message string `json:"message"`
}

type MessageDetails struct {
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}
