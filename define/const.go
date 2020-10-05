package define

type Level int
const (
	DEBUG				Level 	= 0
	INFO				Level	= 1
	WARN				Level	= 2
	ERROR				Level	= 3
	PANIC				Level	= 4
	FATAL				Level	= 5
)

const (
	TIPS_FUNC					= "func"
	TIPS_FILE					= "file"
	TIPS_LINE					= "line"
	TIPS_TIME					= "time"
	TIPS_LEVEL					= "level"
	TIPS_MSG					= "msg"
)

const TIME_FORMAT  				= "2006-01-02 15:04:05"