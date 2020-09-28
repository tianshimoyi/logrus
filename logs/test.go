package logs

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/sirupsen/logrus"
)

var (
	l = logrus.New()
	file, err   = os.OpenFile("1.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	filestus, _ = file.Stat()

	msg = make(chan []byte , 1000)
)

type CutLog struct {
}

func (c *CutLog) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (c *CutLog) Fire(e *logrus.Entry) error {
	filestus, _ = file.Stat()
	fmt.Println(filestus.Size())
	if filestus.Size() >= 1024*10 {
		file.Close()
		name:=strconv.Itoa(int(time.Now().Unix()))
		file,_=os.OpenFile(name+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		l.SetNoLock()
		l.SetOutput(file)
	}
	return nil
}

func LogT() {
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	l.SetReportCaller(true)
	l.SetOutput(file)
	l.AddHook(&CutLog{})
	for {
		l.Error("hello, world!")
		time.Sleep(time.Duration(10000) * time.Microsecond)
	}
}
