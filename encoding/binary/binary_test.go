package binary_test

import (
	"encoding/binary"
	"fmt"
	"testing"
)

//大端模式，是指数据的高字节保存在内存的低地址中，而数据的低字节保存在内存的高地址中，
// 这样的存储模式有点儿类似于把数据当作字符串顺序处理：地址由小向大增加，而数据从高位往低位放；这和我们的阅读习惯一致。

//小端模式，是指数据的高字节保存在内存的高地址中，而数据的低字节保存在内存的低地址中，
// 这种存储模式将地址的高低和数据位权有效地结合起来，高地址部分权值高，低地址部分权值低。
func TestTest(t *testing.T) {
	var num uint64
	num = 0x1234
	fmt.Printf("num = %x\n", num)

	enc := make([]byte, 8)

	// 转化为大端
	binary.BigEndian.PutUint64(enc, num)
	fmt.Printf("bigendian enc = %x\n", enc)

	// 转化为小端
	binary.LittleEndian.PutUint64(enc, num)
	fmt.Printf("littleendian enc = %x\n", enc)
}
