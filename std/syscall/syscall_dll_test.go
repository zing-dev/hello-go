package syscall

import (
	"fmt"
	"syscall"
	"testing"
	"unsafe"
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300

	IDABORT    = 3
	IDCANCEL   = 2
	IDCONTINUE = 11
	IDIGNORE   = 5
	IDNO       = 7
	IDOK       = 1
	IDRETRY    = 4
	IDTRYAGAIN = 10
	IDYES      = 6
)

func TestMessageBox(t *testing.T) {
	msgboxID := InitMessageBox("Golang调用dll函数", "系统调用")
	switch msgboxID {
	case IDABORT:
		fmt.Println("用户选择了终止")
	case IDCANCEL:
		fmt.Println("用户选择了取消")
	case IDCONTINUE:
		fmt.Println("用户选择了继续")
	default:
		fmt.Println("用户选择了", msgboxID)
	}
}

func TestMessageBox2(t *testing.T) {
	msgboxID := InitMessageBox2("Golang调用dll函数", "系统调用")
	switch msgboxID {
	case IDABORT:
		fmt.Println("用户选择了终止")
	case IDCANCEL:
		fmt.Println("用户选择了取消")
	case IDCONTINUE:
		fmt.Println("用户选择了继续")
	default:
		fmt.Println("用户选择了", msgboxID)
	}
}

func InitMessageBox3(content, title string) uintptr {
	return 0
}

func InitMessageBox2(content, title string) uintptr {
	dll, err := syscall.LoadDLL("user32.dll")
	if err != nil {
		panic(err)
	}
	proc, err := dll.FindProc("MessageBoxW")
	if err != nil {
		panic(err)
	}

	//消息框样式
	var uType = uintptr(MB_OK | MB_ICONWARNING | MB_CANCELTRYCONTINUE)
	//消息框内容
	var text, _ = syscall.UTF16PtrFromString(content)
	var lpText = uintptr(unsafe.Pointer(text))
	//消息框标题
	var caption, _ = syscall.UTF16PtrFromString(title)
	var lpCaption = uintptr(unsafe.Pointer(caption))

	call, r2, err := proc.Call(0, lpText, lpCaption, uType)
	fmt.Println(r2)
	fmt.Println(err)
	return call
}

/**
动态加载dll并调用dll的函数
调用MessageBox函数
*/
func InitMessageBox(content, title string) uintptr {
	var user32, _ = syscall.LoadLibrary("user32.dll")
	defer syscall.FreeLibrary(user32)
	var box, _ = syscall.GetProcAddress(user32, "MessageBoxW")

	//消息框样式常量定义

	/*
	   int MessageBoxW(
	     HWND    hWnd,
	     LPCWSTR lpText,
	     LPCWSTR lpCaption,
	     UINT    uType
	   );
	*/

	//消息框MessageBoxA函数参数个数
	var argCount = uintptr(4)
	//消息框样式
	var uType = uintptr(MB_OK | MB_ICONWARNING | MB_CANCELTRYCONTINUE)
	//消息框内容
	var text, _ = syscall.UTF16PtrFromString(content)
	var lpText = uintptr(unsafe.Pointer(text))
	//消息框标题
	var caption, _ = syscall.UTF16PtrFromString(title)
	var lpCaption = uintptr(unsafe.Pointer(caption))

	var msgboxID, _, err = syscall.Syscall6(box, argCount, 0, lpText, lpCaption, uType, 0, 0)
	if err != 0 {
		println(err)
		return 0
	}
	return msgboxID
}
