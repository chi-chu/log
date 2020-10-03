package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var p sync.Pool

type Entry struct {
	Data			map[string]string
	level			Level
	buf				bytes.Buffer
}

func init() {
	p = sync.Pool{New: func() interface{} {
		return &Entry{Data:make(map[string]string), buf: bytes.Buffer{}}
	}}
}

func newEntry(level Level) *Entry {
	o := p.Get().(*Entry)
	o.level = level
	o.getCaller()
	return o
}

func (e *Entry) release() {
	e.Data = map[string]string{}
	e.buf.Reset()
	p.Put(e)
}

func (e *Entry) getCaller() {
	if log.ReportCaller {
		var funcName string
		pc, filename, line, ok := runtime.Caller(DEFAULT_CALLER_TRACE + log.TraceSkip)
		if ok {
			funcName = strings.TrimPrefix(filepath.Ext(runtime.FuncForPC(pc).Name()), ".")
		}
		e.Data[TIPS_FILE] = filename + ":" + strconv.Itoa(line)
		e.Data[TIPS_FUNC] = funcName
	}
	e.Data[TIPS_TIME] = time.Now().Format(TIME_FORMAT)
	e.Data[TIPS_LEVEL] = levelTip[e.level]
}

func (e *Entry) opHook() {
	if log.Hook != nil {
		log.Hook.Set(e)
	}
}

func (e *Entry) format(f string, args ...interface{}) []byte {
	if log.TextFormat == FORMAT_TEXT {
		//e.buf.WriteString(TIPS_TIME + ": ")
		e.buf.WriteString(e.Data[TIPS_TIME])
		e.buf.WriteString(" ")
		//e.buf.WriteString(TIPS_LEVEL + ": ")
		e.buf.WriteString(e.Data[TIPS_LEVEL])
		e.buf.WriteString(" ")
		if log.ReportCaller {
			//e.buf.WriteString(TIPS_FILE + ": ")
			e.buf.WriteString(e.Data[TIPS_FILE])
			e.buf.WriteString(" ")
			//e.buf.WriteString(TIPS_FUNC + ": ")
			e.buf.WriteString(e.Data[TIPS_FUNC])
			e.buf.WriteString(" ")
		}
		e.buf.WriteString(TIPS_MSG + ": ")
		e.buf.WriteString(fmt.Sprintf(f, args...))
	} else if log.TextFormat == FORMAT_JSON {
		e.Data[TIPS_MSG] = fmt.Sprintf(f, args...)
		d, err := json.Marshal(e.Data)
		if err != nil && !stdout {
			fmt.Println(err)
		}
		e.buf.Write(d)
	}
	e.buf.WriteString("\n")
	return nil
}

func (e *Entry) write() {
	err := log.Out.Print(e.buf.Bytes())
	if err != nil && !stdout {
		fmt.Println("[log] error to write err:", err)
	}
	e.release()
}

func Debug(format string, args ...interface{}) {
	if log.Level > DEBUG {
		return
	}
	e := newEntry(DEBUG)
	if stdout && log.WithColorTip{
		e.buf.WriteString(fmt.Sprintf(STDOUT_NONE, 0x1B, DEFAULT_DEBUG_TIPS, 0x1B))
	}
	e.format(format, args...)
	e.write()
}

func Info(format string, args ...interface{}) {
	if log.Level > INFO {
		return
	}
	e := newEntry(INFO)
	if stdout && log.WithColorTip{
		e.buf.WriteString(fmt.Sprintf(STDOUT_GREEN, 0x1B, DEFAULT_INFO_TIPS, 0x1B))
	}
	e.format(format, args...)
	e.write()
}

func Warn(format string, args ...interface{}) {
	if log.Level > WARN {
		return
	}
	e := newEntry(WARN)
	if stdout && log.WithColorTip{
		e.buf.WriteString(fmt.Sprintf(STDOUT_YELLOW, 0x1B, DEFAULT_WARN_TIPS, 0x1B))
	}
	e.format(format, args...)
	e.write()
}

func Error(format string, args ...interface{}) {
	if log.Level > ERROR {
		return
	}
	e := newEntry(ERROR)
	if stdout && log.WithColorTip{
		e.buf.WriteString(fmt.Sprintf(STDOUT_RED, 0x1B, DEFAULT_ERROR_TIPS, 0x1B))
	}
	e.format(format, args...)
	e.write()
}

func Panic(format string, args ...interface{}) {
	if log.Level > PANIC {
		return
	}
	e := newEntry(PANIC)
	if stdout && log.WithColorTip{
		e.buf.WriteString(fmt.Sprintf(STDOUT_CLARET, 0x1B, DEFAULT_PANIC_TIPS, 0x1B))
	}
	e.format(format, args...)
	e.write()
	panic(e.buf.String())
}

func Fatal(format string, args ...interface{}) {
	if log.Level > FATAL {
		return
	}
	e := newEntry(FATAL)
	if stdout && log.WithColorTip{
		e.buf.WriteString(fmt.Sprintf(STDOUT_RED_YELLOW, 0x1B, DEFAULT_FATAL_TIPS, 0x1B))
	}
	e.format(format, args...)
	e.write()
	panic(e.buf.String())
}