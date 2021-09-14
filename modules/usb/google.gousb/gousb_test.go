package google_gousb_test

import (
	"flag"
	"fmt"
	"log"
	"testing"

	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)

var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
)

//Open设备用每个枚举的设备调用开启器。 如果打开器返回true，则打开该设备，如果操作成功，则返回一个设备。
//必须关闭返回的每个设备（无论是否返回错误）。 如果在枚举设备时有任何错误，最后一个与任何成功打开的设备一起返回。
func TestOpenDevices(t *testing.T) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devices, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		//desc.Bus 检测到该设备的总线
		//desc.Address 总线上的设备的地址
		//desc.Port 检测到该设备的usb端口
		fmt.Printf("%03d.%03d %03d %s:%s %s\n", desc.Bus, desc.Address, desc.Port, desc.Vendor, desc.Product, usbid.Describe(desc))
		return false
	})
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("len: %d", len(devices))
}

func TestExample(t *testing.T) {
	flag.Parse()

	// Only one context should be needed for an application.  It should always be closed.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	ctx.Debug(*debug)

	// OpenDevices is used to find the devices to open.
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// The usbid package can be used to print out human readable information.
		fmt.Printf("%03d.%03d %s:%s %s\n", desc.Bus, desc.Address, desc.Vendor, desc.Product, usbid.Describe(desc))
		fmt.Printf("  Protocol: %s\n", usbid.Classify(desc))

		// The configurations can be examined from the DeviceDesc, though they can only
		// be set once the device is opened.  All configuration references must be closed,
		// to free up the memory in libusb.
		for _, cfg := range desc.Configs {
			// This loop just uses more of the built-in and usbid pretty printing to list
			// the USB devices.
			fmt.Printf("  %s:\n", cfg)
			for _, intf := range cfg.Interfaces {
				fmt.Printf("    --------------\n")
				for _, ifSetting := range intf.AltSettings {
					fmt.Printf("    %s\n", ifSetting)
					fmt.Printf("      %s\n", usbid.Classify(ifSetting))
					for _, end := range ifSetting.Endpoints {
						fmt.Printf("      %s\n", end)
					}
				}
			}
			fmt.Printf("    --------------\n")
		}

		// After inspecting the descriptor, return true or false depending on whether
		// the device is "interesting" or not.  Any descriptor for which true is returned
		// opens a Device which is retuned in a slice (and must be subsequently closed).
		return false
	})

	// All Devices returned from OpenDevices must be closed.
	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	// OpenDevices can occasionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
	}

	for _, dev := range devs {
		// Once the device has been selected from OpenDevices, it is opened
		// and can be interacted with.
		_ = dev
	}
}
