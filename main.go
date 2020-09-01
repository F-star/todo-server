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
	redisapp.StartRedis()
	ginvalid.RegisterValidation()
	initapp.SetORM()
	r := routes.SetupRouter()
	r.Run(":4000")
}
