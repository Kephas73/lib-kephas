package constant

const (
	LogBackendPrefix   string = "["
	LogBackendSuffixes string = "]"
	LogBackendLine     string = "-"
	LogBackEnd         string = "BACKEND"
	LogSuccess         string = "success"
	LogError           string = "error"
	LogDebug           string = "debug"

	LogBackEndMainTracePrefix string = "TRACE"
	LogBackEndMainDebugPrefix string = "DEBUG"
	LogBackEndMainInfoPrefix  string = "INFO"
	LogBackEndMainWarnPrefix  string = "WARN"
	LogBackEndMainErrorPrefix string = "ERROR"
	LogBackEndMainFatalPrefix string = "FATAL"
	LogBackEndMainPanicPrefix string = "PANIC"
)
const (
	LogTracePrefix = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainTracePrefix + LogBackendSuffixes
	LogDebugPrefix = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainDebugPrefix + LogBackendSuffixes
	LogInfoPrefix  = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainInfoPrefix + LogBackendSuffixes
	LogWarnPrefix  = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainWarnPrefix + LogBackendSuffixes
	LogErrorPrefix = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainErrorPrefix + LogBackendSuffixes
	LogFatalPrefix = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainFatalPrefix + LogBackendSuffixes
	LogPanicPrefix = LogBackendPrefix + LogBackEnd + LogBackendLine + LogBackEndMainPanicPrefix + LogBackendSuffixes
)
