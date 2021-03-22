package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"
)

func TestWriteFloat64(t *testing.T) {
	var pi = 3.141592653589793
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%x\n", buf.Bytes())
}

func TestWriteStruct(t *testing.T) {
	buf := bytes.Buffer{}
	u := User{
		Id:  0,
		Age: 0,
	}
	err := binary.Write(&buf, binary.BigEndian, u)
	if err != nil {
		t.Fatal(err)
	}
	data := buf.Bytes()
	t.Log(data)
}

func TestWriteStructJSON(t *testing.T) {
	data, _ := json.Marshal(user)
	pack := Package{
		Id:     12,
		Length: len(data),
		Data:   data,
	}
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.LittleEndian, pack.Id)
	err = binary.Write(&buf, binary.LittleEndian, pack.Length)
	err = binary.Write(&buf, binary.LittleEndian, pack.Data)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	b := buf.Bytes()
	fmt.Printf("%d\n", b[0])
	data = make([]byte, 10)
	buf2 := bytes.NewReader(buf.Bytes())
	err = binary.Read(buf2, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(len(data))
	fmt.Println(data[0])
	fmt.Println(data[1])
}

type User1 struct {
	ID        uint16   // 用户ID
	Privilege uint8    // 用户权限
	FpNum     uint8    // 用户登记的指纹数
	SecLevel  uint16   // 用户加密等级
	PIN2      uint32   // 用户编号
	Name      [8]byte  // 用户姓名
	Password  [5]byte  // 密码
	Card      [5]uint8 // 卡ID
}

func NewUser(id uint16) *User1 {
	return &User1{
		ID:       id,
		Name:     [8]byte{},
		Password: [5]byte{},
		Card:     [5]uint8{},
	}
}

func (u *User1) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, u); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (*User1, error) {
	buf := bytes.NewBuffer(b)
	obj := &User1{}
	if err := binary.Read(buf, binary.LittleEndian, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func TestUser(t *testing.T) {
	u := NewUser(6)
	b, err := u.Encode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", b)
	u1, err := Decode(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", u1)
}
