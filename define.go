package log

const (
	STDOUT_NONE					= "%s"
	STDOUT_GREEN				= "%c[1;40;32m%s%c[0m"
	STDOUT_YELLOW				= "%c[1;40;33m%s%c[0m"
	STDOUT_RED					= "%c[1;40;31m%s%c[0m"
	STDOUT_CLARET				= "%c[1;40;35m%s%c[0m"
	STDOUT_RED_YELLOW			= "%c[1;41;33m%s%c[0m"
)

type Level int
type FormatType int

const (
	DEBUG				Level 	= 0
	INFO				Level	= 1
	WARN				Level	= 2
	ERROR				Level	= 3
	PANIC				Level	= 4
	FATAL				Level	= 5
)

var levelTip = map[Level]string{
	DEBUG:		"debug",
	INFO:		"info",
	WARN:		"warn",
	ERROR:		"error",
	PANIC:		"panic",
	FATAL:		"fatal",
}

const (
	DEFAULT_DEBUG_TIPS			= "【DEBUG】"
	DEFAULT_INFO_TIPS			= "【INFO】 "
	DEFAULT_WARN_TIPS			= "【WARN】 "
	DEFAULT_ERROR_TIPS			= "【ERROR】"
	DEFAULT_PANIC_TIPS			= "【PANIC】"
	DEFAULT_FATAL_TIPS			= "【FATAL】"
)

const (
	DEFAULT_CALLER_TRACE		= 3
)

const (
	FORMAT_NONE	FormatType		= 0
	FORMAT_TEXT	FormatType		= 1
	FORMAT_JSON	FormatType		= 2
)

const (
	TIPS_FUNC					= "func"
	TIPS_FILE					= "file"
	TIPS_TIME					= "time"
	TIPS_LEVEL					= "level"
	TIPS_MSG					= "msg"
)

const TIME_FORMAT  				= "2006-01-02 15:04:05"