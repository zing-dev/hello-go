package main

var T = []bool{
	true, true, true, true, true, true, true, true,
	true, true, true, true, true, true, true, true,
	true, true, true, true, true, true, true, true,
	true, true, true, true, true, true, true, true,
}

type Protocol interface {
	Pack() []byte
	UnPack([]byte)
}

type protocol struct {
	length int
	msgId  MsgId
	data   []byte
}

func NewProtocol(msgId MsgId) Protocol {
	d := msgId.ToBytes()
	return &protocol{
		length: len(d),
		msgId:  msgId,
		data:   d,
	}
}

func (p *protocol) Pack() []byte {
	var pack []byte
	pack = append(pack, IntToBytes(p.length)...)
	pack = append(pack, p.msgId.ToByte())
	pack = append(pack, p.data...)
	return pack
}

func (p *protocol) UnPack([]byte) {}
