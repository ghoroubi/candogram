package main

import (
	"github.com/ghoroubi/mt"

	"github.com/ghoroubi/mtx"
	"github.com/gin-gonic/gin"
	"fmt"
)

var m *mtprotox.MTProto
var nm *mtproto.MTProto

func main() {
	//var err error
	router := gin.Default()
	//Connecting Second Object to DC
Init()
fmt.Println("Started")
	router.GET("/send", SendMessageController)
	router.GET("/add/:phone", NewContactController)

	router.Run(":8181") // listen and serve on 0.0.0.0:8080
}
