package demo

import (
	"encoding/binary"
	"encoding/json"
)

type Connected struct {
}

// Serialize serializes Message into bytes.
func (t *Connected) Serialize() ([]byte, error) {
	id := make([]byte, 4)
	binary.BigEndian.PutUint32(id, uint32(t.MessageNumber()))
	return id, nil
}

// MessageNumber returns message type number.
func (t *Connected) MessageNumber() int32 {
	return int32(MsgID_ConnectID)
}

type DeviceRequest struct {
	*SetDeviceRequest
}

// Serialize serializes Message into bytes.
func (t *DeviceRequest) Serialize() ([]byte, error) {
	t.ZoneTempNotifyEnable = true
	t.ZoneAlarmNotifyEnable = true
	t.FiberStatusNotifyEnable = true
	t.TempSignalNotifyEnable = true
	id := make([]byte, 4)
	binary.BigEndian.PutUint32(id, uint32(t.MessageNumber()))
	body, err := json.Marshal(t)
	data := make([]byte, 0)
	data = append(id, body...)
	return data, err
}

// MessageNumber returns message type number.
func (t *DeviceRequest) MessageNumber() int32 {
	return int32(MsgID_SetDeviceReplyID)
}

type TempNotify struct {
	ZoneTempNotify
}

// Serialize serializes Message into bytes.
func (t *TempNotify) Serialize() ([]byte, error) {
	return json.Marshal(t)
}

// MessageNumber returns message type number.
func (t *TempNotify) MessageNumber() int32 {
	return int32(MsgID_ZoneTempNotifyID)
}
