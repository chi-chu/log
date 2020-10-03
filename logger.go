package log

import (
	"fmt"
	"github.com/chi-chu/log/writer/stdoutprinter"
)

type IPrint interface {
	Print([]byte) error
	Rotate() error
	Exit()
}

type logger struct {
	Out				IPrint
	Hook			Hook
	WithColorTip	bool
	TextFormat		FormatType
	ReportCaller	bool
	TraceSkip		int
	Level			Level
	exitFunc		func()
}

var log *logger
var stdout bool
func New(p IPrint) *logger {
	if _, ok := p.(*stdoutprinter.Printer); ok {
		stdout = true
	}
	 err := p.Rotate()
	 if err != nil {
	 	fmt.Println("[log]new log err: ", err)
	 	return nil
	 }
	log = &logger{
		Out:          p,
		Hook:         nil,
		WithColorTip: true,
		TextFormat:   FORMAT_JSON,
		ReportCaller: false,
		Level:        DEBUG,
		exitFunc:     nil,
	}
	return log
}

func (l *logger) SetLevel(e Level) *logger {
	l.Level = e
	return l
}

func (l *logger) SetColorTip(b bool) *logger {
	l.WithColorTip = b
	return l
}

func (l *logger) SetFormat(f FormatType) *logger {
	l.TextFormat = f
	return l
}

func (l *logger) SetHook(h Hook) *logger {
	l.Hook = h
	return l
}

func (l *logger) SetReportCaller(b bool) *logger {
	l.ReportCaller = b
	return l
}

func (l *logger) SetExitFunc(f func()) *logger {
	l.exitFunc = f
	return l
}

func (l *logger) SetTraceSkip(n int) *logger {
	l.TraceSkip = n
	return l
}

func Exit() {
	log.exitFunc()
	log.Out.Exit()
}