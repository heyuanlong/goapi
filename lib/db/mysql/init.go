package mysql

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	kconf "goapi/lib/conf"
)

var (
	MysqlClient     *sql.DB
)

func init() {
	
}

func MysqlInit()  {

	c, _ := kconf.GetConf()
	user,_ := c.String("mysql","user")
	password,_ := c.String("mysql","password")
	ip,_ := c.String("mysql","ip")
	port,_ := c.String("mysql","port")
	mysqldb,_ := c.String("mysql","mysqldb")


	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user,
		password,
		ip,
		port,
		mysqldb,
	)
	MysqlClient, _ = sql.Open("mysql", addr)
	MysqlClient.SetMaxOpenConns(2000)
	MysqlClient.SetMaxIdleConns(10)
	MysqlClient.Ping()
}


