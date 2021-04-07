package msg

import (
	"context"
	"github.com/golang/protobuf/proto"
	"log"
	"test-tao/tao"
)

type TempRequest struct {
	Request *ZoneTempNotify
	Device  *SetDeviceRequest
}

func (t *TempRequest) Handle(ctx context.Context, closer tao.WriteCloser) {
	log.Println("new message...")
	content := tao.MessageFromContext(ctx)
	notify := content.(*TempRequest)
	log.Println("==>   ", len(notify.Request.Zones))
}

func (t *TempRequest) Unmarshaler(data []byte) (tao.Message, error) {
	err := proto.Unmarshal(data, t.Request)
	return t, err
}

// Serialize serializes Message into bytes.
func (t *TempRequest) Serialize() ([]byte, error) {
	return []byte{}, nil
}

// MessageNumber returns message type number.
func (t *TempRequest) MessageNumber() byte {
	return byte(MsgID_ZoneTempNotifyID)
}
