package log

import (
	"github.com/chi-chu/log/define"
	"github.com/chi-chu/log/entry"
)

func Debug(format string, args ...interface{}) {
	if log.Level > define.DEBUG {
		return
	}
	e := entry.NewEntry(define.DEBUG)
	e.Hook(log.Hook)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Info(format string, args ...interface{}) {
	if log.Level > define.INFO {
		return
	}
	e := entry.NewEntry(define.INFO)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Warn(format string, args ...interface{}) {
	if log.Level > define.WARN  {
		return
	}
	e := entry.NewEntry(define.WARN)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Error(format string, args ...interface{}) {
	if log.Level > define.ERROR {
		return
	}
	e := entry.NewEntry(define.ERROR)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Panic(format string, args ...interface{}) {
	if log.Level > define.PANIC {
		return
	}
	e := entry.NewEntry(define.PANIC)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Fatal(format string, args ...interface{}) {
	if log.Level > define.FATAL {
		return
	}
	e := entry.NewEntry(define.FATAL)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}