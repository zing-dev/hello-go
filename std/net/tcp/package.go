package main

type Package struct {
	Id     int
	Length int
	Data   interface{}
}

func (p Package) Pack(data []byte) {
	if len(data) < 1 {
		return
	}

}

func (p Package) Unpack() []byte {
	return nil
}
