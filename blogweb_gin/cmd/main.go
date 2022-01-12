package main

import (
	"blogweb_gin/routers"
	"blogweb_gin/tools"
	"fmt"
)

func main() {
	config, err := tools.ParseConfig("./config/conf.json")
	if err != nil {
		panic(err)
	}
	_, err = tools.OrmEngine(config)
	if err != nil {
		fmt.Println("init orm engine failed, err : ", err.Error())
		return
	}
	router := routers.InitRouter()
	router.Static("/static", "./static")
	router.Run(":8081")
}
