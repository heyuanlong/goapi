# goapi


##依赖包
- go get github.com/robfig/config       (deal config)
- go get github.com/gorilla/mux         (router)
- go get github.com/go-sql-driver/mysql
- go get github.com/garyburd/redigo/redis
- go get github.com/Jeffail/gabs        (deal json)

##关于https
- [https://studygolang.com/articles/4461](https://studygolang.com/articles/4461)

##使用
- 1.cd work
- 2.go run main.go
- 3.[http://127.0.0.1:8000/](http://127.0.0.1:8000/)
    [http://127.0.0.1:8000/testjson](http://127.0.0.1:8000/testjson)


##注意
如果修改了conf/config.cfg的相关配置，才能正常使用 
- http://127.0.0.1:8000/testredis
- http://127.0.0.1:8000/testmysql

