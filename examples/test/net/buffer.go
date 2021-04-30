package net

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"log"
	"sync"
)

const (
	One byte = iota + 1
	List
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Package struct {
	locker sync.Mutex
	Cache  bytes.Buffer
	Cmd    byte
	Size   int
	Data   []byte
}

func NewPackage(cmd byte, data interface{}) *Package {
	d := make([]byte, 0)
	if data != nil {
		d, _ = json.Marshal(data)
	}
	return &Package{
		Cmd:  cmd,
		Size: len(d),
		Data: d,
	}
}

func (p *Package) Pack() []byte {
	b := bytes.Buffer{}
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(p.Size))
	b.Write(size)
	b.WriteByte(p.Cmd)
	b.Write(p.Data)
	return b.Bytes()
}

//0-3 size
//4 cmd
//5: data
func (p *Package) resolve(data []byte) error {
	if len(data) < 5 {
		return errors.New("data's length so small")
	}
	p.Size = int(binary.BigEndian.Uint32(data[:4]))
	p.Cmd = data[4]
	return nil
}

func (p *Package) read(data []byte) error {
	if len(data) < 5 {
		return errors.New("data's length so small")
	}
	p.Size = int(binary.BigEndian.Uint32(data[:4]))
	p.Cmd = data[4]
	p.Data = data[5:]
	if p.Size != len(p.Data) {
		return errors.New("size not equal data's length")
	}
	return nil
}

func (p *Package) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(p.Pack())
	log.Println("WriteTo: ", p.Cmd, p.Size, n)
	return int64(n), err
}

func (p *Package) ReadFrom(r io.Reader) (int64, error) {
	header := make([]byte, 5)
	_, err := io.ReadFull(r, header)
	if err != nil {
		return 0, err
	}
	p.Size = int(binary.BigEndian.Uint32(header[:4]))
	p.Cmd = header[4]
	data := make([]byte, p.Size)
	n, err := io.ReadFull(r, data)
	p.Data = data
	return int64(n), err
}

func (p *Package) Unpack(data []byte) error {
	p.locker.Lock()
	defer p.locker.Unlock()
	log.Println("Unpack start")
	if len(p.Data) == 0 {
		if len(data) < 5 {
			return errors.New("data's length so small")
		}
		p.Size = int(binary.BigEndian.Uint32(data[:4]))
		p.Cmd = data[4]
		p.Data = data[5:]
	} else if len(p.Data) == p.Size {
		//p.Cache.Reset()
		return nil
	} else {
		p.Data = append(p.Data, data...)
	}
	//p.Cache.Write(data)
	//if err := p.read(p.Cache.Bytes()); err != nil {
	//	return err
	//}
	//log.Println("Unpack end")
	//p.Cache.Reset()
	return errors.New("")
}
