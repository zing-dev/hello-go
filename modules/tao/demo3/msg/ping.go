package msg

type PingRequest struct{}

// Serialize serializes Message into bytes.
func (t *PingRequest) Serialize() ([]byte, error) {
	return []byte{}, nil
}

// MessageNumber returns message type number.
func (t *PingRequest) MessageNumber() byte {
	return byte(250)
}
