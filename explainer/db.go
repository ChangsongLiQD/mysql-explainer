package explainer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Conn struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

var (
	db *gorm.DB
)

func InitDb(conn Conn) error{
	var err error
	connText := conn.getConnText()
	db, err = gorm.Open("mysql", connText)
	if err != nil {
		return err
	}
	return nil
}

func GetExplainResult(query string) Explain{
	exp := Explain{}
	db.Raw("EXPLAIN " + query).Scan(exp)

	return exp
}

func (c *Conn)getConnText()string{
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)
}