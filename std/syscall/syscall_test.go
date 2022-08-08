package syscall

import (
	"fmt"
	"github.com/zing-dev/4g-lte-sdk"
	"syscall"
	"testing"
	"unsafe"
)

/**
获取磁盘空间
*/
func GetDiskFreeSpaceExW(diskName string) {
	// 磁盘
	diskNameUtf16Ptr, _ := syscall.UTF16PtrFromString(diskName)
	// 一下参数类型需要跟API 的类型相符
	lpFreeBytesAvailable, lpTotalNumberOfBytes,
		lpTotalNumberOfFreeBytes := int64(0), int64(0), int64(0)

	// 获取方法引用
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("获取方法引用失败:")
	}
	// 释放引用
	defer syscall.FreeLibrary(kernel32)

	getDisFreeSpaceEx, err := syscall.GetProcAddress(kernel32, "GetDiskFreeSpaceExW")
	if err != nil {
		panic("失败1")
	}

	// 根据参数个数使用对象SyscallN方法, 只需要4个参数
	r, _, errno := syscall.Syscall6(uintptr(getDisFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(diskNameUtf16Ptr)), //
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)),
		0, 0)
	// 此处的errno不是error接口， 而是type Errno uintptr
	// MSDN GetDiskFreeSpaceEx function 文档说明：
	// Return value
	// 		If the function succeeds, the return value is nonzero.
	// 		If the function fails, the return value is zero (0). To get extended error information, call GetLastError.
	// 只要是0 就是错误
	if r != 0 {
		fmt.Printf("剩余空间 %d M.\n", lpFreeBytesAvailable/1024/1204)
		fmt.Printf("用户可用总空间 %d G.\n", lpTotalNumberOfBytes/1024/1204/1024)
		fmt.Printf("剩余空间2 %d M.\n", lpTotalNumberOfFreeBytes/1024/1204)
	} else {
		fmt.Println("失败2")
		panic(errno)
	}
}

func TestGetDiskFreeSpaceExW(t *testing.T) {
	GetDiskFreeSpaceExW("C:")
}

func GetDriveTypeA(lpRootPathName string) {
	lpRootPathNameUtf16Ptr, _ := syscall.UTF16PtrFromString(lpRootPathName)
	// 获取方法引用
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("获取方法引用失败:")
	}
	// 释放引用
	defer syscall.FreeLibrary(kernel32)

	getDriveTypeA, err := syscall.GetProcAddress(kernel32, "GetDriveTypeA")
	if err != nil {
		panic("失败1")
	}
	typeDrive := 0
	r1, r2, errno := syscall.Syscall(getDriveTypeA, 2, uintptr(unsafe.Pointer(lpRootPathNameUtf16Ptr)),
		uintptr(unsafe.Pointer(&typeDrive)), 0)
	if r1 != 0 {
		fmt.Printf("磁盘驱动编号为 %d \n", typeDrive)
	} else {
		fmt.Printf("errno %d\n", errno)
	}
	fmt.Printf("r2 %d\n", r2)
}

func TestGetDriveTypeA(t *testing.T) {
	GetDriveTypeA("E:")
}

func GetFileSizeEx(hFile string) {
	hFileUtf16Ptr, _ := syscall.UTF16PtrFromString(hFile)
	// 获取方法引用
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("获取方法引用失败:")
	}
	// 释放引用
	defer syscall.FreeLibrary(kernel32)
	getFileSizeEx, err := syscall.GetProcAddress(kernel32, "GetFileSizeEx")
	if err != nil {
		panic("失败1")
	}
	lpFileSize := int64(0)
	r1, r2, errno := syscall.SyscallN(getFileSizeEx, 2,
		uintptr(unsafe.Pointer(hFileUtf16Ptr)),
		uintptr(unsafe.Pointer(&lpFileSize)), 0)
	if r1 != 0 {
		fmt.Printf("文件大小为 %d \n", lpFileSize)
	} else {
		fmt.Printf("errno %d\n", errno)
	}
	fmt.Printf("r2 %d\n", r2)
}
func TestGetFileSizeEx(t *testing.T) {
	GetFileSizeEx(`syscall_test.go`)
}

func TestTLE(t *testing.T) {
	err := lte.OpenModem(8, 115200)
	if err != nil {
		t.Fatal(err)
	}
	err = lte.SendSms("test", "18630954530")
	if err != nil {
		t.Fatal(err)
	}
}

func TestName(t *testing.T) {
}
