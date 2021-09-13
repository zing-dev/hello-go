package msg

import (
	"context"
	"github.com/golang/protobuf/proto"
	"log"
	"test-tao/tao"
)

type SignResponse struct {
	Request *TempSignalNotify
	Device  *SetDeviceRequest
}

func (t *SignResponse) Handle(ctx context.Context, closer tao.WriteCloser) {
	log.Println("new message...")
	content := tao.MessageFromContext(ctx)
	notify := content.(*SignResponse)
	log.Println("==>   ", len(notify.Request.Signal))
}

func (t *SignResponse) Unmarshaler(data []byte) (tao.Message, error) {
	err := proto.Unmarshal(data, t.Request)
	return t, err
}

// Serialize serializes Message into bytes.
func (t *SignResponse) Serialize() ([]byte, error) {
	return []byte{}, nil
}

// MessageNumber returns message type number.
func (t *SignResponse) MessageNumber() byte {
	return byte(MsgID_TempSignalNotifyID)
}
