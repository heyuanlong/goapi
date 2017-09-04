package models

import (
	_ "github.com/go-sql-driver/mysql"

	kmysql "goapi/lib/db/mysql"
	klog "goapi/lib/log"
)
func AddLog(accountid int ,version string) int {
	stmt, err := kmysql.MysqlClient.Prepare(`INSERT log_first_login (accountid,version,logtime) values (?,?,UNIX_TIMESTAMP())`)
	defer stmt.Close()
	if err != nil {
		klog.Klog.Println(err)
		return 0
	}
	res, err_ := stmt.Exec( accountid,version)
	if err_ != nil {
		klog.Klog.Println(err_)
		return 0
	}
	id, err__ := res.LastInsertId()
	if err__ != nil {
		klog.Klog.Println(err__)
		return 0
	}
	return int(id)
}