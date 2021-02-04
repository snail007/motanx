package motanx

import (
	_ "github.com/snail007/gmc"
	gcore "github.com/snail007/gmc/core"
	gconfig "github.com/snail007/gmc/module/config"
	glog "github.com/snail007/gmc/module/log"
	vlog "github.com/weibocom/motan-go/log"
)

func init() {
	apiLogger := NewLoggerFromConfig(gconfig.NewConfig())
	vlog.LogInit(apiLogger)
}

type Logger struct {
	*glog.GMCLog
}

func NewLoggerFromConfig(config gcore.Config, prefix ...string) vlog.Logger {
	pre := ""
	if len(prefix) == 1 {
		pre = prefix[0]
	}
	log := glog.NewFromConfig(config, pre).(*glog.GMCLog)
	log.SetCallerSkip(4)
	return &Logger{
		log,
	}
}

func (l Logger) Info(i ...interface{}) {
	l.GMCLog.Info(i...)
}

func (l Logger) Infof(s string, i ...interface{}) {
	l.GMCLog.Infof(s, i...)
}

func (l Logger) Infoln(i ...interface{}) {
	l.GMCLog.Info(i...)
}

func (l Logger) Warn(i ...interface{}) {
	l.GMCLog.Warn(i...)
}

func (l Logger) Warningf(s string, i ...interface{}) {
	l.GMCLog.Warnf(s, i...)
}

func (l Logger) Warningln(i ...interface{}) {
	l.GMCLog.Warn(i...)
}

func (l Logger) Error(i ...interface{}) {
	l.GMCLog.Warn(i...)
}

func (l Logger) Errorf(s string, i ...interface{}) {
	l.GMCLog.Warnf(s, i...)
}

func (l Logger) Errorln(i ...interface{}) {
	l.GMCLog.Warn(i...)
}

func (l Logger) Fatal(i ...interface{}) {
	l.GMCLog.Panic(i...)
}

func (l Logger) Fatalf(s string, i ...interface{}) {
	l.GMCLog.Panicf(s, i...)
}

func (l Logger) Fatalln(i ...interface{}) {
	l.GMCLog.Panic(i...)
}

func (l Logger) AccessLog(entity *vlog.AccessLogEntity) {
}

func (l Logger) MetricsLog(s string) {
}

func (l Logger) Flush() {
	l.WaitAsyncDone()
}

func (l Logger) SetAsync(b bool) {
	if b {
		l.EnableAsync()
	}
}

func (l Logger) GetLevel() vlog.LogLevel {
	return vlog.LogLevel(l.Level())
}

func (l Logger) SetLevel(level vlog.LogLevel) {
	theLevel := gcore.LTRACE
	switch level {
	case vlog.TraceLevel:
		theLevel = gcore.LTRACE
	case vlog.DebugLevel:
		theLevel = gcore.LDEBUG
	case vlog.InfoLevel:
		theLevel = gcore.LINFO
	case vlog.WarnLevel:
		theLevel = gcore.LWARN
	case vlog.ErrorLevel:
		theLevel = gcore.LWARN
	case vlog.DPanicLevel:
		theLevel = gcore.LPANIC
	case vlog.PanicLevel:
		theLevel = gcore.LPANIC
	case vlog.FatalLevel:
		theLevel = gcore.LERROR
	}
	l.GMCLog.SetLevel(gcore.LogLevel(theLevel))
}

func (l Logger) SetAccessStructured(b bool) {

}

func (l Logger) GetAccessLogAvailable() bool {
	return false
}

func (l Logger) SetAccessLogAvailable(b bool) {

}

func (l Logger) GetMetricsLogAvailable() bool {
	return false
}

func (l Logger) SetMetricsLogAvailable(b bool) {

}
