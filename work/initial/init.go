package initial

import (
	kconf "goapi/lib/conf"
	kredis "goapi/lib/db/redis"
	kmysql "goapi/lib/db/mysql"
)
func init() {
	kconf.SetFile("conf/config.cfg")
	kredis.RedisInit()
	kmysql.MysqlInit()
}