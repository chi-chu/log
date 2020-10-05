package file

import (
	"errors"
	"fmt"
	"github.com/chi-chu/log/entry"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type printer struct {
	w			io.Writer
	dir			string
	fileName	string
	ext			string
	lock		sync.RWMutex
}

func New(path string) (*printer, error) {
	l := len(path)
	if []byte(path)[l-1] == '/' || []byte(path)[l-1] == '\\' {
		return nil, errors.New("[log] invalid filepath path: " + path)
	}
	fName := filepath.Base(path)
	ext := filepath.Ext(fName)
	fName = strings.Trim(fName, ext)
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	o := &printer{w: nil, fileName: fName, dir: dir, ext:ext}
	return o, nil
}

func (p *printer) Print(e *entry.Entry) {
	e.Buf.WriteString("\n")
	p.lock.RLock()
	defer p.lock.RUnlock()
	_, _ = fmt.Fprint(p.w, e.Buf.String())
}

func (p *printer) Rotate(b bool) error {
	var fn string
	if b {
		fn = p.defaultRotate()
	} else {
		fn = p.dir + string(os.PathSeparator) + p.fileName + p.ext
	}
	var f *os.File
	var err error
	f, err = os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	p.lock.Lock()
	if p.w != nil {
		p.w.(*os.File).Close()
	}
	p.w = f
	p.lock.Unlock()
	return nil
}

func (p *printer) Exit() {
	if p.w != nil {
		_ = p.w.(*os.File).Close()
	}
}

func (p *printer) defaultRotate() string {
	return p.dir + string(os.PathSeparator) + p.fileName + "_" + time.Now().Format("200601021504") + p.ext
}