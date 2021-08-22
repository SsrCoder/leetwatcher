package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type logHook struct {
	path       string
	lock       *sync.RWMutex
	timeFormat string
	preTime    string
	preWriters map[logrus.Level]io.WriteCloser
}

func NewLogHook(path string, timeFormat string) *logHook {
	return &logHook{
		path:       path,
		timeFormat: timeFormat,
		lock:       &sync.RWMutex{},
		preWriters: make(map[logrus.Level]io.WriteCloser),
	}
}

func (h *logHook) Fire(e *logrus.Entry) error {
	time := e.Time.Format(h.timeFormat)
	h.lock.RLock()
	if h.preTime == time {
		if e != nil {
			if h.preWriters[e.Level] == nil {
				h.lock.RUnlock()
				h.lock.Lock()
				writer, err := h.newFileWriter(e.Level, time)
				if err != nil {
					h.lock.Unlock()
					return err
				}
				h.preWriters[e.Level] = writer
				h.lock.Unlock()
				h.lock.RLock()
			}
			bytes, err := e.Logger.Formatter.Format(e)
			if err != nil {
				return err
			}
			h.preWriters[e.Level].Write(bytes)
		}
		h.lock.RUnlock()
		return nil
	}
	h.lock.RUnlock()
	h.lock.Lock()

	for level, writer := range h.preWriters {
		if writer != nil {
			writer.Close()
		}
		h.preWriters[level] = nil
	}
	h.preTime = time
	h.lock.Unlock()
	return h.Fire(e)
}

func (h *logHook) newFileWriter(level logrus.Level, time string) (fd io.WriteCloser, err error) {
	path := fmt.Sprintf("%s/%s_%s.log", h.path, level.String(), time)
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)

	fd, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return
}

func (*logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func callerPrettyfier(f *runtime.Frame) (function string, file string) {
	file, _ = filepath.Rel(absPath, f.File)
	file = fmt.Sprintf("%s:%d", file, f.Line)
	function = f.Function
	if idx := strings.LastIndex(function, "/"); idx != -1 {
		function = function[idx+1:]
	}
	return
}

func InitLogrus() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: callerPrettyfier,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(NewLogHook("./logs", "2006010215"))
	logrus.SetOutput(&NilWriter{})
}
