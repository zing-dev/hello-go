package main

import (
	"fmt"
	"github.com/Atian-OE/DTSSDK_Golang/dtssdk"
	"github.com/Atian-OE/DTSSDK_Golang/dtssdk/model"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Temperature struct {
	Max float32   `json:"max"`
	Avg float32   `json:"avg"`
	Min float32   `json:"min"`
	At  TimeLocal `json:"at,omitempty"`
}

type AlarmInfo struct {
	Zone
	Temperature
	Location  float32                `json:"location"`
	AlarmAt   TimeLocal              `json:"alarm_at"`
	AlarmType model.DefenceAreaState `json:"alarm_type"`
}

type ZoneExtend struct {
	Warehouse string `json:"warehouse,omitempty"`
	Group     string `json:"group,omitempty"`
	Row       int    `json:"row,omitempty"`
	Column    int    `json:"column,omitempty"`
	Layer     int    `json:"layer,omitempty"`
}

type Relay struct {
	Tag      string   `json:"tag"`
	Branches []string `json:"branches"`
}

type Zone struct {
	Id        int32   `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	ChannelId int32   `json:"channel_id,omitempty"`
	Host      string  `json:"host,omitempty"`
	Start     float32 `json:"start,omitempty"`
	Finish    float32 `json:"finish,omitempty"`
	Tag       string  `json:"tag,omitempty"`
	Relays    []Relay `json:"relays,omitempty"`
	ZoneExtend
}

type N struct {
	Warehouse map[string]struct {
		Group map[string][]Zone `json:"group"`
	} `json:"warehouse"`
}

type N2 map[string]map[string][]Zone

type DTSModel struct {
	Id        int    `json:"id"`
	ChannelId int    `json:"channel_id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	SDKPort   int    `json:"sdk_port"`
	HTTPPort  int    `json:"http_port"`
}

func (d DTSModel) SDKAddress() string {
	return fmt.Sprintf("%s:%d", d.Host, d.SDKPort)
}

func (d DTSModel) HTTPAddress() string {
	return fmt.Sprintf("%s:%d", d.Host, d.HTTPPort)
}

type DTS struct {
	DeviceId string         `json:"device_id,omitempty"`
	Zones    []Zone         `json:"zones,omitempty"`
	Model    DTSModel       `json:"model,omitempty"`
	Status   int32          `json:"status,omitempty"`
	Client   *dtssdk.Client `json:"-"`
	Locker   sync.RWMutex   `json:"-"`
}

func NewDTS(m DTSModel) *DTS {
	return &DTS{Zones: []Zone{}, Model: m}
}

func (d *DTS) GetZones() []Zone {
	d.Locker.Lock()
	defer d.Locker.Unlock()
	return d.Zones
}

func (d *DTS) getZonesFromSDK() {
	for i := 1; i <= d.Model.ChannelId; i++ {
		res, err := d.Client.GetDefenceZone(i, "")
		if err != nil {
			log.Println("GetDefenceZone", err)
		} else if res.Success {
			log.Println(fmt.Sprintf("获取 %s %d %d 防区", d.Model.Host, i, len(res.Rows)))
			if len(res.Rows) == 0 {
				continue
			}
			zones := make([]Zone, len(res.Rows))
			for i, r := range res.Rows {
				zones[i] = Zone{
					Id:        r.ID,
					Name:      r.ZoneName,
					ChannelId: r.ChannelID,
					Host:      d.Model.Host,
					Start:     r.Start,
					Finish:    r.Finish,
					Tag:       r.Tag,
				}
			}
			d.Locker.Lock()
			d.Zones = append(d.Zones, zones...)
			d.Locker.Unlock()
		}
	}
}

func (d *DTS) Run() {
	if atomic.LoadInt32(&d.Status) == 1 {
		log.Println("already run")
		return
	}
	d.Client = dtssdk.NewDTSClient(d.Model.Host)
	d.Client.CallConnected(func(s string) {
		log.Println("success connected ", s)
		res, err := d.Client.GetDeviceID()
		if err != nil {
			log.Println("GetDeviceID", err)
		} else if res.Success {
			d.DeviceId = res.DeviceID
		}

		atomic.StoreInt32(&d.Status, 1)
		go d.getZonesFromSDK()

		_ = d.Client.CallZoneAlarmNotify(func(notify *model.ZoneAlarmNotify, err error) {

		})
		_ = d.Client.CallZoneTempNotify(func(notify *model.ZoneTempNotify, err error) {

		})

		d.Client.CallDisconnected(func(s string) {
			atomic.StoreInt32(&d.Status, -1)
		})
	})
}

func (d *DTS) Close() {
	d.Client.Close()
	time.Sleep(time.Millisecond * 100)
}
