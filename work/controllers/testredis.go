package controllers

import (
	"net/http"
	"github.com/garyburd/redigo/redis"

	kredis "goapi/lib/db/redis"
	klog "goapi/lib/log"
)

func Testredis(w http.ResponseWriter, r *http.Request) {
	//https://godoc.org/github.com/garyburd/redigo/redis

	redis_set()
	redis_get()

	redis_list_set()
	redis_list_get()

	redis_hash_set()
	redis_hash_get()

	redis_key_exists()
	redis_set_expire()

	w.Write([]byte("goapi Testredis!\n"))
}




func redis_set()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)
	_, err := c.Do("set", "key1", "red")
	if err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}
func redis_get()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	v, err := redis.String(c.Do("get", "key1"))
	if err !=nil {
		klog.Klog.Println(err)
	}
	klog.Klog.Println(v)
}

func redis_list_set()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	if _, err := c.Do("del", "redlist");err != nil {
		klog.Klog.Println(err)
		return
	}

	if _, err := c.Do("lpush", "redlist", "qqq");err != nil {
		klog.Klog.Println(err)
		return
	}
	if _, err := c.Do("lpush", "redlist", "www");err != nil {
		klog.Klog.Println(err)
		return
	}
	if _, err := c.Do("lpush", "redlist", "eee");err != nil {
		klog.Klog.Println(err)
		return
	}
	return
}
func redis_list_get()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	values, err := redis.Strings(c.Do("lrange", "redlist","0","100"))
	if err !=nil {
		klog.Klog.Println(err)
		return
	}
	for _, v :=  range values {
		klog.Klog.Println(v)
	}
	return
}

func redis_hash_set()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	if _, err := c.Do("hset", "hkey1", "field1", "value1");err != nil {
		klog.Klog.Println(err)
		return
	}
	if _, err := c.Do("hset", "hkey1", "field2", "value2");err != nil {
		klog.Klog.Println(err)
		return
	}

	if _, err := c.Do("hmset", "hkey2", "field1", "value1", "field2", "value2");err != nil {
		klog.Klog.Println(err)
		return
	}

	return
}
func redis_hash_get()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	v, err := redis.String(c.Do("hget", "hkey1", "field1"))
	if err !=nil {
		klog.Klog.Println(err)
	}
	klog.Klog.Println(v)

	values, err_ := redis.StringMap(c.Do("hgetall", "hkey1"))
	if err_ !=nil {
		klog.Klog.Println(err_)
		return
	}
	for k, v :=  range values {
		klog.Klog.Println(k,v)
	}

	values2, err2_ := redis.StringMap(c.Do("hgetall", "hkey2"))
	if err2_ !=nil {
		klog.Klog.Println(err2_)
		return
	}
	for k, v :=  range values2 {
		klog.Klog.Println(k,v)
	}

	return
}




func redis_key_exists()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	is, err := redis.Bool(c.Do("EXISTS", "xxxxxxx"))
	if err !=nil {
		klog.Klog.Println(err)
		return
	}
	klog.Klog.Println(is)
	return
}

func redis_set_expire()  {
	c := kredis.GetRedis()
	defer kredis.CloseRedis(c)

	if _, err := c.Do("EXPIRE", "hkey1", 24*3600);err != nil {
		klog.Klog.Println(err)
		return
	}

	return
}