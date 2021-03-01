package basic

import (
	"testing"
)

type Struct struct {
	Byte  byte
	PByte *byte

	Str  string
	PStr *string

	Bool  bool
	PBool *bool

	Arr  [4]byte
	PArr *[4]byte

	Slice  []byte
	PSlice *[]byte

	Chan  chan byte
	PChan *chan byte
}

func TestNewStruct(t *testing.T) {
	s := Struct{}
	t.Log(s)
	s.Byte = 10
	s.PByte = &s.Byte
	t.Log(s)

	s.Str = "hello world"
	s.PStr = &s.Str
	t.Log(s.Str, *s.PStr, &s.Str, s.PStr)

	s.PBool = new(bool)
	s.Bool = *s.PBool
	t.Log(s.Bool, *s.PBool, &s.Bool, s.PBool)

	s.PArr = new([4]byte)
	s.Arr = [4]byte{}
	t.Log(s.Arr, *s.PArr, &s.Arr, s.PArr, s.Arr[0], &s.Arr[0])
	s.PArr[0] = 1
	s.Arr[0] = 1
	t.Log(s.Arr, *s.PArr, &s.Arr, s.PArr, s.Arr[0], &s.Arr[0])
	t.Log(s.Arr == *s.PArr, &s.Arr == s.PArr, s.Arr[0] == s.PArr[0], &s.Arr[0] == &s.PArr[0])

	s.Slice = s.Arr[:]
	s.Slice = *new([]byte)
	s.Slice = []byte{}
	s.PSlice = &s.Slice
	s.PSlice = new([]byte)
	t.Log("slice", s.Slice, *s.PSlice, &s.Slice, s.PSlice)
	s.Slice = make([]byte, 4)
	s.Slice[2] = 1
	//b := s.PSlice[2]// does not support indexing
	//s.PSlice = make([]byte,4) //
	t.Log("slice", s.Slice, *s.PSlice, &s.Slice, s.PSlice)
}

func TestNew(t *testing.T) {
	t.Log(new(bool))
	t.Log(*new(bool))

	t.Log(new(byte))
	t.Log(*new(byte))

	t.Log(new(float32))
	t.Log(*new(float32))

	t.Log(new([4]byte))
	t.Log(*new([4]byte))

	t.Log(new(string))
	t.Log(*new(string))

	t.Log(new(struct{}))
	t.Log(*new(struct{}))

	t.Log(new(Struct))
	t.Log(*new(Struct))
}
