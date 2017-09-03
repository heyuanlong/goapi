package mysql

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"
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


func Test()  {
	rows, err := MysqlClient.Query("SELECT accountId,deviceId FROM account limit 2")
	if err!= nil {
		klog.Klog.Println("test mysql fail ")
		return
	}
	defer rows.Close()

	for rows.Next(){
		var accountId int
		var deviceId string

		if err := rows.Scan(&accountId,&deviceId); err != nil {
			klog.Klog.Println(err)
			return
		}
		klog.Klog.Printf("accountId=%v,deviceId=%v",accountId,deviceId)
	}

	//-------------------------------------------------------------
	accountId := SelectAccount("xxxxxxxxxx")
	klog.Klog.Println(accountId)

}
