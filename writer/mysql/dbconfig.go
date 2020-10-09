package mysql

import "fmt"

type Config struct {
	User			string
	Password		string
	IP				string
	Port			int
	DBname			string
}

const DSN_URL 		= "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func (c *Config) String() string {
	if c.Port == 0 {
		c.Port = 3306
	}
	return fmt.Sprintf(DSN_URL, c.User, c.Password, c.IP, c.Port, c.DBname)
}