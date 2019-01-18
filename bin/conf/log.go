package conf

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/pkg/errors"
	"path"
)

// config logrus log to local filesystem, with file rotation
func LocalLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)

	lowLv := baseLogPaht+".info"
	writer, err := rotatelogs.New(
		lowLv+".%Y%m%d%H%M",
		// rotatelogs.WithLinkName(lowLv), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	higLv := baseLogPaht+".error"
	writerErr, err := rotatelogs.New(
		higLv+".%Y%m%d%H%M",
		// rotatelogs.WithLinkName(higLv), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)

	hook := lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writerErr,
		log.FatalLevel: writerErr,
		log.PanicLevel: writerErr,
	}
	lfHook := lfshook.NewHook(hook,&log.TextFormatter{DisableColors: true})
	log.AddHook(lfHook)
}