package mysql

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type printer struct {
	tableName			string
	r					func(string)string
	engine				*gorm.DB
	rotate				bool
	rotatePlan			string
	tableMod			interface{}
}

func New(dsn string, tableName string, tableMod interface{}) (*printer, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	o := &printer{
		tableName: tableName,
		rotatePlan: "0 0 * * *",//default daily
		engine: db,
		tableMod: tableMod,
	}
	// 自动迁移模式
	//db.AutoMigrate(&o.tableMod)
	return o, nil
}

func (p *printer) Print(data []byte) error {
	o := p.tableMod
	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	return p.engine.Model(p.tableName).Create(&o).Error
}

func (p *printer) Rotate() error {

	return nil
}

func (p *printer) Exit() {

}

func (p *printer) SetRotateFlag(b bool) *printer {
	p.rotate = b
	return p
}

func (p *printer) SetRotatePlan(plan string) *printer {
	p.rotate = true
	p.rotatePlan = plan
	return p
}
