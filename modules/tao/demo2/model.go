package demo2

import (
	"encoding/json"
	"github.com/leesper/tao"
	"time"
)

const (
	CodeUser = 1 + iota
	CodeUsers
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Age       uint8     `json:"age"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

// Serialize serializes Message into bytes.
func (u *User) Serialize() ([]byte, error) {
	return json.Marshal(u)
}

// MessageNumber returns message type number.
func (u *User) MessageNumber() int32 {
	return CodeUser
}

func (u *User) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

// DeserializeMessage deserializes bytes into Message.
func (u *User) DeserializeMessage(data []byte) (tao.Message, error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	err := json.Unmarshal(data, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

type Users []*User

func (u *Users) Serialize() ([]byte, error) {
	return json.Marshal(u)
}

// MessageNumber returns message type number.
func (u *Users) MessageNumber() int32 {
	return CodeUsers
}

func (u *Users) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

// DeserializeMessage deserializes bytes into Message.
func (u *Users) DeserializeMessage(data []byte) (tao.Message, error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	err := json.Unmarshal(data, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
