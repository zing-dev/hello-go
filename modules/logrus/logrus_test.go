package logrus

import (
	"bytes"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	log2 "log"
	"os"
	"testing"
	"time"
)

func TestLog1(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func TestLog2(t *testing.T) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

func TestLog3(t *testing.T) {
	var l = log.New()
	// The API for setting attributes is a little different than the modules level
	// exported logger. See Godoc.
	l.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	}).Fatal("Failed to send event")
}

type Log struct {
	buffer  bytes.Buffer
	Content chan string
}

func (l *Log) String() {
	for {
		content := <-l.Content
		log2.Println(content)
	}
}

func (l *Log) Write(p []byte) (n int, err error) {
	go func() {
		l.Content <- string(p)
	}()
	return l.buffer.Write(p)
}

func TestLog4(t *testing.T) {
	l := &Log{Content: make(chan string)}
	go func() {
		l.String()
	}()
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	//logger.SetOutput(l)
	logger.AddHook(lfshook.NewHook(l, &log.JSONFormatter{}))
	logger.Info("info")
	logger.Warn("warn")
	time.Sleep(time.Minute)
}
