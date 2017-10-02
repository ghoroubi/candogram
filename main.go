package main

import (
	"github.com/ghoroubi/mt"

	"log"
	"os"

	"github.com/ghoroubi/mtx"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var M *mtprotox.MTProto
var Mx *mtproto.MTProto
var LogFile *os.File
var DB *sqlx.DB

func main() {
	//var err error
	LoggerInit()
	InitSend()
	InitImport()
	router := gin.Default()
	//Connecting Second Object to DC
	log.Println("Started")
	router.GET("/send", SendMessageController)
	router.GET("/new", NewContactController)

	router.Run(":9191") // listen and serve on 0.0.0.0:8080
}
