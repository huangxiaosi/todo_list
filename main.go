package main

import (
	"todo_list/conf"
	"todo_list/routes"
)

func main() {
	conf.Init()
	r := routes.NewRoute()
	_ = r.Run(conf.HttpPort)

}
