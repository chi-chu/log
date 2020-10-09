package mongo

import "fmt"

type Config struct {
	User			string
	Password		string
	IP				string
	Port			int
	Database		string
}

const DSN_URL 		= "mongodb://%s:%s@%s:%d/%s?authSource=%s"
const DSN_WITHOUT_PWD = "mongodb://%s:%d/%s"
func (c *Config) String() string {
	if c.Port == 0 {
		c.Port = 27017
	}
	if c.User == "" && c.Password == "" {
		return fmt.Sprintf(DSN_WITHOUT_PWD, c.IP, c.Port, c.Database)
	}
	return fmt.Sprintf(DSN_URL, c.User, c.Password, c.IP, c.Port, c.Database, c.Database)
}
