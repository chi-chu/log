package log

import (
	"bytes"
	"fmt"
	"github.com/chi-chu/log/entry"
	"io"
	"os"
	"sync"
)

type printer struct {
	w			io.Writer
	p			sync.Pool
}

func newStdout() *printer {
	o := &printer{w:os.Stdout}
	o.p.New = func() interface{} {
		return bytes.Buffer{}
	}
	return o
}

func (p *printer) Print(e *entry.Entry) {
	buf := p.p.Get().(bytes.Buffer)
	buf.WriteString(fmt.Sprintf(stdoutColor[e.Level], 0x1B, stdoutMsg[e.Level], 0x1B))
	buf.Write(e.Buf.Bytes())
	buf.WriteString("\n")
	_, _ = fmt.Fprint(p.w, buf.String())
	buf.Reset()
	p.p.Put(buf)
}

func (p *printer) Rotate(b bool) error {
	return nil
}

func (p *printer) Exit() {
}