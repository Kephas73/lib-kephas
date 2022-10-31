package logger

import (
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/timer"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
)

var mLog *logrus.Logger

func NewLogger(logPath string, logPrefix string, debug bool) *logrus.Logger {
	if mLog != nil {
		return mLog
	}

	logPathMap := lfshook.PathMap{
		logrus.InfoLevel:  path.Join(logPath, timer.Now(), logPrefix+"_success.log"),
		logrus.TraceLevel: path.Join(logPath, timer.Now(), logPrefix+"_success.log"),
		logrus.WarnLevel:  path.Join(logPath, timer.Now(), logPrefix+"_success.log"),

		logrus.DebugLevel: path.Join(logPath, timer.Now(), logPrefix+"_debug.log"),

		logrus.ErrorLevel: path.Join(logPath, timer.Now(), logPrefix+"_error.log"),
		logrus.FatalLevel: path.Join(logPath, timer.Now(), logPrefix+"_error.log"),
		logrus.PanicLevel: path.Join(logPath, timer.Now(), logPrefix+"_error.log"),
	}

	logFormatter := new(logrus.TextFormatter)
	logFormatter.TimestampFormat = constant.TimeFormatYYMMDDHHmmSS
	logFormatter.FullTimestamp = true

	mLog = logrus.New()
	mLog.Hooks.Add(lfshook.NewHook(
		logPathMap,
		logFormatter,
	))

	if !debug {
		mLog.Out = ioutil.Discard
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
	//go func() {
	//	msg := fmt.Sprintf(constant.LogErrorPrefix+format, v...)
	//	err := errors.New(msg)
	//	sentry.CaptureException(err)
	//}()

	mLog.Errorf(constant.LogErrorPrefix+format, v...)
}

func Fatal(format string, v ...interface{}) {
	//sentry.CaptureMessage(constant.LogFatalPrefix + format)
	mLog.Fatalf(constant.LogFatalPrefix+format, v...)
}

func Panic(format string, v ...interface{}) {
	//sentry.CaptureMessage(constant.LogPanicPrefix + format)
	mLog.Panicf(constant.LogPanicPrefix+format, v...)
}
