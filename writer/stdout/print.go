package stdout

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	w			io.Writer
	r			func()string
}

func New() (*Printer, error) {
	return &Printer{w:os.Stdout}, nil
}

func (p *Printer) Print(data []byte) error {
	_, err := fmt.Fprint(p.w, string(data))
	return err
}

func (p *Printer) Rotate() error {
	return nil
}

func (p *Printer) Exit() {
}