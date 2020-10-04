package mysql

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type printer struct {
	tableName			string
	r					func(string)string
	engine				*xorm.Engine
	rotate				bool
	rotatePlan			string
}

func New() (*printer, error) {
	engine, err := xorm.NewEngine("mysql", sqlUrl)
	if err != nil {
		return nil, err
	}
	if err = engine.Ping(); err != nil {
		return nil, errors.New("xorm engine ping fail")
	}
	engine.ShowSQL(false)
	engine.SetMaxOpenConns(300)
	engine.SetMaxIdleConns(300)
	return &printer{
		engine:			engine,
		rotatePlan: 	"0 0 * * *",
	}, nil
}

func (p *printer) Print(data []byte) error {
	return nil
}

func (p *printer) Rotate() error {
	return nil
}

func (p *printer) Exit() {

}
