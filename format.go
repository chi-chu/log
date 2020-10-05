package log

import (
	"encoding/json"
	"fmt"
	"github.com/chi-chu/log/define"
	"github.com/chi-chu/log/entry"
)

var formatFunc = map[FormatType]func(*entry.Entry, string, ...interface{}){
	FORMAT_NONE: nil,
	FORMAT_TEXT: formatText,
	FORMAT_JSON: formatJson,
}

func formatText(e *entry.Entry, f string, args ...interface{}) {
	e.Buf.WriteString(e.Data[define.TIPS_TIME])
	e.Buf.WriteString(" ")
	e.Buf.WriteString(e.Data[define.TIPS_LEVEL])
	e.Buf.WriteString(" ")
	if log.ReportCaller {
		e.Buf.WriteString(e.Data[define.TIPS_FILE])
		e.Buf.WriteString(":")
		e.Buf.WriteString(e.Data[define.TIPS_LINE])
		e.Buf.WriteString(" ")
		e.Buf.WriteString(e.Data[define.TIPS_FUNC])
		e.Buf.WriteString(" ")
	}
	e.Buf.WriteString(define.TIPS_MSG + ": ")
	e.Buf.WriteString(fmt.Sprintf(f, args...))
	//e.Data[define.TIPS_MSG] = fmt.Sprintf(f, args...)
	for k, v := range e.Data {
		if k == define.TIPS_TIME || k == define.TIPS_FUNC ||
			k == define.TIPS_FILE || k == define.TIPS_MSG ||
			k == define.TIPS_LEVEL || k == define.TIPS_LINE {
			continue
		}
		e.Buf.WriteString(" ")
		e.Buf.WriteString(k)
		e.Buf.WriteString(":")
		e.Buf.WriteString(v)
	}
}

func formatJson(e *entry.Entry, f string, args ...interface{}) {
	e.Data[define.TIPS_MSG] = fmt.Sprintf(f, args...)
	d, err := json.Marshal(e.Data)
	if err != nil {
		fmt.Println(err)
	}
	e.Buf.Write(d)
}