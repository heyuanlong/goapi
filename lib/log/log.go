package log

import (
	"os"
	syslog "log"
	kconf "goklmmx/lib/conf"
)

var Klog *syslog.Logger

func init() {
	Klog = syslog.New(os.Stdout,"",syslog.LstdFlags | syslog.Lshortfile )
}

func SetLogfile( ) error  {
	var logfile string

	c, _ := kconf.GetConf()
	daemon,_ := c.Bool("server","daemon")
	if daemon == true{
		logfile,_ = c.String("server","logfile")
	}

	if logfile != ""{
		logFile,err  := os.Create(logfile)
		if err != nil {
			syslog.Fatalln("open file error")
		}
		Klog = syslog.New(logFile,"",syslog.LstdFlags | syslog.Lshortfile )
	}

	return nil
}
