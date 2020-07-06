package app

import (
	"5z7Game/pkg/constant"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type LoggerInstance struct {
	once     sync.Once
	instance *logrus.Logger
}

var logger *LoggerInstance

func init() {
	logger = new(LoggerInstance)
}

// getInstance
func (l *LoggerInstance) getInstance(writer io.Writer) {
	// 实例化,实际项目中一般用全局变量来初始化一个日志管理器
	l.instance = logrus.New()

	// 设置日志内容为json格式
	l.instance.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   constant.TimeFormat, // 时间格式
		DisableTimestamp:  false,               //是否禁用日期
		DisableHTMLEscape: false,               // 是否禁用html转义
		DataKey:           "",
		CallerPrettyfier:  nil,
		PrettyPrint:       false, // 是否需要格式化
	})

	l.instance.Level = logrus.DebugLevel

	l.instance.Out = writer

}
