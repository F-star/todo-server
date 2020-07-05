package main

import (
	"todo/config"
	initapp "todo/init"
	"todo/module/ginvalid"
	"todo/redis"
	"todo/routes"
)

func main() {
	defer config.DB.Close()
	initapp.SetEnv()
	ginvalid.RegisterValidation()

	initapp.SetORM()
	redis.StartRedis()
	r := routes.SetupRouter()
	r.Run()
}
