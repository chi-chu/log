package log

import "github.com/chi-chu/log/define"

const (
	STDOUT_NONE					= "%c【%s%c"
	STDOUT_GREEN				= "%c[1;40;32m%s%c[0m"
	STDOUT_YELLOW				= "%c[1;40;33m%s%c[0m"
	STDOUT_RED					= "%c[1;40;31m%s%c[0m"
	STDOUT_CLARET				= "%c[1;40;35m%s%c[0m"
	STDOUT_RED_YELLOW			= "%c[1;41;33m%s%c[0m"
)

var stdoutColor = map[define.Level]string{
	define.DEBUG:	STDOUT_NONE,
	define.INFO:	STDOUT_GREEN,
	define.WARN:	STDOUT_YELLOW,
	define.ERROR:	STDOUT_RED,
	define.PANIC:	STDOUT_CLARET,
	define.FATAL:	STDOUT_RED_YELLOW,
}

const (
	DEFAULT_DEBUG_TIPS			= "【DEBUG】"
	DEFAULT_INFO_TIPS			= "【INFO】 "
	DEFAULT_WARN_TIPS			= "【WARN】 "
	DEFAULT_ERROR_TIPS			= "【ERROR】"
	DEFAULT_PANIC_TIPS			= "【PANIC】"
	DEFAULT_FATAL_TIPS			= "【FATAL】"
)

var stdoutMsg = map[define.Level]string{
	define.DEBUG:	DEFAULT_DEBUG_TIPS,
	define.INFO:	DEFAULT_INFO_TIPS,
	define.WARN:	DEFAULT_WARN_TIPS,
	define.ERROR:	DEFAULT_ERROR_TIPS,
	define.PANIC:	DEFAULT_PANIC_TIPS,
	define.FATAL:	DEFAULT_FATAL_TIPS,
}

type FormatType int
const (
	FORMAT_NONE	FormatType		= 0
	FORMAT_TEXT	FormatType		= 1
	FORMAT_JSON	FormatType		= 2
)

const (
	ROTATE_MINITE				= "*/1 * * * *"
	ROTATE_HOUR					= "0 * * * *"
	ROTATE_DAY					= "0 0 * * *"
	ROTATE_WEEK					= "0 0 * * 0"
	ROTATE_MONTH				= "0 0 1 * *"
)