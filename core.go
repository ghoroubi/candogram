package main

import (
	"os"
	"github.com/ghoroubi/mt"
	"github.com/ghoroubi/mtx"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerInit() {
	var err error
	LogFile, err = os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	log.New(LogFile, time.Now().String(), 0)
	log.SetOutput(LogFile)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer LogFile.Close()
}
func InitImport() {
	var err error

	//EhMTP init and Connect
	Mx, err = mtproto.NewMTProto(os.Getenv("HOME")+"/.telegram_go", "", 1)
	if err != nil {
		log.Printf("Create failed: %s\n", err)

	}
	err = Mx.Connect()
	if err != nil {
		log.Printf("Connect failed: %s\n", err)

	}
}
func InitSend() {
	var err error

	//MainMTP init and Connect
	M, err = mtprotox.NewMTProto(os.Getenv("HOME") + "/.telegram_go")
	if err != nil {
		log.Printf("Create failed: %s\n", err)

	}

	//Connecting to telegram Nearest DataCenter
	err = M.Connect()
	if err != nil {
		log.Printf("Connect failed: %s\n", err)

	}
}
func NewContactController(c *gin.Context) {
	log.Println(" NewContactController")
	var id int32
	phone := c.Query("phone")
	log.Println("Phone Number Entered:", phone)
	if CheckContactExisting(phone) {
		log.Printf("\n%s:phone is already exist.\n", phone)
		c.BindJSON(gin.H{"Already exist": phone})
		return
	} else {
	 	NewContact(phone)
		id = GetUserID(phone)
		log.Println(id)
		c.BindJSON(gin.H{"New Contact ID": id})
		return
	}
}
func NewContact(phone string) {
	log.Println(" NewContact")
	firstName, lastName := phone, "new Contact"
	contact := mtproto.Contact{Firstname: firstName, Lastname: lastName, Phone: phone}
	Mx.Contacts_ImportContacts([]mtproto.TL{contact.GetInputContact()})

}
func SendMessageController(c *gin.Context) {
	log.Println(" SendMessageController")
	var userId int32
	phone := c.Query("phone")
	message := c.Query("message")
	log.Println("Debug:Phone:", phone, " Message:", message, "Imported UserID:", userId)
	if CheckContactExisting(phone) {
		log.Printf("\n%s:phone is already exist.\n", phone)
		userId = GetUserID(phone)

	go 	M.SendMsg(userId, message)
		c.JSON(200, gin.H{"status": "Send"})
		log.Println("Debug:Phone:", phone, " Message:", message, "Imported UserID:", userId)
	} else {

		NewContact(phone)
		userId = GetUserID(phone)

	go 	M.SendMsg(userId, message)
		c.JSON(200, gin.H{"status": "Send"})
	}

}
func CheckContactExisting(phone string) bool {
	var is bool
	log.Println(" CheckContactExisting")
	users := Mx.Contacts_GetContacts()
	for _, v := range users {
		log.Println("phone:", phone, "in contact phone:", v.Phone)
		if v.Phone == phone {
			is = true
			break
		} else {
			is = false
			break
		}
	}
log.Println("is:",is)
	return is
}

func GetUserID(phone string) int32 {
	log.Println("GetUserID")
	var uid int32 = 0
	users := Mx.Contacts_GetContacts()
	for _, v := range users {
		log.Println("phone:", phone, "v.phone", v.Phone)
		if v.Phone == phone {
			uid = v.ID
			return uid
		}
	}
	return uid

}
