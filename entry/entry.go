package entry

import (
	"bytes"
	"github.com/chi-chu/log/define"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)
const (
	default_caller_trace		= 3
)

var levelTip = map[define.Level]string{
	define.DEBUG:		"debug",
	define.INFO:		"info",
	define.WARN:		"warn",
	define.ERROR:		"error",
	define.PANIC:		"panic",
	define.FATAL:		"fatal",
}

var p sync.Pool
var doOnce sync.Once
var projectPath int
type Entry struct {
	Data  map[string]string
	Level define.Level
	Buf   bytes.Buffer
}

func init() {
	p = sync.Pool{New: func() interface{} {
		return &Entry{Data:make(map[string]string), Buf: bytes.Buffer{}}
	}}
}

func NewEntry(level define.Level) *Entry {
	//use go.mod file to check this project...
	//maybe its not perfect
	doOnce.Do(func(){
		_, filename, _, ok := runtime.Caller(default_caller_trace+2)
		if ok {
			for {
				filename = filepath.Dir(filename)
				if filename == string(os.PathSeparator) {
					break
				} else {
					_, err := os.Stat(filename + string(os.PathSeparator) + "go.mod")
					if err == nil {
						projectPath = len(filepath.Dir(filename))
					}
				}
			}
		}
	})
	o := p.Get().(*Entry)
	o.Level = level
	o.Data[define.TIPS_LEVEL] = levelTip[level]
	o.Data[define.TIPS_TIME] = time.Now().Format(define.TIME_FORMAT)
	o.getCaller()
	return o
}

func (e *Entry) Release() {
	e.Data = map[string]string{}
	e.Buf.Reset()
	p.Put(e)
}

func (e *Entry) getCaller() {
	var funcName string
	pc, filename, line, ok := runtime.Caller(default_caller_trace)
	if ok {
		funcName = strings.TrimPrefix(filepath.Ext(runtime.FuncForPC(pc).Name()), ".")
	}
	e.Data[define.TIPS_FILE] = filename[projectPath:]
	e.Data[define.TIPS_LINE] = strconv.Itoa(line)
	e.Data[define.TIPS_FUNC] = funcName
}

func (e *Entry) Hook(h Hook) {
	if h != nil {
		h.Set(e)
	}
}