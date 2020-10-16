package proto

type Msg struct {
	Ver       int    `json:"ver"`  // protocol version
	Operation int    `json:"op"`   // operation for request
	SeqId     string `json:"seq"`  // sequence number chosen by client
	Body      []byte `json:"body"` // binary body bytes
}

type ConnectRequest struct {
	Angle    string `json:"angle"`
	Icon     string `json:"icon"`
	Momentum string `json:"momentum"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Type     string `json:"type"`
	X        string `json:"x"`
	Y        string `json:"y"`
}
