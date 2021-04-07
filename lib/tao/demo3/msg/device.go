package msg

import "github.com/golang/protobuf/proto"

type DeviceRequest struct {
	Request *SetDeviceRequest
}

// Serialize serializes Message into bytes.
func (t *DeviceRequest) Serialize() ([]byte, error) {
	return proto.Marshal(t.Request)
}

// MessageNumber returns message type number.
func (t *DeviceRequest) MessageNumber() byte {
	return byte(MsgID_SetDeviceReplyID)
}
