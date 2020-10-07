package log

import (
	"github.com/chi-chu/log/define"
	"github.com/chi-chu/log/entry"
	"os"
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
	e.Hook(log.Hook)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Warn(format string, args ...interface{}) {
	if log.Level > define.WARN  {
		return
	}
	e := entry.NewEntry(define.WARN)
	e.Hook(log.Hook)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Error(format string, args ...interface{}) {
	if log.Level > define.ERROR {
		return
	}
	e := entry.NewEntry(define.ERROR)
	e.Hook(log.Hook)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
}

func Panic(format string, args ...interface{}) {
	if log.Level > define.PANIC {
		return
	}
	e := entry.NewEntry(define.PANIC)
	e.Hook(log.Hook)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
	os.Exit(1)
}

func Fatal(format string, args ...interface{}) {
	if log.Level > define.FATAL {
		return
	}
	e := entry.NewEntry(define.FATAL)
	e.Hook(log.Hook)
	formatFunc[log.TextFormat](e, format, args...)
	log.Out.Print(e)
	e.Release()
	os.Exit(1)
}