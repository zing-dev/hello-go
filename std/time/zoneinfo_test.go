package time

import (
	"testing"
	"time"
)

func TestLocal(t *testing.T) {
	t.Log(time.Local.String())
}

func TestLoadLocation(t *testing.T) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(location)
	time.Now().In(location)
	t.Log(time.Now().In(location))
	t.Log(time.Now())
	t.Log(time.Local.String())

	now := time.Now().In(location)
	zone, offset := now.Zone()
	t.Log(zone, offset/3600)

	now = time.Now()
	zone, offset = now.Zone()
	t.Log(zone, offset/3600)

	now = time.Now().UTC()
	zone, offset = now.Zone()
	t.Log(zone, offset/3600)
}

func TestFixedZone(t *testing.T) {
	location := time.FixedZone("UTC", 3600*8)
	t.Log(time.Now().In(location))

	location = time.FixedZone("UTC", 0)
	t.Log(time.Now().In(location))

	location = time.FixedZone("CST", 3600*8)
	t.Log(time.Now().In(location))

	location = time.FixedZone("Asia/Shanghai", 3600*8)
	t.Log(time.Now().In(location))

}
