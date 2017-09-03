package conf

import (
	"github.com/robfig/config"
)
var logfile string

func SetFile(f string)  {
	logfile = f
}
func GetConf() (*config.Config,error) {
	c, err := config.ReadDefault(logfile)
	if err != nil {
		return nil,err
	}
	return c,nil
}





