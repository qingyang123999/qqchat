package qqlog

import (
	"github.com/sirupsen/logrus"
	"time"
)

// DefaultFieldHook 添加默认字段
type DefaultFieldHook struct {
	AppName string
	Env     string
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	entry.Data["env"] = h.Env
	entry.Data["timestamp"] = time.Now().UnixNano()
	return nil
}
