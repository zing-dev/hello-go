package expvar_test

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var vi = expvar.NewInt("vi")
	var vf = expvar.NewFloat("vf")
	var vs = expvar.NewString("vs")
	var vm = expvar.NewMap("vm")
	vi.Add(1)
	vf.Set((float64(vi.Value())) * 10)
	vs.Set(fmt.Sprintf("%d %f", vi.Value(), vf.Value()))
	vm.Set("i", vi)
	vm.Set("f", vf)
	vm.Set("s", vs)
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func TestNew(t *testing.T) {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8080", nil)
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}

func NewUser(name string, u *User) *User {
	expvar.Publish(name, u)
	return u
}

func NewDefaultUser(name string) *User {
	u := new(User)
	expvar.Publish(name, u)
	return u
}

func (u *User) Set() string {
	pointer := unsafe.Pointer(u)
	atomic.StorePointer(&pointer, pointer)
	data, _ := json.Marshal(u)
	return string(data)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	NewUser("user", &User{
		Id:   1,
		Name: "zing",
	})
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func TestStruct(t *testing.T) {
	http.HandleFunc("/", handler2)
	_ = http.ListenAndServe(":8080", nil)
}

func TestDo(t *testing.T) {
	var vi = expvar.NewInt("vi")
	go func() {
		for i := 0; i < 100; i++ {
			vi.Add(1)
			time.Sleep(time.Second)
		}
	}()
	m := map[string]string{}
	for {
		expvar.Do(func(value expvar.KeyValue) {
			m[value.Key] = value.Value.String()
		})

		t.Log(m)
		time.Sleep(time.Second)
	}
}
