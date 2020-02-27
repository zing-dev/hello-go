package tcp2

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

type Message struct {
	Status  bool
	Message string
	Data    interface{}
}

func Encode(message []byte) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) ([]byte, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	if int32(reader.Buffered()) < length+4 {
		return nil, err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return nil, err
	}
	return pack[4:], nil
}

func EncodeStr(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func DecodeStr(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
