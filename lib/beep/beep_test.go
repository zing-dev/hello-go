package beep

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"testing"
	"time"

	"github.com/faiface/beep/speaker"
)

const (
	alarm = "alarm.mp3"
)

func TestWav(t *testing.T) {
	f, err := os.Open("alarm.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer))
	time.Sleep(time.Second * 30)
	speaker.Close()
	time.Sleep(time.Second * 10)
}

func TestStop(t *testing.T) {
	f, err := os.Open("alarm.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer))
	time.Sleep(time.Second * 3)
	speaker.Close()
	time.Sleep(time.Second * 10)
}

func TestPlay(t *testing.T) {
	f, err := os.Open(alarm)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	log.Println(time.Now())
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		log.Println("结束....")
		log.Println(time.Now())
	})))
	log.Println(time.Now())
	select {}
}

func TestName(t *testing.T) {
	f, err := os.Open(alarm)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
