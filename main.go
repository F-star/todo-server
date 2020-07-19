package main

import (
	"todo/config"
	initapp "todo/init"
	"todo/module/ginvalid"
	"todo/redisapp"
	"todo/routes"
)

func main() {
	defer config.DB.Close()
	initapp.SetEnv()
	ginvalid.RegisterValidation()

	initapp.SetORM()
	redisapp.StartRedis()
	r := routes.SetupRouter()
	r.Run()
}
