package fileprinter

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type printer struct {
	w			io.Writer
	dir			string
	fileName	string
	ext			string
	r			func(fileName string) (filename string)
	cron		*cron.Cron
	rotateFlag	bool
	rotateType	RotateType
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
			return nil, errors.New("[log] error to mkdir path: "+dir+"   err: " + err.Error())
		}
	}
	o := &printer{w: nil, fileName: fName, dir: dir, ext:ext, rotateType: ROTATE_DAY}
	return o, nil
}

func (p *printer) Print(data []byte) error {
	_, err := fmt.Fprint(p.w, string(data))
	return err
}

func (p *printer) Rotate() error {
	if p.rotateFlag {
		p.cron = cron.New()
		_, err := p.cron.AddFunc(string(p.rotateType), func(){ _ = p.rotate()})
		if err != nil {
			return err
		}
		return p.rotate()
	} else {
		f, err := os.OpenFile(p.dir + string(os.PathSeparator) + p.fileName + p.ext, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		p.w = f
	}
	return nil
}

func (p *printer) SetRotatePlan(r RotateType) *printer {
	p.SetRotateFlag(true)
	p.rotateType = r
	return p
}

func (p *printer) SetRotateFlag(b bool) *printer {
	p.rotateFlag = b
	return p
}

func (p *printer) SetRotateFunc(f func(string)string) *printer {
	p.SetRotateFlag(true)
	p.r = f
	return p
}

func (p *printer) rotate() error {
	var fn string
	if p.r != nil {
		fn = p.r(p.fileName)
		if fn == p.fileName {
			return nil
		} else {
			fn = p.dir + string(os.PathSeparator) + fn + p.ext
		}
	} else {
		fn = p.defaultRotate()
	}
	if fn == "" {
		return errors.New("[log]rotate func get nil string")
	}
	var f, o *os.File
	var err error
	f, err = os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	if p.w != nil {
		o = p.w.(*os.File)
		defer o.Close()
	}
	p.w = f
	return nil
}

func (p *printer) Exit() {
	p.cron.Stop()
	if p.w != nil {
		_ = p.w.(*os.File).Close()
	}
}

func (p *printer) defaultRotate() string {
	return p.dir + string(os.PathSeparator) + p.fileName + "_" + time.Now().Format("200601021304") + p.ext
}