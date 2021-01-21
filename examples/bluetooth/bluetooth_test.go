package bluetooth

import (
	"fmt"
	"github.com/nobonobo/joycon"
	"log"
	"os"
	"testing"

	"github.com/ecc1/ble"
)

func TestName2(t *testing.T) {
	devices, err := joycon.Search(joycon.ProCon)
	if err != nil {
		log.Fatalln(err)
	}
	jc, err := joycon.NewJoycon(devices[0].Path, false)
	if err != nil {
		log.Fatalln(err)
	}
	s := <-jc.State()
	fmt.Println(s.Buttons)  // Button bits
	fmt.Println(s.LeftAdj)  // Left Analog Stick State
	fmt.Println(s.RightAdj) // Right Analog Stick State
	a := <-jc.Sensor()
	fmt.Println(a.Accel) // Acceleration Sensor State
	fmt.Println(a.Gyro)  // Gyro Sensor State

	jc.Close()
}

func TestName(t *testing.T) {
	conn, err := ble.Open()
	if err != nil {
		log.Fatal("open ", err)
	}
	device := ble.Device(nil)
	if len(os.Args) == 2 && !ble.ValidUUID(os.Args[1]) {
		device, err = conn.GetDeviceByName(os.Args[1])
	} else {
		uuids := os.Args[1:]
		device, err = conn.GetDevice(uuids...)
	}
	if err != nil {
		log.Fatal(err)
	}
	device.Print(os.Stdout)
}
