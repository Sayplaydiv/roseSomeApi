package main

import (
	"roseSomeApi/router"
)

const (
	PATH="config/conf.ini"
)
func main()  {
	router.InitRouter(PATH)
}


