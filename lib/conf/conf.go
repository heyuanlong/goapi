package conf

import (
	"github.com/robfig/config"
	"fmt"
)
var logfile string

func SetFile(f string)  {
	logfile = f
}
func GetConf() (*config.Config,error) {
	c, err := config.ReadDefault(logfile)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return c,nil
}





