package controllers

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"

	kmysql "goapi/lib/db/mysql"
	klog "goapi/lib/log"
)

func Testmysql(w http.ResponseWriter, r *http.Request) {
	add()
	update()
	query()
	del()
	query()

	w.Write([]byte("goapi Testmysql!\n"))
}

func query()  {
	stmt, err := kmysql.MysqlClient.Prepare(`select accountid,version ,logtime from log_first_login`)
	defer stmt.Close()
	if err != nil {
		klog.Klog.Println(err)
		return
	}
	res, err_ := stmt.Query( )
	defer res.Close()
	if err_ != nil {
		klog.Klog.Println(err_)
		return
	}
	for res.Next(){
		var accountId int
		var version string
		var logtime string

		if err := res.Scan(&accountId,&version,&logtime); err != nil {
			klog.Klog.Println(err)
			return
		}
		klog.Klog.Printf("accountId=%d,version=%s,logtime=%s",accountId,version,logtime)
	}
}
func add()  {
	stmt, err := kmysql.MysqlClient.Prepare(`INSERT log_first_login (accountid,version,logtime) values (?,?,UNIX_TIMESTAMP())`)
	defer stmt.Close()
	if err != nil {
		klog.Klog.Println(err)
		return
	}
	res, err_ := stmt.Exec( 104,"v1.0")
	if err_ != nil {
		klog.Klog.Println(err_)
		return
	}
	id, err__ := res.LastInsertId()
	if err__ != nil {
		klog.Klog.Println(err__)
		return
	}
	klog.Klog.Println(id)
}

func update()  {
	stmt, err := kmysql.MysqlClient.Prepare(`update log_first_login set version=? where accountid=?`)
	defer stmt.Close()
	if err != nil {
		klog.Klog.Println(err)
		return
	}
	res, err_ := stmt.Exec( "v1.2",104)
	if err_ != nil {
		klog.Klog.Println(err_)
		return
	}
	id, err__ := res.RowsAffected()
	if err__ != nil {
		klog.Klog.Println(err__)
		return
	}
	klog.Klog.Println(id)
}
func del()  {
	stmt, err := kmysql.MysqlClient.Prepare(`delete from  log_first_login where accountid=?`)
	defer stmt.Close()
	if err != nil {
		klog.Klog.Println(err)
		return
	}
	res, err_ := stmt.Exec( 104)
	if err_ != nil {
		klog.Klog.Println(err_)
		return
	}
	id, err__ := res.RowsAffected()
	if err__ != nil {
		klog.Klog.Println(err__)
		return
	}
	klog.Klog.Println(id)

}