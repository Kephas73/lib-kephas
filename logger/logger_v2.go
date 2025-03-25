package logger

import (
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"path"
)

var mLog *logrus.Logger

// WriterHook giúp logrus ghi log ra nhiều file tùy theo cấp độ
type WriterHook struct {
	Writer    *lumberjack.Logger
	LogLevels []logrus.Level
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func NewWriterHook(writer *lumberjack.Logger, levels []logrus.Level) *WriterHook {
	return &WriterHook{
		Writer:    writer,
		LogLevels: levels,
	}
}

func NewLogger(debug bool) *logrus.Logger {
	if mLog != nil {
		return mLog
	}

	createConfigFromEnv()

	// Cấu hình log rotation cho từng cấp độ
	successLogger := &lumberjack.Logger{
		Filename:   path.Join(loggerConf.Path, loggerConf.Prefix+"_success.log"),
		MaxSize:    loggerConf.MaxSize,    // MB
		MaxBackups: loggerConf.MaxBackups, // Giữ tối đa 5 file backup
		MaxAge:     loggerConf.MaxAge,     // Giữ log trong 30 ngày
		Compress:   loggerConf.Compress,   // Nén log cũ
	}

	// Cấu hình log rotation cho từng cấp độ
	debugLogger := &lumberjack.Logger{
		Filename:   path.Join(loggerConf.Path, loggerConf.Prefix+"_debug.log"),
		MaxSize:    loggerConf.MaxSize,    // MB
		MaxBackups: loggerConf.MaxBackups, // Giữ tối đa 5 file backup
		MaxAge:     loggerConf.MaxAge,     // Giữ log trong 30 ngày
		Compress:   loggerConf.Compress,   // Nén log cũ
	}

	// Cấu hình log rotation cho từng cấp độ
	errorLogger := &lumberjack.Logger{
		Filename:   path.Join(loggerConf.Path, loggerConf.Prefix+"_error.log"),
		MaxSize:    loggerConf.MaxSize,    // MB
		MaxBackups: loggerConf.MaxBackups, // Giữ tối đa 5 file backup
		MaxAge:     loggerConf.MaxAge,     // Giữ log trong 30 ngày
		Compress:   loggerConf.Compress,   // Nén log cũ
	}

	// Cấu hình log rotation cho từng cấp độ
	panicLogger := &lumberjack.Logger{
		Filename:   path.Join(loggerConf.Path, loggerConf.Prefix+"panic.log"),
		MaxSize:    loggerConf.MaxSize,    // MB
		MaxBackups: loggerConf.MaxBackups, // Giữ tối đa 5 file backup
		MaxAge:     loggerConf.MaxAge,     // Giữ log trong 30 ngày
		Compress:   loggerConf.Compress,   // Nén log cũ
	}

	mLog = logrus.New()
	logFormatter := new(logrus.TextFormatter)
	logFormatter.TimestampFormat = constant.TimeFormatYYMMDDHHmmSS
	logFormatter.FullTimestamp = true

	mLog.SetFormatter(logFormatter)

	mLog.AddHook(NewWriterHook(successLogger, []logrus.Level{
		logrus.InfoLevel, logrus.TraceLevel, logrus.WarnLevel,
	}))

	mLog.AddHook(NewWriterHook(debugLogger, []logrus.Level{
		logrus.DebugLevel,
	}))

	mLog.AddHook(NewWriterHook(errorLogger, []logrus.Level{
		logrus.ErrorLevel, logrus.FatalLevel,
	}))

	mLog.AddHook(NewWriterHook(panicLogger, []logrus.Level{
		logrus.PanicLevel,
	}))

	if !debug {
		mLog.Out = io.Discard
	}

	return mLog
}

func Trace(format string, v ...interface{}) {
	mLog.Tracef(constant.LogTracePrefix+format, v)
}

func Debug(format string, v ...interface{}) {
	mLog.Debugf(constant.LogDebugPrefix+format, v...)
}

func Info(format string, v ...interface{}) {
	mLog.Infof(constant.LogInfoPrefix+format, v...)
}

func Warn(format string, v ...interface{}) {
	mLog.Warnf(constant.LogWarnPrefix+format, v...)
}

func Error(format string, v ...interface{}) {
	mLog.Errorf(constant.LogErrorPrefix+format, v...)
}

func Fatal(format string, v ...interface{}) {
	mLog.Fatalf(constant.LogFatalPrefix+format, v...)
}

func Panic(format string, v ...interface{}) {
	mLog.Panicf(constant.LogPanicPrefix+format, v...)
}
