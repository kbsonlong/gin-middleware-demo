package main

import (
	"github.com/kbsonlong/gin-middleware-demo/routers"
)

func main() {

	router := routers.InitRouter()

	router.Run()
}
