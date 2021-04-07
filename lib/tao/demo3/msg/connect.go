package msg

type ConnectMsg struct {
}

// Serialize serializes Message into bytes.
func (t *ConnectMsg) Serialize() ([]byte, error) {
	return []byte{0}, nil
}

// MessageNumber returns message type number.
func (t *ConnectMsg) MessageNumber() byte {
	return byte(MsgID_ConnectID)
}
