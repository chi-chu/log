package mysql

import (
	"errors"
	"fmt"
	"github.com/chi-chu/log/entry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type printer struct {
	tableName			string
	rotateName			string
	engine				*gorm.DB
	tableMod			interface{}
	pool				sync.Pool
}

func New(dsn string, tableName string, tableMod interface{}) (*printer, error) {
	if tableName == "" {
		return nil, errors.New("table name can`t be nil")
	}
	if tableMod == nil {
		return nil, errors.New("tableMod can`t be nil")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	o := &printer{
		tableName: tableName,
		engine: db,
		tableMod: tableMod,
	}
	o.pool.New = func()interface{}{return make(map[string]interface{})}
	return o, nil
}

func (p *printer) Print(e *entry.Entry) {
	o := p.tableMod
	//code to json and decode to struct test slowly
	//t1 := time.Now().UnixNano()
	//d, err := json.Marshal(e.Data)
	//if err != nil {
	//	panic(err)
	//}
	//err = json.Unmarshal(d, &o)
	//if err != nil {
	//	panic(err)
	//}
	//err = p.engine.Table(p.tableName).Create(o).Error
	//if err!= nil {
	//	fmt.Println("log insert into mysql err: ", err)
	//}
	//t2 := time.Now().UnixNano()

	b := p.pool.Get().(map[string]interface{})
	for k,v := range e.Data {
		b[k] = v
	}
	err := p.engine.Table(p.rotateName).Model(&o).Create(b).Error
	if err != nil {
		fmt.Println("log insert into mysql err: ", err)
	}
	p.pool.Put(b)
	//fmt.Println(t2-t1,  time.Now().UnixNano() - t2)
}

func (p *printer) Rotate(b bool) error {
	rotateName := p.tableName
	if b {
		rotateName = p.defaultRotate()
	}
	//create table
	//exist table  check column
	if !p.engine.Migrator().HasTable(rotateName) {
		err := p.engine.Set("gorm:table_options","ENGINE=Archive").
			Table(rotateName).Migrator().CreateTable(p.tableMod)
		if err != nil {
			return err
		}
	}
	p.rotateName = rotateName
	return nil
}

func (p *printer) Exit() {
}

func (p *printer) defaultRotate() string {
	return p.tableName + "_" + time.Now().Format("200601021504")
}