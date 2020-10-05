package log

import (
	"github.com/chi-chu/log/define"
	"github.com/chi-chu/log/entry"
	"github.com/robfig/cron/v3"
)

type IPrint interface {
	Print(*entry.Entry)
	Rotate(bool) error
	Exit()
}

type Option func(*logger)

type logger struct {
	Out          IPrint
	Hook         entry.Hook
	TextFormat   FormatType
	ReportCaller bool
	Level        define.Level
	RotateFlag	 bool
	RotateType	 string
	cron		 *cron.Cron
}

var log *logger
var newPrinter bool
func init() {
	log = &logger{
		Out:          newStdout(),
		Hook:         nil,
		TextFormat:   FORMAT_JSON,
		ReportCaller: true,
		Level:        define.DEBUG,
		RotateFlag:	  false,
		RotateType:	  ROTATE_DAY,
	}
}

func Opt(opt ...Option) {
	for _, f := range opt {
		f(log)
	}
	//to insure that the rotate action after user inject IPrint
	if newPrinter {
		for _, f := range opt {
			f(log)
		}
	}
}

func SetWriterAndRotate(print IPrint, on bool, rotateType string) Option {
	return func(l *logger) {
		l.Out = print
		l.RotateFlag = on
		if rotateType != "" {
			rotateType = ROTATE_DAY
		}
		l.RotateType = rotateType
		err := log.Out.Rotate(on)
		if err != nil {
			Error("log rotate err: ", err)
		}
		if on {
			log.cron = cron.New()
			_, err := log.cron.AddFunc(rotateType, func() {
				err := log.Out.Rotate(on)
				if err != nil {
					Error("log rotate err: ", err)
				}
			})
			if err != nil {
				Error("log rotate err: ", err)
				return
			}
			l.cron.Start()
		}
	}
}

func SetLevel(level define.Level) Option {
	return func(l *logger) {
		l.Level = level
	}
}

func SetFormat(formatType FormatType) Option {
	return func(l *logger) {
		l.TextFormat = formatType
	}
}

func SetHook(hook entry.Hook) Option {
	return func(l *logger) {
		l.Hook = hook
	}
}

func SetReportCaller(on bool) Option {
	return func(l *logger) {
		l.ReportCaller = on
	}
}

func Exit() {
	if log.cron != nil {
		log.cron.Stop()
	}
	log.Out.Exit()
}