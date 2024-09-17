package logs

import (
	"sync"
	"github.com/sirupsen/logrus"
)

type LoggerAdapter struct {
	logs []logrus.Entry
	mu   sync.Mutex
}

func NewLoggerAdapter() *LoggerAdapter {
	return &LoggerAdapter{
		logs: []logrus.Entry{},
	}
}

func (l *LoggerAdapter) LogError(err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry := logrus.Entry{
		Message: err.Error(),
		Level:   logrus.ErrorLevel,
	}
	l.logs = append(l.logs, entry)
	logrus.WithFields(logrus.Fields{
		"type": "error",
		"msg":  err.Error(),
	}).Error("Error occurred")
}

func (l *LoggerAdapter) LogInfo(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry := logrus.Entry{
		Message: msg,
		Level:   logrus.InfoLevel,
	}
	l.logs = append(l.logs, entry)
	logrus.WithFields(logrus.Fields{
		"type": "info",
		"msg":  msg,
	}).Info("Info logged")
}

func (l *LoggerAdapter) LogFileStats(stats string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry := logrus.Entry{
		Message: stats,
		Level:   logrus.InfoLevel,
	}
	l.logs = append(l.logs, entry)
	logrus.WithFields(logrus.Fields{
		"type": "file_stats",
		"data": stats,
	}).Info("File stats logged")
}

func (l *LoggerAdapter) RetrieveLogs() []logrus.Entry {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.logs
}
